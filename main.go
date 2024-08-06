package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"project/authentication"
	"project/course"
	"project/todo"
	"project/utils"
)

func main() {
		EnterCond()
	for {
		fmt.Println("1. Sign Up")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")

		choice := utils.ReadInput("Choose an option: ")

		switch choice {
		case "1":
			fmt.Println("===============Sign Up================")
			username := utils.ReadInput("Enter username: ")
			fmt.Print("Enter password: ")

			// Check if the terminal is available
			if !terminal.IsTerminal(int(os.Stdin.Fd())) {
				fmt.Println("No terminal detected.")
				return
			}

			password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("\nError reading password:", err)
			}
			favorite := utils.ReadInput("\nWhat is your hobby: ")
			address := utils.ReadInput("Enter address: ")
			if err := authentication.SignUp(username, string(password), favorite, address); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Signup successful.")
			}
		case "2":
			fmt.Println("================Login=================")
			username := utils.ReadInput("Enter username: ")
			//password := utils.ReadInput("Enter password: ")
			fmt.Print("Enter password: ")

			// Check if the terminal is available
			if !terminal.IsTerminal(int(os.Stdin.Fd())) {
				fmt.Println("No terminal detected.")
				return
			}

			password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				fmt.Println("\nError reading password:", err)
			}
			if err, fav, address := authentication.Login(username, string(password)); err != nil {
				fmt.Println("Error:Invalid username or password.")
			} else {
				fmt.Println("---------Login successful---------")
				fmt.Println("Name:", username)
				fmt.Println("Your Address:", address)
				fmt.Println("you like ", fav)
				fmt.Println("")
				for {
					fmt.Println("Enter 1 for todo and 2 for courses and 3 to exit")
					choice := utils.ReadInput("Choose an option: ")
					switch choice {
					case "1":
						todoCall(username)
					case "2":
						CourseCall(username)
					case "3":

						os.Exit(0)
					}
				}
			}
		case "3":
			os.Exit(0)

		}
	}
}

func todoCall(username string) {
	for {
		fmt.Println("==============To Do==============")
		fmt.Println("enter 1 to add Todo,2 to list todo ,3 to delete todo and 4 to Go Back")
		choice2 := utils.ReadInput("Choose an option:")
		switch choice2 {
		case "1":
			task := utils.ReadInput("Enter todo task: ")
			if err := todo.AddTodo(username, task); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Todo added.")
			}
		case "2":
			if err := todo.ListTodos(username); err != nil {
				fmt.Println("Error:", err)
			}
		case "3":
			fmt.Println("Enter todo task to delete: ")
			var del int
			fmt.Scanln(&del)
			if err := todo.DeleteTodo(username, del); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Todo deleted.")
			}
		case "4":
			return
		}
	}
}

func CourseCall(username string) {
	for {
		fmt.Println("==========Courses==========")
		fmt.Println("Enter 1 to set course status , 2 to list course status and 3 for overall progress and 4 to Go Back")
		choice3 := utils.ReadInput("Choose an option:")
		switch choice3 {
		case "1":
			coarseName := utils.ReadInput("Enter coarseName:")
			progress := utils.ReadInput("Enter progress:")
			err := course.SetCourseProgress(username, coarseName, progress)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "2":
			err := course.ListCourseProgress(username)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "3":
			fmt.Println("Overall Progress:", course.OverallProgress(username))
		case "4":
			return
		}
	}
}

func EnterCond() {
	answer := utils.ReadInput("Are you a female?[y/n]\n")
	if answer == "y" {
		return
	} else {
		err := errors.New("YOU ARE NOT A VALID USER !")
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}


// readHiddenInput reads input from the user while masking the input with maskChar.
func readHiddenInput(prompt string, maskChar rune) (string, error) {
	fmt.Print(prompt)
	var password []rune
	//stdin := int(os.Stdin.Fd())

	for {
		ch, err := readSingleChar()
		if err != nil {
			return "", err
		}

		switch ch {
		case '\n', '\r':
			fmt.Println() // Move to the next line after Enter key
			return string(password), nil
		case 0x7F: // Handle backspace
			if len(password) > 0 {
				password = password[:len(password)-1]
				fmt.Print("\b \b") // Erase the last masked character
			}
		default:
			if ch >= 32 && ch <= 126 { // Printable ASCII characters
				password = append(password, ch)
				fmt.Print(string(maskChar))
			}
		}
	}
}

// readSingleChar reads a single character from the standard input.
func readSingleChar() (rune, error) {
	buf := make([]byte, 1)
	_, err := os.Stdin.Read(buf)
	if err != nil {
		if err == io.EOF {
			return 0, io.EOF
		}
		return 0, err
	}
	return rune(buf[0]), nil
}

