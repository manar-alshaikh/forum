// package functions

// import (
// 	"database/sql"
// 	// "fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// type Post struct {
// 	Post_id int
// 	Content string
// 	Categories string
// 	Image string
// 	Time string
// 	Username string
// 	NumberOfComments int
// }

// func FetchAllPosts(db *sql.DB) *[]Post {
// 	fetchAllPosts := `SELECT post_id, content, categories, image, time, username FROM posts`
// 	rows, err := db.Query(fetchAllPosts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var posts []Post
// 	for rows.Next() {
// 		var post Post
// 		err = rows.Scan(&post.Post_id, &post.Content, &post.Categories, &post.Image, &post.Time, &post.Username)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		post.NumberOfComments = CountCommentsbyPostId(db, post.Post_id)
// 		// fmt.Println(post.NumberOfComments)
// 		posts = append(posts, post)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &posts
// }

// func FetchPostById(db *sql.DB, post_id int) *Post {
// 	fetchPost:="SELECT post_id, content, categories, image, time, username FROM posts WHERE post_id = ?"
// 	var post Post
// 	row := db.QueryRow(fetchPost, post_id)
// 	if row == nil {
// 		return nil
// 	}
// 	err := row.Scan(&post.Post_id, &post.Content, &post.Categories, &post.Image, &post.Time, &post.Username)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &post
// }
package functions

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type Post struct {
    Post_id          int
    Title            string
    Content          string
    Categories       string
    Image            string
    Time             string
    Username         string
    NumberOfComments int
}

func FetchAllPosts(db *sql.DB) *[]Post {
    fetchAllPosts := `SELECT post_id, title, content, categories, image, time, username FROM posts`
    rows, err := db.Query(fetchAllPosts)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    var posts []Post
    for rows.Next() {
        var post Post
        err = rows.Scan(&post.Post_id, &post.Title, &post.Content, &post.Categories, &post.Image, &post.Time, &post.Username)
        if err != nil {
            log.Fatal(err)
        }
        post.NumberOfComments = CountCommentsbyPostId(db, post.Post_id)
        posts = append(posts, post)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
    return &posts
}

func FetchPostById(db *sql.DB, post_id int) *Post {
    fetchPost := "SELECT post_id, title, content, categories, image, time, username FROM posts WHERE post_id = ?"
    var post Post
    row := db.QueryRow(fetchPost, post_id)
    if row == nil {
        return nil
    }
    err := row.Scan(&post.Post_id, &post.Title, &post.Content, &post.Categories, &post.Image, &post.Time, &post.Username)
    if err != nil {
        log.Fatal(err)
    }
    return &post
}