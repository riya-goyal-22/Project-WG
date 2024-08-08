package authentication

import (
	"encoding/json"
	"io"
	"os"
)

const userFile = "users.json"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Favorite string `json:"favorite"`
	Address  string `json:"address"`
}

func LoadUsers() ([]User, error) {
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

func SaveUsers(users []User) error {
	file, err := os.Create(userFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(users)
}
