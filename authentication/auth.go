package authentication

import (
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"project/utils"
	"strings"
)

const userFile = "users.json"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Favorite string `json:"favorite"`
	Address  string `json:"address"`
}

func loadUsers() ([]User, error) {
	file, err := os.Open(userFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []User{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var users []User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return users, nil
}

func saveUsers(users []User) error {
	file, err := os.Create(userFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(users)
}

func SignUp(username, password, favorite, address string) error {
	if !utils.IsPassCorrect(password) {
		return errors.New("password must be at least 8 characters long and must contain atleast one special character and number")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users, err := loadUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return errors.New("username already exists")
		}
	}

	newUser := User{Username: username, Password: string(hashedPassword), Favorite: favorite, Address: address}
	users = append(users, newUser)
	return saveUsers(users)
}

func Login(username, password string) (error, string, string) {
	users, err := loadUsers()
	if err != nil {
		return err, "", ""
	}

	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)), user.Favorite, user.Address
		}
	}

	return errors.New("invalid username or password"), "", ""
}
