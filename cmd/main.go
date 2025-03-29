package main

import (
	"database/sql"
	"log"
	"net/http"
	"forum/database"

	"forum/internal/functions"
	"forum/internal/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database
	db, err := sql.Open("sqlite3", "database/mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables if they don't exist
	err = database.CreateTables(db)
	if err != nil {
		log.Fatal(err)
	}

	database.InitDummyData(db)

	functions.InitDB(db)
	
	handlers.InitHandlers(db)

	// Setup routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/dashboard", handlers.DashboardHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
