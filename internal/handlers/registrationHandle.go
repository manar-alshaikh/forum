package handlers

import (
	"database/sql"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"forum/internal/functions"
	"forum/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	tmpl       *template.Template
	db         *sql.DB
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	storedHash string
	username   string
)

var user struct {
	Username string
	Name     string
	Email    string
}

func InitHandlers(database *sql.DB) {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}
	db = database
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := functions.GetSession(r); err == nil {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := tmpl.ExecuteTemplate(w, "registration.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	formData := models.FormData{
		Username: strings.TrimSpace(r.FormValue("username")),
		Name:     strings.TrimSpace(r.FormValue("name")),
		Email:    strings.TrimSpace(r.FormValue("email")),
		Dob:      r.FormValue("dob"),
		Hobby:    strings.TrimSpace(r.FormValue("hobby")),
	}

	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	file, handler, err := r.FormFile("fileInput")
	var profilePicturePath string
	if err == nil {
		defer file.Close()

		if _, err := os.Stat("static/uploads"); os.IsNotExist(err) {
			os.Mkdir("static/uploads", 0755)
		}

		ext := filepath.Ext(handler.Filename)
		profilePicturePath = "uploads/" + uuid.New().String() + ext

		dst, err := os.Create("static/" + profilePicturePath)
		if err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
	} else if err != http.ErrMissingFile {
		http.Error(w, "Error processing image", http.StatusBadRequest)
		return
	}

	functions.ValidateUsername(formData.Username, &formData)
	functions.ValidateName(formData.Name, &formData)
	functions.ValidateEmail(formData.Email, &formData)
	functions.ValidateDob(formData.Dob, &formData)
	functions.ValidatePassword(password, confirmPassword, &formData)

	if functions.HasErrors(&formData) {
		err := tmpl.ExecuteTemplate(w, "registration.html", formData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	_, err = db.Exec(`INSERT INTO users 
	(username, name, email, date_of_birth, hobby, password, profile_picture) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`,
		formData.Username, formData.Name, formData.Email,
		formData.Dob, formData.Hobby, hashedPassword, profilePicturePath)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	functions.CreateSession(formData.Username, w)
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	loginInput  := strings.TrimSpace(r.FormValue("loginInput"))
	password := r.FormValue("password")

	// Determine if the input is an email or username
	var query string
	var queryArg string
	if emailRegex.MatchString(loginInput) {
		query = "SELECT username, password FROM users WHERE email = ?"
		queryArg = loginInput
	} else {
		query = "SELECT username, password FROM users WHERE username = ?"
		queryArg = loginInput
	}

	err = db.QueryRow(query, queryArg).Scan(&username, &storedHash)

	if err == sql.ErrNoRows {
		err := tmpl.ExecuteTemplate(w, "registration.html", models.FormData{
			LoginEmailError: "No account found with this Username / Email",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	} else if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		err := tmpl.ExecuteTemplate(w, "registration.html", models.FormData{
			LoginPasswordError: "Incorrect password",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	functions.CreateSession(username, w)
	log.Println("Redirecting to dashboard...")
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, err := functions.GetSession(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err = db.QueryRow("SELECT username, name, email FROM users WHERE username = ?", session.Username).Scan(
		&user.Username, &user.Name, &user.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "dashboard.html", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	functions.DeleteSession(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
