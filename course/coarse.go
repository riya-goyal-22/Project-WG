package course

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"project/dailyStatus"
	"strings"
)

const courseFile = "courses.json"

type Course struct {
	Username string                       `json:"username"`
	Progress map[string]map[string]string `json:"progress"`
}

func loadCourses() ([]Course, error) {
	file, err := os.Open(courseFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Course{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var courses []Course
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&courses)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return courses, nil
}

func saveCourses(courses []Course) error {
	file, err := os.Create(courseFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(courses)
}

func SetCourseProgress(username, courseName, progress, cID string) error {
	// Load existing courses
	courses, err := loadCourses()
	if err != nil {
		return err
	}

	// Find the user's course entry
	var userCourse *Course
	for i, course := range courses {
		if course.Username == username {
			userCourse = &courses[i]
			break
		}
	}

	// If no course entry for the user is found, create a new one
	if userCourse == nil {
		userCourse = &Course{
			Username: username,
			Progress: make(map[string]map[string]string),
		}
		courses = append(courses, *userCourse)
	}

	// Initialize the inner map if it doesn't exist
	if _, exists := userCourse.Progress[cID]; !exists {
		userCourse.Progress[cID] = make(map[string]string)
	}

	// Set the course progress
	userCourse.Progress[cID][courseName] = progress

	// Save updated courses
	return saveCourses(courses)
}

func ListCourseProgress(username string) error {
	courses, err := loadCourses()
	if err != nil {
		return err
	}

	var userCourse *Course
	for _, course := range courses {
		if course.Username == username {
			userCourse = &course
			break
		}
	}

	if userCourse == nil {
		fmt.Println("\nNo course progress found for user.")
		return nil
	}

	if len(userCourse.Progress) == 0 {
		fmt.Println("\nNo course progress found.")
		return nil
	}

	fmt.Printf("\nCourse progress for %s:\n", username)
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("%-15s | %-25s | %s\n", "CourseID", "Course Name", "Course Progress")
	fmt.Println(strings.Repeat("-", 60))

	for cID, courseMap := range userCourse.Progress {
		for courseName, progress := range courseMap {
			fmt.Printf("%-15s | %-25s | %s\n", cID, courseName, progress)
		}
	}

	fmt.Println(strings.Repeat("-", 60))
	return nil
}

func UpdateCourse(username, cID, progress string) error {
	// Load existing courses
	courses, err := loadCourses()
	if err != nil {
		return err
	}

	// Find the user's course entry
	var userCourse *Course
	for i, course := range courses {
		if course.Username == username {
			userCourse = &courses[i]
			break
		}
	}

	// If no course entry for the user is found, return an error
	if userCourse == nil {
		return errors.New("no courses found for user")
	}

	// Check if the course ID exists
	if _, exists := userCourse.Progress[cID]; !exists {
		return errors.New("course ID not found")
	}

	// Prepare to log updated courses
	updatedCourses := make(map[string]string)
	for courseName, currentStatus := range userCourse.Progress[cID] {
		if currentStatus != progress {
			updatedCourses[courseName] = progress
		}
	}

	// Update all course names under the given course ID with the new progress
	for courseName := range userCourse.Progress[cID] {
		userCourse.Progress[cID][courseName] = progress
	}

	// Save updated courses
	if err := saveCourses(courses); err != nil {
		return err
	}

	// Log updated courses
	if len(updatedCourses) > 0 {
		err := dailyStatus.LogUpdatedCourses(cID, updatedCourses)
		if err != nil {
			return err
		}
	}

	return nil
}
func DeleteCourse(username, cID string) error {
	// Load the existing courses
	courses, err := loadCourses()
	if err != nil {
		return err
	}

	// Find the user's course entry
	var userCourse *Course
	for i, course := range courses {
		if course.Username == username {
			userCourse = &courses[i]
			break
		}
	}

	// If no course entry for the user is found, return an error
	if userCourse == nil {
		return errors.New("no course found for user")
	}

	// Remove the specified course
	if _, exists := userCourse.Progress[cID]; exists {
		delete(userCourse.Progress, cID)
	} else {
		return errors.New("course ID not found")
	}

	// Remove user entry if no courses remain
	if len(userCourse.Progress) == 0 {
		for i := range courses {
			if courses[i].Username == username {
				courses = append(courses[:i], courses[i+1:]...)
				break
			}
		}
	}

	// Save updated courses
	if err := saveCourses(courses); err != nil {
		return err
	}

	return nil
}

func OverallProgress(username string) (float64, error) {
	// Load the existing courses
	courses, err := loadCourses()
	if err != nil {
		return 0, err
	}

	// Find the user's course entry
	var userCourse *Course
	for _, course := range courses {
		if course.Username == username {
			userCourse = &course
			break
		}
	}

	// If no course entry for the user is found, return an error
	if userCourse == nil {
		return 0, errors.New("no course progress found for user")
	}

	// Calculate the total progress
	totalCourses := 0
	completedCourses := 0

	for _, courseMap := range userCourse.Progress {
		for _, status := range courseMap {
			totalCourses++
			if status == "Done" || status == "done" {
				completedCourses++
			}
		}
	}

	// Handle division by zero if no courses are present
	if totalCourses == 0 {
		return 0, nil
	}

	// Calculate the percentage
	percentage := (float64(completedCourses) / float64(totalCourses)) * 100
	return percentage, nil
}
