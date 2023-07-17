package validate

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isVaildUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isVaildFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(str string, minLen int, maxLen int) error {
	if len(str) < minLen || len(str) > maxLen {
		return fmt.Errorf("must contain from %d - %d characters", minLen, maxLen)
	}
	return nil
}

func ValidateUsername(username string) error {
	if err := ValidateString(username, 3, 100); err != nil {
		return err
	}
	if !isVaildUsername(username) {
		return fmt.Errorf("must contain only alphanumeric characters and underscore")
	}
	return nil
}

func ValidatePassword(password string) error {
	if err := ValidateString(password, 6, 100); err != nil {
		return err
	}
	return nil
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

func ValidateFullName(fullName string) error {
	if err := ValidateString(fullName, 3, 100); err != nil {
		return err
	}
	if !isVaildFullName(fullName) {
		return fmt.Errorf("must contain only letters and spaces")
	}
	return nil
}

func ValidateEmailId(value int64) error {
	if value <= 0 {
		return fmt.Errorf("invalid email id")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	if err := ValidateString(value, 32, 128); err != nil {
		return err
	}
	return nil
}
