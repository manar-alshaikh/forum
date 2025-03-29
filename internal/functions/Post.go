package functions

// import (
// 	"database/sql"
// 	"log"
// 	"net/http"
// 	// "text/template"
// )

// // func InsertPost(db *sql.DB, post *Post) {
// // 	insertPost := `INSERT INTO posts (
// //         title, content, categories, image, username, time
// //     ) VALUES (?, ?, ?, ?,?,datetime('now'))`

// // 	result, err := db.Exec(insertPost,
// //         post.Title,
// //         post.Content,
// //         post.Categories,
// //         post.Image,
// //         post.Username)
// //     if err != nil {
// //         log.Fatal(err)
// //     }

// // 	lastID, err := result.LastInsertId()
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	post.Post_id = int(lastID)

// // 	db.Exec(insertPost, "Welcome to Forum", "This is our first post!", "General", "", "admin")
// //     db.Exec(insertPost, "Gaming Discussion", "What games are you playing?", "Entertainment", "", "gamer123")
// //     db.Exec(insertPost, "Travel Tips", "Share your best travel experiences", "Travel", "", "traveler")
// // }

// // // Example post data
// // _, err = db.Exec(insertPost, "welcoming", "Hello, World!", "General", "", "john_doe")
// // if err != nil {
// // 	log.Fatal(err)
// // }

// // _, err = db.Exec(insertPost, "welcoming", "This is a test post.", "Sport", "image.jpg", "noorz")
// // if err != nil {
// // 	log.Fatal(err)
// // }

// // _, err = db.Exec(insertPost, "welcoming", "Another post.", "General", "", "haleema")
// // if err != nil {
// // 	log.Fatal(err)
// // }

// // _, err = db.Exec(insertPost, "welcoming", "Hello, World!", "General", "image.jpg", "Kawther")
// // if err != nil {
// // 	log.Fatal(err)
// // }

// // testPost := &Post{
// // 	Title:      "Test Post",
// // 	Content:    "This is a test post",
// // 	Categories: "Test",
// // 	Username:   "admin",
// // }

// // InsertPost(db, testPost)

// func InsertPost(db *sql.DB) {
// 	insertPost := `INSERT INTO posts (
// 			title, content, categories, image, username, time
// 		) VALUES (?, ?, ?, ?, ?, datetime('now'))`

// 	// Example post data
// 	_, err := db.Exec(insertPost, "welcoming", "Hello, World!", "General", "", "john_doe")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = db.Exec(insertPost, "welcoming", "This is a test post.", "Sport", "image.jpg", "noorz")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = db.Exec(insertPost, "welcoming", "Another post.", "General", "", "haleema")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = db.Exec(insertPost, "welcoming", "Hello, World!", "General", "image.jpg", "Kawther")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func CreatPostHandler(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == "Post" {
// 			http.Error(w, "Methode not allowed", http.StatusMethodNotAllowed)
// 			return
// 		}

// 		// Parse the multipart form data (for image upload)
// 		r.ParseMultipartForm(10 << 20) // 10MB max size

// 		title := r.FormValue("title")
// 		content := r.FormValue("content")
// 		categories := r.FormValue("categories")
// 		username := GetCurrentUsername(r)

// 		file, handler, err := r.FormFile("image")
// 		imagePath := ""
// 		if err == nil {
// 			defer file.Close()
// 			imagePath = handler.Filename
// 		}

// 		insertPost := `INSERT INTO posts (title, content, categories, image, username, time)
// 	 VALUES (?, ?, ?, ?, ?, datetime('now'))`

// 		_, err = db.Exec(insertPost, title, content, categories, imagePath, username)
// 		if err != nil {
// 			log.Fatal(err)
// 			http.Error(w, "Error creatin post", http.StatusInternalServerError)
// 			return
// 		}
// 		http.Redirect(w, r, "/", http.StatusSeeOther)

// 	}
// }
