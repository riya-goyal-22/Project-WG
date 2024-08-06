package course

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

const courseFile = "courses.json"

type Course struct {
	Username string            `json:"username"`
	Progress map[string]string `json:"progress"` // key: course name, value: progress (e.g., "50%")
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

func SetCourseProgress(username, courseName, progress string) error {
	courses, err := loadCourses()
	if err != nil {
		return err
	}

	var userCourse *Course
	for i, course := range courses {
		if course.Username == username {
			userCourse = &courses[i]
			break
		}
	}

	if userCourse == nil {
		userCourse = &Course{Username: username, Progress: make(map[string]string)}
		courses = append(courses, *userCourse)
	}

	userCourse.Progress[courseName] = progress
	return saveCourses(courses)
}

func ListCourseProgress(username string) error {
	courses, err := loadCourses()
	if err != nil {
		return err
	}

	for _, course := range courses {
		if course.Username == username {
			if len(course.Progress) == 0 {
				fmt.Println("No course progress found.")
				return nil
			}
			fmt.Printf("Course progress for %s:\n", username)
			for courseName, progress := range course.Progress {
				fmt.Printf("%s: %s\n", courseName, progress)
			}
			return nil
		}
	}

	return errors.New("no course progress found for user")
}

func OverallProgress(username string) float64 {
	courses, err := loadCourses()
	if err != nil {
		fmt.Println("Error loading courses")
	}
	var sum int
	mp := make(map[string]string)
	for i, course := range courses {
		if course.Username == username {
			mp = courses[i].Progress
		}
	}
	len := len(mp)
	for _, val := range mp {
		v, _ := strconv.Atoi(val)
		sum += v
	}
	return (float64(sum)) / (float64(len))
}
