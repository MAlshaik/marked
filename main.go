package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
    "sync"
    "time"

    "github.com/MAlshaik/marked/handlers"
    "github.com/gorilla/websocket"
    _ "github.com/lib/pq"
)

var (
    db           *sql.DB
    upgrader     = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool {
            return true // Be cautious with this in production
        },
    }
    clients       = make(map[*websocket.Conn]bool)
    clientsMutex  sync.Mutex
    broadcast     = make(chan Message)
)

type Message struct {
    Type    string `json:"type"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

func main() {
    fmt.Println("Connecting to Postgres...")

    postgresURL := os.Getenv("POSTGRES_URL")

    log.Println("Attempting to connect to database...")
    var err error
    db, err = sql.Open("postgres", postgresURL)
    if err != nil {
        log.Fatal("Error opening database connection: ", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal("Error pinging database: ", err)
    }
    log.Println("Successfully connected to database")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8000" // fallback to 8000 if not set
    }

    app := http.NewServeMux()

    app.HandleFunc("GET /", handlers.HandleRoot(db))
    app.HandleFunc("GET /create", handlers.CreateDocument(db))
    app.HandleFunc("GET /document/{id}", handlers.GetDocument(db))
    app.HandleFunc("POST /document/{id}", handlers.UpdateDocument(db))
    app.HandleFunc("POST /document/{id}/delete", handlers.DeleteDocument(db))
    app.HandleFunc("GET /static/", serveStatic)

    app.HandleFunc("GET /ws/{id}", handleWebSocket)

    app.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    go handleMessages()

    log.Println("Server starting...")
    log.Println("Listening on port:", port)
    log.Fatal(http.ListenAndServe(":"+port, app))
}

func serveStatic(w http.ResponseWriter, r *http.Request) {
    path := r.URL.Path[len("/static/"):]
    if strings.HasSuffix(path, "/") {
        http.NotFound(w, r)
        return
    }
    http.ServeFile(w, r, "./static/"+path)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    documentID := r.URL.Path[len("/ws/"):]

    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading to WebSocket:", err)
        return
    }
    defer conn.Close()

    // Set up ping-pong to keep connection alive
    conn.SetReadDeadline(time.Now().Add(60 * time.Second))
    conn.SetPongHandler(func(string) error {
        conn.SetReadDeadline(time.Now().Add(60 * time.Second))
        return nil
    })

    // Start a goroutine to send pings
    go func() {
        ticker := time.NewTicker(30 * time.Second)
        defer ticker.Stop()
        for {
            select {
            case <-ticker.C:
                if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(10*time.Second)); err != nil {
                    log.Println("Ping error:", err)
                    return
                }
            }
        }
    }()

    clientsMutex.Lock()
    clients[conn] = true
    clientsMutex.Unlock()

    for {
        var msg Message
        err := conn.ReadJSON(&msg)
        if err != nil {
            log.Println("Error reading message:", err)
            clientsMutex.Lock()
            delete(clients, conn)
            clientsMutex.Unlock()
            break
        }

        msg.Type = "update" // Ensure type is set for broadcasting
        broadcast <- msg

        if msg.Type == "save" {
            // Update the document in the database
            _, err = db.Exec("UPDATE documents SET title = $1, content = $2 WHERE id = $3", msg.Title, msg.Content, documentID)
            if err != nil {
                log.Println("Error updating document in database:", err)
                continue
            }
        }
    }
}

func handleMessages() {
    for {
        msg := <-broadcast
        clientsMutex.Lock()
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("Error broadcasting message: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
        clientsMutex.Unlock()
    }
}
