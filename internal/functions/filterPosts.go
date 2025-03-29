package functions

import (
	"database/sql"
	"log"
	"net/http"
)

type filterPost struct {
    PostID      int
    Content     string
    Categories  string
    Image       string
    Time        string
    Username    string
}

func FilterPosts(db *sql.DB, filterType string, category string, username string) []filterPost {
    var rows *sql.Rows
    var err error

    switch filterType {
    case "categories":
        query := `SELECT post_id, content, categories, image, time, username FROM posts WHERE categories LIKE ?`
        rows, err = db.Query(query, "%"+category+"%")
    case "myPosts":
        query := `SELECT post_id, content, categories, image, time, username FROM posts WHERE username = ?`
        rows, err = db.Query(query, username)
    case "likedPosts":
        query := `SELECT p.post_id, p.content, p.categories, p.image, p.time, p.username FROM posts p JOIN likes l ON p.post_id = l.post_id WHERE l.username = ? AND l.like_type = 'LIKE'`
        rows, err = db.Query(query, username)
    }

    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

   var posts []filterPost
   for rows.Next() {
        var post filterPost
        err := rows.Scan(&post.PostID, &post.Content, &post.Categories, &post.Image, &post.Time, &post.Username)
        if err != nil {
            log.Fatal(err)
            continue
        }
        posts = append(posts, post)
   }
   return posts
}
func FilterPostsHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        filterType := r.URL.Query().Get("filterType")
        category := r.URL.Query().Get("category")
        username := ""

        posts := FilterPosts(db, filterType, category, username)
        if posts == nil {
            http.Error(w, "No posts found", http.StatusNotFound)
            return
        }
    }
}





