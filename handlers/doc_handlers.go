package handlers

import (
    "database/sql"
    "fmt"
    "net/http"

    "github.com/MAlshaik/marked/models"
    "github.com/MAlshaik/marked/templates/pages"
    "github.com/google/uuid"
)

func HandleRoot(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Creating document")
        sessionID := uuid.New().String()
        
        document := models.Document{
            ID:      sessionID,
            Title:   "New Document",
            Content: "Start typing here...",
        }

        _, err := db.Exec("INSERT INTO documents (id, title, content) VALUES ($1, $2, $3)", document.ID, document.Title, document.Content)
        if err != nil {
            http.Error(w, "Error creating document", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/document/"+sessionID, http.StatusSeeOther)
    }
}

func GetDocument(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Getting document")

        documentID := r.URL.Path[len("/document/"):]

        var document models.Document
        err := db.QueryRow("SELECT id, title, content FROM documents WHERE id = $1", documentID).Scan(&document.ID, &document.Title, &document.Content)
        if err == sql.ErrNoRows {
            // Document not found, create a new one
            document = models.Document{
                ID:      documentID,
                Title:   "New Document",
                Content: "Start typing here...",
            }
            
            _, err := db.Exec("INSERT INTO documents (id, title, content) VALUES ($1, $2, $3)", document.ID, document.Title, document.Content)
            if err != nil {
                http.Error(w, "Error creating document", http.StatusInternalServerError)
                return
            }
        } else if err != nil {
            http.Error(w, "Error retrieving document", http.StatusInternalServerError)
            return
        }

        pages.DocumentPage(document).Render(r.Context(), w)
    }
}

func CreateDocument(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Creating document")
        sessionID := uuid.New().String()
        
        document := models.Document{
            ID:      sessionID,
            Title:   "New Document",
            Content: "Start typing here...",
        }

        _, err := db.Exec("INSERT INTO documents (id, title, content) VALUES ($1, $2, $3)", document.ID, document.Title, document.Content)
        if err != nil {
            http.Error(w, "Error creating document", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/document/"+sessionID, http.StatusSeeOther)
    }
}

func UpdateDocument(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        documentID := r.URL.Path[len("/document/"):]
        title := r.FormValue("title")
        content := r.FormValue("content")

        _, err := db.Exec("UPDATE documents SET title = $1, content = $2 WHERE id = $3", title, content, documentID)
        if err != nil {
            http.Error(w, "Error updating document", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/document/"+documentID, http.StatusSeeOther)
    }
}

func DeleteDocument(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
            return
        }

        documentID := r.URL.Path[len("/document/"):]

        _, err := db.Exec("DELETE FROM documents WHERE id = $1", documentID)
        if err != nil {
            http.Error(w, "Error deleting document", http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
}
