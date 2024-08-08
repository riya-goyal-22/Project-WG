package todo

import (
	"os"
	"strings"
	"testing"
)

const testTodoFile = "test_todos.json"

// Helper function to setup a clean test environment
func setup() {
	err := os.Remove(testTodoFile)
	if err != nil {
		return
	}
}

// Helper function to cleanup after tests
func teardown() {
	err := os.Remove(testTodoFile)
	if err != nil {
		return
	}
}

func TestAddTodo(t *testing.T) {
	setup()
	defer teardown()

	err := AddTodo("user1", "Task 1")
	if err != nil {
		t.Fatalf("Failed to add todo: %v", err)
	}

	err = AddTodo("user1", "Task 2")
	if err != nil {
		t.Fatalf("Failed to add todo: %v", err)
	}

	todos, err := loadTodos()
	if err != nil {
		t.Fatalf("Failed to load todos: %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(todos))
	}

	if len(todos[0].Tasks) != 2 {
		t.Fatalf("Expected 2 tasks, got %d", len(todos[0].Tasks))
	}

	if todos[0].Tasks[0] != "Task 1" || todos[0].Tasks[1] != "Task 2" {
		t.Fatalf("Tasks do not match expected values")
	}
}

func TestDeleteTodo(t *testing.T) {
	setup()
	defer teardown()

	err := AddTodo("user1", "Task 1")
	if err != nil {
		t.Fatalf("Failed to add todo: %v", err)
	}

	err = AddTodo("user1", "Task 2")
	if err != nil {
		t.Fatalf("Failed to add todo: %v", err)
	}

	err = DeleteTodo("user1", 1)
	if err != nil {
		t.Fatalf("Failed to delete todo: %v", err)
	}

	todos, err := loadTodos()
	if err != nil {
		t.Fatalf("Failed to load todos: %v", err)
	}

	if len(todos) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(todos))
	}

	if len(todos[0].Tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(todos[0].Tasks))
	}

	if todos[0].Tasks[0] != "Task 2" {
		t.Fatalf("Task 2 was not found after deletion")
	}
}

// Helper function to check if a string contains a substring
func contains(str, substr string) bool {
	return strings.Contains(str, substr)
}
