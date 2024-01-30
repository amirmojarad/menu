package service

import (
	"regexp"
)

type Validator interface {
	Validate() bool
}

func ValidateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

func ValidatePassword(password string) bool {
	passwordRegex := `^(?=.*[A-Z])(?=.*[a-z])(?=.*\d).+$`
	return regexp.MustCompile(passwordRegex).MatchString(password)
}

func (l LoginRequest) Validate() bool {
	return ValidateEmail(l.Email) && ValidatePassword(l.Password)
}

func (s SignupRequest) Validate() bool {
	return ValidateEmail(s.Email) || ValidatePassword(s.Password)
}
