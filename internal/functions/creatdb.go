package functions

import (
	"database/sql"
	"log"
)

func Createdb(db *sql.DB) {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
    username TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    hide_name BOOLEAN NOT NULL DEFAULT 0 CHECK(hide_name IN (0, 1)),
    email TEXT NOT NULL UNIQUE,
    hide_email BOOLEAN NOT NULL DEFAULT 0 CHECK(hide_email IN (0, 1)),
    date_of_birth DATE NOT NULL,
    hide_age BOOLEAN NOT NULL DEFAULT 0 CHECK(hide_age IN (0, 1)),
    hobby TEXT NOT NULL DEFAULT 'NONE',
    hide_hobby BOOLEAN NOT NULL DEFAULT 0 CHECK(hide_hobby IN (0, 1)),
    role TEXT NOT NULL DEFAULT 'USER' CHECK(role IN ('USER', 'ADMIN')),
    profile_picture TEXT NOT NULL DEFAULT 'NONE',
    password TEXT NOT NULL,
    blocked_posts INTEGER NOT NULL DEFAULT 0,
    banned BOOLEAN NOT NULL DEFAULT 0 CHECK(banned IN (0, 1)),
    banned_date DATE NOT NULL DEFAULT '0000-00-00'
);`

	createPostsTable := `CREATE TABLE IF NOT EXISTS posts (

        post_id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        categories TEXT,
        image TEXT,
        time DATETIME DEFAULT CURRENT_TIMESTAMP,
        username TEXT NOT NULL,
     
        FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
    );`

	createLikesTable := `CREATE TABLE IF NOT EXISTS likes (
        like_id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER,
        comment_id INTEGER,
        username TEXT NOT NULL,
        like_type TEXT NOT NULL CHECK(like_type IN ('LIKE', 'DISLIKE')),
        FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
        FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
        FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
    );`

	createCommentsTable := `CREATE TABLE IF NOT EXISTS comments (
        comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
        post_id INTEGER NOT NULL,
        content TEXT NOT NULL,
        created_date DATETIME DEFAULT CURRENT_TIMESTAMP,
        username TEXT NOT NULL,
        pin BOOLEAN NOT NULL DEFAULT 0 CHECK(pin IN (0, 1)),
        FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
        FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
    );`

	createComplaintsTable := `CREATE TABLE IF NOT EXISTS complaints (
        complaints_id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        post_id INTEGER,
        comment_id INTEGER,
        FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
        FOREIGN KEY (comment_id) REFERENCES comments(comment_id) ON DELETE CASCADE,
        FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
    );`
    createComplaintsPostsTable := `CREATE TABLE IF NOT EXISTS complaintsPosts (
        complaints_id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        post_id INTEGER,
        FOREIGN KEY (post_id) REFERENCES posts(post_id) ON DELETE CASCADE,
        FOREIGN KEY (username) REFERENCES users(username) ON DELETE CASCADE
    );`
	
	statements := []string{createUsersTable, createPostsTable, createLikesTable, createCommentsTable, createComplaintsTable, createComplaintsPostsTable}
	for _, stmt := range statements {
		_, err := db.Exec(stmt)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Println("Tables created successfully.")
	// InsertUser(db)
    // InsertPost(db)
    // InsertComment(db)
    // InsertComplaint(db)

}



func InsertUser(db *sql.DB) {
	insertUser := `INSERT INTO users (
        username, name, hide_name, email, hide_email,
        date_of_birth, hide_age, hobby, hide_hobby,
        role, password, blocked_posts, banned
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	// Example user data
	_, err := db.Exec(insertUser, "john_doe", "John Doe", 0, "john@example.com", 0,
		"1990-01-01", 0, "Reading", 0, "USER", "password123", 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(insertUser, "noorz", "Noor Zaher", 0, "noorz@example.com", 0,
		"1990-01-01", 0, "Reading", 0, "ADMIN", "password123", 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(insertUser, "haleema", "haleema abbas", 0, "haleema@example.com", 0,
		"1990-01-01", 0, "Reading", 0, "USER", "password123", 3, 0)
	if err != nil {
		log.Fatal(err)
	}

	insertUser = `INSERT INTO users (
        username, name, email,
        date_of_birth, hobby, 
        role, password
    ) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err = db.Exec(insertUser, "Kawther", "Kawther abbas", "Kawther@example.com",
		"1990-01-01", "Reading", "USER", "password123")
	if err != nil {
		log.Fatal(err)
	}
}



func InsertComment(db *sql.DB) {
	insertComment := `INSERT INTO comments (
        post_id, content, username
    ) VALUES (?, ?, ? )`

	// Example comment data
	_, err := db.Exec(insertComment, 1, "This is a test comment.", "noorz")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(insertComment, 2, "Another comment.", "haleema")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(insertComment, 3, "Hello, World!", "Kawther")
	if err != nil {
		log.Fatal(err)
	}
}

func InsertComplaint(db *sql.DB) {
	insertComplaint := `INSERT INTO complaints (
        username, post_id, comment_id
    ) VALUES (?, ?, ?)`

	// Example complaint data
	_, err := db.Exec(insertComplaint, "noorz", 1, 1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(insertComplaint, "haleema", 2, 1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(insertComplaint, "Kawther", 3, 0)
	if err != nil {
		log.Fatal(err)
	}
}

