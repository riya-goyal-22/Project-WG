// login/login_test.go
package login

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func TestLogin_InvalidPassword(t *testing.T) {
	password, _ := hashPassword("password123")
	err := Login("testUser", password)
	if err == nil {
		t.Fatal("expected error, got none")
	}
	if err.Error() != "invalid username or password" {
		t.Fatalf("expected 'invalid username or password' error, got %v", err)
	}
}

func TestLogin_InvalidUsername(t *testing.T) {
	password, _ := hashPassword("password123")
	err := Login("wrongUser", password)
	if err == nil {
		t.Fatal("expected error, got none")
	}
}
