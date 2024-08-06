package main

import (
	"database/sql"
	"encoding/json"
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
)

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
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            break
        }

        var message struct {
            Type    string `json:"type"`
            Title   string `json:"title"`
            Content string `json:"content"`
        }
        err = json.Unmarshal(p, &message)
        if err != nil {
            log.Println("Error unmarshaling message:", err)
            continue
        }

        if message.Type == "update" {
            // Only broadcast to other clients
            broadcastMessageExcept(messageType, p, conn)
        } else if message.Type == "save" {
            // Update the document in the database
            _, err = db.Exec("UPDATE documents SET title = $1, content = $2 WHERE id = $3", message.Title, message.Content, documentID)
            if err != nil {
                log.Println("Error updating document in database:", err)
                continue
            }

            // Broadcast a 'saved' message to all clients
            savedMsg, _ := json.Marshal(struct {
                Type    string `json:"type"`
                Title   string `json:"title"`
                Content string `json:"content"`
            }{
                Type:    "saved",
                Title:   message.Title,
                Content: message.Content,
            })
            broadcastMessage(websocket.TextMessage, savedMsg)
        }
    }

    clientsMutex.Lock()
    delete(clients, conn)
    clientsMutex.Unlock()
}

func broadcastMessageExcept(messageType int, message []byte, except *websocket.Conn) {
    clientsMutex.Lock()
    defer clientsMutex.Unlock()

    for client := range clients {
        if client != except {
            err := client.WriteMessage(messageType, message)
            if err != nil {
                log.Printf("Error broadcasting message: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func broadcastMessage(messageType int, message []byte) {
    clientsMutex.Lock()
    defer clientsMutex.Unlock()

    for client := range clients {
        err := client.WriteMessage(messageType, message)
        if err != nil {
            log.Printf("Error broadcasting message: %v", err)
            client.Close()
            delete(clients, client)
        }
    }
}
