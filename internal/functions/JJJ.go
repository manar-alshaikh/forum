package functions

import (
    "database/sql"
    "log"
)

// CountPosts returns the number of rows in the posts table
func CountPosts(db *sql.DB) int {
    var count int
    err := db.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
    if err != nil {
        log.Fatal(err)
    }
    return count
}