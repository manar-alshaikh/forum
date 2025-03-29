package functions

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type Comments struct {
	comment_id int
	Username   string
	post_id	  int
	content	  string
}

func FetchAllComments(db *sql.DB) *[]Comments {
	fetchAllcomments := `SELECT comment_id, username, post_id, comment_id, content FROM comments`
	rows, err := db.Query(fetchAllcomments)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var comments []Comments
	for rows.Next() {
		var comment Comments
		err = rows.Scan(&comment.comment_id, &comment.Username, &comment.post_id, &comment.comment_id, &comment.content)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &comments
}

func FetchCommentById(db *sql.DB, comment_id int) *Comments {
	fetchComment:="SELECT comment_id, username, post_id, content FROM comments WHERE comment_id = ?"
	var comment Comments
	row := db.QueryRow(fetchComment, comment_id)
	if row == nil {
		return nil
	}
	err := row.Scan(&comment.comment_id, &comment.Username, &comment.post_id, &comment.content)
	if err != nil {
		log.Fatal(err)
	}
	return &comment
}

func FetchCommentsByPostId(db *sql.DB, post_id int) *[]Comments {
	fetchComments:="SELECT comment_id, username, post_id, content FROM comments WHERE post_id = ?"
	rows, err := db.Query(fetchComments, post_id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var comments []Comments
	for rows.Next() {
		var comment Comments
		err = rows.Scan(&comment.comment_id, &comment.Username, &comment.post_id, &comment.content)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &comments
}

func CountCommentsbyPostId(db *sql.DB, post_id int) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM comments WHERE post_id = ?", post_id).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}