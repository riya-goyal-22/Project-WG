package profile

import (
	"fmt"
	"project/authentication"
	"strings"
)

func Profile(username string) {
	users, err := authentication.LoadUsers()
	if err != nil {
		fmt.Println("Error loading users: ", err)
	}
	var userprofile authentication.User
	for _, user := range users {
		if strings.EqualFold(user.Username, username) {
			userprofile = user
		}
	}
	fmt.Println("=====================================")
	fmt.Println("Your Profile")
	fmt.Println("=====================================")
	fmt.Println("Welcome ", userprofile.Username)
	fmt.Println("You live at ", userprofile.Address)
	fmt.Println("You like to ", userprofile.Favorite)
	fmt.Println("")
}
