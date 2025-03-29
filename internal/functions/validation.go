package functions

import (
	"database/sql"
	"log"
	"regexp"
	"strings"
	"forum/internal/models"
	"time"
	"unicode"
)

var (
	db            *sql.DB
	emailRegex    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]{4,20}$`)
)

// InitDB initializes the package with a database connection
func InitDB(database *sql.DB) {
	db = database
}

func ValidateUsername(username string, formData *models.FormData) {
	if !usernameRegex.MatchString(username) {
		formData.UsernameError = "Must be 4-20 characters (letters, numbers, underscores only)"
		return
	}

	var exists bool
	if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists); err != nil {
		log.Printf("Username check error: %v", err)
		formData.UsernameError = "System error checking username availability"
		return
	}
	if exists {
		formData.UsernameError = "Username already taken"
	}
}

func ValidateName(name string, formData *models.FormData) {
	if len(name) < 2 {
		formData.NameError = "Must be at least 2 characters"
	}
}

func ValidateEmail(email string, formData *models.FormData) {
	if !emailRegex.MatchString(email) {
		formData.RegistrationEmailError = "Please enter a valid email address"
		return
	}

	var exists bool
	if err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = ?)", email).Scan(&exists); err != nil {
		log.Printf("Email check error: %v", err)
		formData.RegistrationEmailError = "System error checking email availability"
		return
	}
	if exists {
		formData.RegistrationEmailError = "Email already registered"
	}
}

func ValidateDob(dob string, formData *models.FormData) {
	parsedDob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		formData.DobError = "Invalid date format (YYYY-MM-DD)"
		return
	}

	age := time.Since(parsedDob).Hours() / 24 / 365
	if age < 18 {
		formData.DobError = "You must be at least 18 years old"
	}
}

func ValidatePassword(password, confirmPassword string, formData *models.FormData) {
	var requirements []string

	if len(password) < 8 {
		requirements = append(requirements, "at least 8 characters")
	}

	var (
		hasNumber   = false
		hasLower    = false
		hasUpper    = false
		hasSpecial  = false
		letterCount = 0
	)

	for _, char := range password {
		switch {
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsLower(char):
			hasLower = true
			letterCount++
		case unicode.IsUpper(char):
			hasUpper = true
			letterCount++
		case strings.ContainsRune("!@#$%^&*()-_=+{};:,<.>", char):
			hasSpecial = true
		}
	}

	if !hasNumber {
		requirements = append(requirements, "one number")
	}
	if !hasLower {
		requirements = append(requirements, "one lowercase letter")
	}
	if !hasUpper {
		requirements = append(requirements, "one uppercase letter")
	}
	if !hasSpecial {
		requirements = append(requirements, "one special character")
	}
	if letterCount < 5 {
		requirements = append(requirements, "at least five letters")
	}

	if len(requirements) > 0 {
		formData.RegistrationPasswordError = "Password must contain: " + strings.Join(requirements, ", ")
	}

	if password != confirmPassword {
		formData.ConfirmPasswordError = "Passwords do not match"
	}
}

func HasErrors(formData *models.FormData) bool {
	return formData.UsernameError != "" ||
		formData.NameError != "" ||
		formData.RegistrationEmailError != "" ||
		formData.DobError != "" ||
		formData.RegistrationPasswordError != "" ||
		formData.ConfirmPasswordError != ""
}
