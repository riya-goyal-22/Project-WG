package course

import (
	"os"
	"strings"
	"testing"
)

const testCourseFile = "test_courses.json"

// Helper function to setup a clean test environment
func setup() {
	os.Remove(testCourseFile)
}

// Helper function to cleanup after tests
func teardown() {
	os.Remove(testCourseFile)
}

func TestSetCourseProgress(t *testing.T) {
	setup()
	defer teardown()

	err := SetCourseProgress("user1", "Math", "In Progress", "C1")
	if err != nil {
		t.Fatalf("Failed to set course progress: %v", err)
	}

	courses, err := loadCourses()
	if err != nil {
		t.Fatalf("Failed to load courses: %v", err)
	}

	if len(courses) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(courses))
	}

	if len(courses[0].Progress) != 1 {
		t.Fatalf("Expected 1 course ID, got %d", len(courses[0].Progress))
	}

	if courses[0].Progress["C1"]["Math"] != "In Progress" {
		t.Fatalf("Expected progress 'In Progress', got %s", courses[0].Progress["C1"]["Math"])
	}
}

func TestUpdateCourse(t *testing.T) {
	setup()
	defer teardown()

	err := SetCourseProgress("user1", "Math", "In Progress", "C1")
	if err != nil {
		t.Fatalf("Failed to set course progress: %v", err)
	}

	err = UpdateCourse("user1", "C1", "Done")
	if err != nil {
		t.Fatalf("Failed to update course progress: %v", err)
	}

	courses, err := loadCourses()
	if err != nil {
		t.Fatalf("Failed to load courses: %v", err)
	}

	if len(courses) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(courses))
	}

	if courses[0].Progress["C1"]["Math"] != "Done" {
		t.Fatalf("Expected progress 'Done', got %s", courses[0].Progress["C1"]["Math"])
	}
}

func TestDeleteCourse(t *testing.T) {
	setup()
	defer teardown()

	err := SetCourseProgress("user1", "Math", "In Progress", "C1")
	if err != nil {
		t.Fatalf("Failed to set course progress: %v", err)
	}

	err = DeleteCourse("user1", "C1")
	if err != nil {
		t.Fatalf("Failed to delete course: %v", err)
	}

	courses, err := loadCourses()
	if err != nil {
		t.Fatalf("Failed to load courses: %v", err)
	}

	if len(courses) != 0 {
		t.Fatalf("Expected 0 courses, got %d", len(courses))
	}
}

func TestOverallProgress(t *testing.T) {
	setup()
	defer teardown()

	err := SetCourseProgress("user1", "Math", "Done", "C1")
	if err != nil {
		t.Fatalf("Failed to set course progress: %v", err)
	}

	err = SetCourseProgress("user1", "Science", "In Progress", "C2")
	if err != nil {
		t.Fatalf("Failed to set course progress: %v", err)
	}

	progress, err := OverallProgress("user1")
	if err != nil {
		t.Fatalf("Failed to calculate overall progress: %v", err)
	}

	if progress != 50 {
		t.Fatalf("Expected overall progress 50, got %f", progress)
	}
}

// Helper function to check if a string contains a substring
func contains(str, substr string) bool {
	return strings.Contains(str, substr)
}
