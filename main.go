package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"project/course"
	"project/dailyStatus"
	"project/login"
	"project/profile"
	"project/signUp"
	"project/todo"
	"project/utils"
	"strings"
)

func init() {
	EnterCond()
}
func main() {
	fmt.Println("===================================")
	fmt.Println("\tWelcome")
	fmt.Println("===================================")
	for {
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println("Choose an option :")
		fmt.Println("1. Sign Up")
		fmt.Println("2. Login")
		fmt.Println("3. Exit")

		choice := utils.ReadInput("Your choice : ")

		switch choice {
		case "1":
			fmt.Println("\n===============Sign Up================")
			username := utils.ReadInput("Enter username: ")
			fmt.Print("[\"password must be at least 8 characters long and must contain atleast one special character and number\"]\nEnter password: ")

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
			if err := signUp.SignUp(username, string(password), favorite, address); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("\nSignup successful.")
			}
		case "2":
			fmt.Println("\n================Login=================")
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
			if err := login.Login(username, string(password)); err != nil {
				fmt.Println("Error:Invalid username or password.")
			} else {
				fmt.Println("\n---------Login successful---------")
				for {
					fmt.Println(strings.Repeat("-", 30))
					fmt.Println("Enter your choice:")
					fmt.Println("1. Manage Todos")
					fmt.Println("2. Manage Courses")
					fmt.Println("3. Daily Status")
					fmt.Println("4. View Profile")
					fmt.Println("5. Exit")
					choice := utils.ReadInput("\nYour choice: ")
					switch choice {
					case "1":
						todoCall(username)
					case "2":
						CourseCall(username)
					case "3":
						err := dailyStatus.ShowDailyStatus()
						if err != nil {
							fmt.Println("\nError showing daily status:", err)
						}
					case "4":
						profile.Profile(username)
					case "5":
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
		fmt.Println("\n==============To Do==============")
		fmt.Println("Enter your choice :")
		fmt.Println("1. Add todo")
		fmt.Println("2. List todo")
		fmt.Println("3. Delete todo")
		fmt.Println("4. Go Back")
		choice2 := utils.ReadInput("\nYour choice:")
		switch choice2 {
		case "1":
			task := utils.ReadInput("\nEnter todo task: ")
			if err := todo.AddTodo(username, task); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("\nTodo added.")
			}
		case "2":
			if err := todo.ListTodos(username); err != nil {
				fmt.Println("Error:", err)
			}
		case "3":
			if err := todo.ListTodos(username); err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println("\nEnter todo task to delete: ")
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
		fmt.Println("\n==========Courses==========")
		fmt.Println("Enter your choice:")
		fmt.Println("1. Set courses")
		fmt.Println("2. List courses")
		fmt.Println("3. Update courses")
		fmt.Println("4. Delete courses")
		fmt.Println("5. Overall Progress")
		fmt.Println("6. Go Back")
		choice3 := utils.ReadInput("\nYour choice:")
		switch choice3 {
		case "1":
			coarseName := utils.ReadInput("\nEnter coarseName :")
			progress := utils.ReadInput("Enter progress [done/pending] :")
			cId := utils.ReadInput("Enter course id :")
			err := course.SetCourseProgress(username, coarseName, progress, cId)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "2":
			err := course.ListCourseProgress(username)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "3":
			err := course.ListCourseProgress(username)
			if err != nil {
				fmt.Println("Error:", err)
			}
			cId := utils.ReadInput("\nEnter coarseId :")
			progress := utils.ReadInput("Enter progress [done/pending]:")
			err = course.UpdateCourse(username, cId, progress)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("\nCourse updated.")
			}
			err = course.ListCourseProgress(username)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "4":
			err := course.ListCourseProgress(username)
			if err != nil {
				fmt.Println("Error:", err)
			}
			cId := utils.ReadInput("\nEnter coarseId :")
			err = course.DeleteCourse(username, cId)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Course deleted.")
			}
		case "5":
			progress, err := course.OverallProgress(username)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println("\nOverall Progress:", progress, "%")
		case "6":
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
			os.Exit(0)
		}
	}
}
