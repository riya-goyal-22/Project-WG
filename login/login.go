// login/login.go
package login

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"project/authentication"
	"strings"
)

func Login(username, password string) error {
	users, err := authentication.LoadUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		}
	}

	return errors.New("invalid username or password")
}
