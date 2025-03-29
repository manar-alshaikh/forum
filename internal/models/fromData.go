package models

type FormData struct {
	Username                  string
	// RegistrationUsername    string   // For registration username
	// LoginUsername           string   // For login username
	Name                      string
	Email                     string
	Dob                       string
	Hobby                     string
	UsernameError             string
	NameError                 string
	RegistrationEmailError    string 
	LoginEmailError           string 
	DobError                  string
	RegistrationPasswordError string 
	LoginPasswordError        string 
	ConfirmPasswordError      string
	View                      string
}
