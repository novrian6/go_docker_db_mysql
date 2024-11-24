package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialize DB connection and create the messages table
func init() {
	var err error
	db, err = sql.Open("mysql", "root:rootpassword@tcp(db:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}

	// Ensure the messages table exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id INT AUTO_INCREMENT PRIMARY KEY,
		content VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

// Fetch all messages from the database
func getMessages(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, content, created_at FROM messages")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var id int
		var content string
		var createdAt string
		if err := rows.Scan(&id, &content, &createdAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		messages = append(messages, fmt.Sprintf("ID: %d, Content: %s, CreatedAt: %s", id, content, createdAt))
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, msg := range messages {
		fmt.Fprintln(w, msg)
	}
}

func main() {
	http.HandleFunc("/messages", getMessages)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
