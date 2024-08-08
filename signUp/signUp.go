package signUp

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"project/authentication"
	"project/utils"
	"strings"
)

func SignUp(username, password, favorite, address string) error {
	if !utils.IsPassCorrect(password) {
		return errors.New("password must be at least 8 characters long and must contain atleast one special character and number")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users, err := authentication.LoadUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return errors.New("username already exists")
		}
	}

	newUser := authentication.User{Username: username, Password: string(hashedPassword), Favorite: favorite, Address: address}
	users = append(users, newUser)
	return authentication.SaveUsers(users)
}
