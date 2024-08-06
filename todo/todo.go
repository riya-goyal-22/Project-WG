package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

const todoFile = "todos.json"

type Todo struct {
	Username string   `json:"username"`
	Tasks    []string `json:"tasks"`
}

func loadTodos() ([]Todo, error) {
	file, err := os.Open(todoFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var todos []Todo
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&todos)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return todos, nil
}

func saveTodos(todos []Todo) error {
	file, err := os.Create(todoFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(todos)
}

func AddTodo(username, task string) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}

	var userTodo *Todo
	for i, todo := range todos {
		if todo.Username == username {
			userTodo = &todos[i]
			break
		}
	}

	if userTodo == nil {
		userTodo = &Todo{Username: username, Tasks: []string{}}
		//todos = append(todos, *userTodo)
	}

	userTodo.Tasks = append(userTodo.Tasks, task)
	todos = append(todos, *userTodo)
	//fmt.Println(userTodo.Tasks)
	return saveTodos(todos)
}

func ListTodos(username string) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}

	for _, todo := range todos {
		if todo.Username == username {
			if len(todo.Tasks) == 0 {
				fmt.Println("No todos found.")
				return nil
			}
			fmt.Printf("Todos for %s:\n", username)
			for i, task := range todo.Tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}
			return nil
		}
	}

	return errors.New("no todos found for user")
}

func DeleteTodo(username string, index int) error {
	todos, err := loadTodos()
	if err != nil {
		return err
	}

	for i, todo := range todos {
		if todo.Username == username {
			if index < 1 || index > len(todo.Tasks) {
				return errors.New("invalid task number")
			}

			todo.Tasks = append(todo.Tasks[:index-1], todo.Tasks[index:]...)
			todos[i] = todo
			return saveTodos(todos)
		}
	}

	return errors.New("no todos found for user")
}
