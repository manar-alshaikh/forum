package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func CreateTables(db *sql.DB) (error) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			username TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			hide_name BOOLEAN NOT NULL DEFAULT 0 CHECK(hide_name IN (0, 1)),
			email TEXT NOT NULL UNIQUE,
			hide_email BOOLEAN NOT NULL DEFAULT 0 CHECK(hide_email IN (0, 1)),
			date_of_birth DATE NOT NULL,
			hide_age BOOLEAN NOT NULL DEFAULT 0 CHECK(hide_age IN (0, 1)),
			hobby TEXT,
			hide_hobby BOOLEAN NOT NULL DEFAULT 0 CHECK(hide_hobby IN (0, 1)),
			role TEXT NOT NULL DEFAULT 'USER' CHECK(role IN ('USER', 'ADMIN')),
			profile_picture TEXT,
			password VARCHAR(60) NOT NULL,
			blocked_posts INTEGER NOT NULL DEFAULT 0,
			banned BOOLEAN NOT NULL DEFAULT 0 CHECK(banned IN (0, 1))
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
			post_id INTEGER PRIMARY KEY AUTOINCREMENT,
			content TEXT NOT NULL,
			categories TEXT,
			image TEXT,
			time DATETIME DEFAULT CURRENT_TIMESTAMP,
			username TEXT NOT NULL,
			FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS likes (
			like_id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER,
			comment_id INTEGER,
			username TEXT NOT NULL,
			like_type TEXT NOT NULL CHECK(like_type IN ('LIKE', 'DISLIKE')),
			FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
			FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
			FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS comments (
			comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
			post_id INTEGER NOT NULL,
			content TEXT NOT NULL,
			created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
			username TEXT NOT NULL,
			pin BOOLEAN NOT NULL CHECK(pin IN (0, 1)),
			FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
			FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS complaints (
			complaints_id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			post_id INTEGER,
			comment_id INTEGER,
			FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
			FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
			FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
		);`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	log.Println("Tables created successfully.")
	return nil
}

func InitDummyData(db *sql.DB) {
	dummyUsers := []struct {
		username    string
		name        string
		email       string
		dateOfBirth string
		hobby       string
		password    string
	}{
		{"john_doe", "John Doe", "john@example.com", "1990-01-01", "Reading", "password123"},
		{"jane_smith", "Jane Smith", "jane@example.com", "1992-05-15", "Traveling", "password456"},
		{"admin_user", "Admin User", "admin@example.com", "1985-08-20", "Gaming", "admin123"},
	}

	for _, user := range dummyUsers {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error hashing password for user %s: %v", user.username, err)
			continue
		}

		_, err = db.Exec(`
			INSERT OR IGNORE INTO users (
				username, name, email, date_of_birth, hobby, password
			) VALUES (?, ?, ?, ?, ?, ?)`,
			user.username, user.name, user.email, user.dateOfBirth, user.hobby, hashedPassword)
		if err != nil {
			log.Printf("Error inserting dummy data for user %s: %v", user.username, err)
		}
	}
}

// printTable prints all rows from a given table.
func printTable(db *sql.DB, tableName string) error {
	colQuery := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	rows, err := db.Query(colQuery)
	if err != nil {
		return err
	}
	defer rows.Close()

	var columnNames []string
	for rows.Next() {
		var cid int
		var name, ctype string
		var notnull, pk int
		var dfltValue sql.NullString
		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dfltValue, &pk); err != nil {
			return err
		}
		columnNames = append(columnNames, name)
	}
	if err = rows.Err(); err != nil {
		return err
	}

	fmt.Println(columnNames)

	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	dataRows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer dataRows.Close()

	for dataRows.Next() {
		columns := make([]interface{}, len(columnNames))
		columnPointers := make([]interface{}, len(columnNames))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := dataRows.Scan(columnPointers...); err != nil {
			return err
		}

		for _, col := range columns {
			fmt.Printf("%v\t", col)
		}
		fmt.Println()
	}

	return dataRows.Err()
}
