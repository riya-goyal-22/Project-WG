package dailyStatus

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

const dailyStatusFile = "daily_status.json"

type DailyStatus struct {
	Date           string                       `json:"date"`
	UpdatedCourses map[string]map[string]string `json:"deleted_courses"` // map[course ID]map[course name]status
}

func loadDailyStatus() ([]DailyStatus, error) {
	file, err := os.Open(dailyStatusFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []DailyStatus{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var statuses []DailyStatus
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&statuses)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return statuses, nil
}

func saveDailyStatus(statuses []DailyStatus) error {
	file, err := os.Create(dailyStatusFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(statuses)
}

func LogUpdatedCourses(cID string, courseMap map[string]string) error {
	today := time.Now().Format("2006-01-02")
	statuses, err := loadDailyStatus()
	if err != nil {
		return err
	}

	var dailyStatus *DailyStatus
	for i, status := range statuses {
		if status.Date == today {
			dailyStatus = &statuses[i]
			break
		}
	}

	if dailyStatus == nil {
		dailyStatus = &DailyStatus{
			Date:           today,
			UpdatedCourses: make(map[string]map[string]string),
		}
		statuses = append(statuses, *dailyStatus)
	}
	dailyStatus.UpdatedCourses[cID] = courseMap
	return saveDailyStatus(statuses)
}

func ShowDailyStatus() error {
	// Load the daily status
	statuses, err := loadDailyStatus()
	if err != nil {
		return err
	}

	// Check if there are any records
	if len(statuses) == 0 {
		fmt.Println("No daily status records found.")
		return nil
	}

	// Print each daily status record
	fmt.Println("Daily Status Records:")
	fmt.Println("----------------------")
	for _, status := range statuses {
		fmt.Printf("Date: %s\n", status.Date)
		for cID, courses := range status.UpdatedCourses {
			fmt.Printf("Course ID: %s\n", cID)
			for courseName, status := range courses {
				fmt.Printf("  Course Name: %s - Status: %s\n", courseName, status)
			}
		}
		fmt.Println("----------------------")
	}

	return nil
}
