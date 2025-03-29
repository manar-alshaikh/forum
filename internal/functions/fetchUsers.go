package functions

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Username    string
	Name        string
	hide_name   bool
	Email       string
	hide_email  bool
	DateOfBirth string
	hide_age    bool
	Hobby       string
	hide_hobby  bool
	Role        string
	profile_picture string
	Blocked_posts int
	banned bool
	Banned_date string
}

func FetchAllUsers(db *sql.DB) *[]User {
	fetchAllUsers := `SELECT username, name, hide_name, email, hide_email, date_of_birth, hide_age, hobby, hide_hobby, role, profile_picture, Blocked_posts, banned, banned_date  FROM users`
	rows, err := db.Query(fetchAllUsers)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Username, &user.Name, &user.hide_name, &user.Email, &user.hide_email, &user.DateOfBirth, &user.hide_age, &user.Hobby, &user.hide_hobby, &user.Role, &user.profile_picture, &user.Blocked_posts, &user.banned, &user.Banned_date)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &users
}

func FetchUserByUsername(db *sql.DB, username string) *User {
	fetchUser:="SELECT username, name, hide_name, email, hide_email, date_of_birth, hide_age, hobby, hide_hobby, role, profile_picture, Blocked_posts, banned, banned_date  FROM users WHERE username = ?"
	var user User
	row := db.QueryRow(fetchUser, username)
	err := row.Scan(&user.Username, &user.Name, &user.hide_name, &user.Email, &user.hide_email, &user.DateOfBirth, &user.hide_age, &user.Hobby, &user.hide_hobby, &user.Role, &user.profile_picture, &user.Blocked_posts, &user.banned, &user.Banned_date)
	if err != nil {
		log.Fatal(err)
	}
	return &user
}

func FetchUserByBlocked_posts(db *sql.DB) *User {
	fetchUser:="SELECT username, name, hide_name, email, hide_email, date_of_birth, hide_age, hobby, hide_hobby, role, profile_picture, Blocked_posts, banned  FROM users WHERE Blocked_posts > 0"
	var user User
	row := db.QueryRow(fetchUser)
	err := row.Scan(&user.Username, &user.Name, &user.hide_name, &user.Email, &user.hide_email, &user.DateOfBirth, &user.hide_age, &user.Hobby, &user.hide_hobby, &user.Role, &user.profile_picture, &user.Blocked_posts, &user.banned, &user.Banned_date)
	if err != nil {
		log.Fatal(err)
	}
	return &user
}

func FetchUserBanned(db *sql.DB) *User {
	fetchUser:="SELECT username, name, hide_name, email, hide_email, date_of_birth, hide_age, hobby, hide_hobby, role, profile_picture, Blocked_posts, banned  FROM users WHERE banned = 1"
	var user User
	row := db.QueryRow(fetchUser)
	err := row.Scan(&user.Username, &user.Name, &user.hide_name, &user.Email, &user.hide_email, &user.DateOfBirth, &user.hide_age, &user.Hobby, &user.hide_hobby, &user.Role, &user.profile_picture, &user.Blocked_posts, &user.banned, &user.Banned_date)
	if err != nil {
		log.Fatal(err)
	}
	return &user
}

func FetchAllUsersDescendingBlockedPosts(db *sql.DB) *[]User {
 
			  fetchAllUsers := `SELECT username, name, hide_name, email, hide_email, date_of_birth, hide_age, hobby, hide_hobby, 
                     role, profile_picture, Blocked_posts, banned, banned_date 
              FROM users 
              ORDER BY Blocked_posts DESC;`
			  rows, err := db.Query(fetchAllUsers)
			  if err != nil {
				  log.Fatal(err)
			  }
			  defer rows.Close()
		  
			  var users []User
			  for rows.Next() {
				  var user User
				  err = rows.Scan(&user.Username, &user.Name, &user.hide_name, &user.Email, &user.hide_email, &user.DateOfBirth, &user.hide_age, &user.Hobby, &user.hide_hobby, &user.Role, &user.profile_picture, &user.Blocked_posts, &user.banned, &user.Banned_date)
				  if err != nil {
					  log.Fatal(err)
				  }
				  users = append(users, user)
			  }
			  err = rows.Err()
			  if err != nil {
				  log.Fatal(err)
			  }
			  return &users
}
