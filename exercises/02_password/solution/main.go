package main

import (
	"fmt"
	"unicode"
)

func ValidatePassword(password string) bool {

	// Password must be at least 8 characters long
	if len(password) < 8 {
		return false
	}

	// Password must contain at least one uppercase letter, one lowercase letter, and one digit
	var hasUpper, hasLower, hasDigit bool
	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}
	}
	if !hasUpper || !hasLower || !hasDigit {
		return false
	}

	// Password is valid
	return true
}

func main() {
	// Prompt the user to input a password
	var password string
	fmt.Print("Enter a password: ")
	fmt.Scanln(&password)

	// Validate the password
	isValid := ValidatePassword(password)
	if isValid {
		fmt.Println("Password is valid")
	} else {
		fmt.Println("Password is invalid")
	}
}
