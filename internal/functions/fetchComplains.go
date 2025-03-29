// package functions

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/mattn/go-sqlite3"
// )

// type Complaints struct {
// 	Complaint_id int
// 	Username     string
// 	post	  *Post
// 	comment	  *Comments
// }

// func FetchAllComplaints(db *sql.DB) *[]Complaints {
// 	fetchAllComplaints := `SELECT complaints_id, username, post_id, comment_id FROM complaints`
// 	rows, err := db.Query(fetchAllComplaints)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	fmt.Println("Fetching all complaints")

// 	var complaints []Complaints
// 	for rows.Next() {
// 		var complaint Complaints
// 		var postid int
// 		var commentid int
// 		err = rows.Scan(&complaint.Complaint_id, &complaint.Username, postid, commentid)
// 		fmt.Println(&complaint.Complaint_id, &complaint.Username, postid, commentid)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println("Scan complaint")
// 		if postid != 0 {

// 		complaint.post = FetchPostById(db, postid)
// 		}
// 		if commentid != 0 {
// 		complaint.comment = FetchCommentById(db, commentid)
// 		}
// 		complaints = append(complaints, complaint)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return &complaints
// }

package functions

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type Complaints struct {
    Complaint_id int
    Username     string
    Post         int
    Comment      int
}
type PComplaints struct {
    Complaint_id int
    Username     string
    Post         int
}


func FetchAllComplaints(db *sql.DB) *[]Complaints {
    // SQL query to fetch all complaints
    fetchAllComplaints := `SELECT complaints_id, username, post_id, comment_id FROM complaints`
    rows, err := db.Query(fetchAllComplaints)
    if err != nil {
        log.Fatal("Error executing query:", err)
    }
    defer rows.Close()

    fmt.Println("Fetching all complaints...")

    // Slice to hold complaints
    var complaints []Complaints
    for rows.Next() {
        var complaint Complaints
        // Scan row data into the struct
        err = rows.Scan(&complaint.Complaint_id, &complaint.Username, &complaint.Post, &complaint.Comment)
        if err != nil {
            log.Fatal("Error scanning row:", err)
        }
        complaints = append(complaints, complaint)
    }

    // Check for any errors during iteration
    if err = rows.Err(); err != nil {
        log.Fatal("Error iterating rows:", err)
    }

    fmt.Println("Complaints fetched successfully:", complaints)
    return &complaints
}
func FetchAllPComplaints(db *sql.DB) *[]PComplaints {
    fetchAllPComplaints := `SELECT complaints_id, username, post_id FROM complaintsPosts`
    rows, err := db.Query(fetchAllPComplaints)
    if err != nil {
        log.Fatal("Error executing query:", err)
    }
    defer rows.Close()

    fmt.Println("Fetching all post complaints...")

    var complaints []PComplaints
    for rows.Next() {
        var complaint PComplaints
        err = rows.Scan(&complaint.Complaint_id, &complaint.Username, &complaint.Post)
        if err != nil {
            log.Fatal("Error scanning row:", err)
        }
        complaints = append(complaints, complaint)
    }

    if err = rows.Err(); err != nil {
        log.Fatal("Error iterating rows:", err)
    }

    fmt.Println("Post complaints fetched successfully:", complaints)
    return &complaints
}