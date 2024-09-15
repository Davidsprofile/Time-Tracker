package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

// struct to store task name and duration
type Session struct {
	TaskName string
	Duration time.Duration
}

// starttimer, starts the timer for the session
func StartTimer() time.Time {
	fmt.Println("Timer started")
	return time.Now()
}

// stoptimer, stop the timer for the session
func StopTimer(start time.Time) time.Duration {
	elapsed := time.Since(start) // calculate time difference
	fmt.Printf("Timer stoped, Duration %v\n", elapsed)
	return elapsed
}

// Save session to the CSV file
func SaveSession(taskname string, duration time.Duration) {
	file, err := os.OpenFile("sessions.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// write session details to the csv files
	err = writer.Write([]string{taskname, duration.String()})
	if err != nil {
		fmt.Println("Error writing to a file", err)
	}
}

// reads and displays all the recorded sessions from the CSV file
func ViewSessions() {
	file, err := os.Open("sessions.csv")
	if err != nil {
		fmt.Println("Could not open File", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Could not read file", err)
		return
	}
	for _, record := range records {
		fmt.Printf("Task: %s, Duration: %s\n", record[0], record[1])
	}
}

func main() {
	var choice int
	var taskName string
	var startTime time.Time

	for {
		fmt.Println("1. Start Timer")
		fmt.Println("2. Stop Timer")
		fmt.Println("3. View Sessions")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter the task name: ")
			fmt.Scan(&taskName)
			startTime = StartTimer()
		case 2:
			duration := StopTimer(startTime)
			SaveSession(taskName, duration)
		case 3:
			ViewSessions()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
