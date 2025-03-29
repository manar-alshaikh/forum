package main

import (
	"database/sql"
	"fmt"
	"forum/functions"
	"log"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type Forum struct {
    Users []functions.User
    Posts []functions.Post
    Comments []functions.Comments
    Complaints []functions.Complaints
    PComplaints []functions.PComplaints
}

func main() {
    // Open a connection to the SQLite database
    db, err := sql.Open("sqlite3", "./Database/database.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
//     insertCommentComplaint := `INSERT INTO complaints (username, post_id, comment_id) VALUES (?, ?, ?)`
// _, err = db.Exec(insertCommentComplaint, "noorz", 1, 1) // Example data: user "noorz" complains about comment ID 1 in post ID 1
// if err != nil {
//     log.Fatal(err)
// }
// insertPostComplaint := `INSERT INTO complaintsPosts (username, post_id) VALUES (?, ?)`
// _, err = db.Exec(insertPostComplaint, "haleema", 2) // Example data: user "haleema" complains about post ID 2
// if err != nil {
//     log.Fatal(err)
// }
	functions.Createdb(db)
    fmt.Println("Database created successfully")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var forum Forum
        forum.Users = *functions.FetchAllUsersDescendingBlockedPosts(db)
        fmt.Println("Users")
        forum.Posts = *functions.FetchAllPosts(db)
        fmt.Println("Posts")
        forum.Comments = *functions.FetchAllComments(db)
        fmt.Println("Comments")
        forum.Complaints = *functions.FetchAllComplaints(db)
        fmt.Println("Complaints", len(*&forum.Complaints))
        fmt.Println("Complaints", len(*&forum.Comments))

        forum.PComplaints = *functions.FetchAllPComplaints(db)
        fmt.Println("Complaints", len(*&forum.PComplaints))
        // users := functions.FetchAllUsers(db)
        // posts := functions.FetchAllPosts(db)
        // comments := functions.FetchAllComments(db)
        // complaints := functions.FetchAllComplaints(db)
        http.HandleFunc("/filter/posts", functions.FilterPostsHandler(db))
        tmpl, err := template.ParseFiles("./templets/index.html")
		// log.Println("hii")
        if err != nil {
            log.Fatal(err)
        }
        tmpl.Execute(w, forum)
    })

    fmt.Println(functions.CountPosts(db))
    log.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}