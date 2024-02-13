package repository

import (
	"bytes"
	"class-reminder-be/config"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// const (
// 	RundeckURL   = "http://localhost:4440"
// 	APIToken     = "wfXV0ceZ3ejCufNZrEKPlEFeypq9C0pH"
// 	ProjectName  = "class_reminder"
// 	JobID        = "your-job-id"
// 	ScheduleName = "your-schedule-name"
// 	ScheduleCron = "0 0 * * * ?" // Example: Run daily at midnight
// )

func CreateJobToRundeck(id int64, title string, schedule string, jobEvery string) error {
	jobId := generateJobId(id, title)
	// scheduleCron := generateCronExpression(schedule)
	hr, min, sec := splitSchedule(schedule)
	fmt.Println(jobId)
	// fmt.Println(scheduleCron)
	// Create a JSON payload for the schedule request
	// payload := []byte(fmt.Sprintf(`{
	//     "jobId": "%s",
	//     "name": "%s",
	//     "group": "default",
	//     "schedule": "R1/%s",
	//     "scheduleTimeZone": "UTC"
	// }`, jobId, jobId, scheduleCron))
	payload := []byte(fmt.Sprintf(`[
		{
			"name": "%s",
			"description": "",
			"loglevel": "INFO",
			"sequence": {
				"commands": [
					{
						"script": "curl --location \"http://localhost:9090/notif/send\" ^ --header \"Content-Type: application/json\" ^ --data \"{ \\\"event_id\\\": %d }\""
					}
				]
			},
			"schedule": {
				"time": {
					"hour": "%s",
					"minute": "%s",
					"seconds": "%s"
				},
				"month": "*",
				"year": "*",
				"weekday": {
					"day": "%s"
				}
			}
		}
	]
	 `, jobId, id, hr, min, sec, jobEvery))
	fmt.Println("payload", string(payload))
	// Create a new HTTP client
	client := &http.Client{}

	// Create a request to create a schedule in Rundeck
	req, err := http.NewRequest("POST", fmt.Sprintf(config.RundeckURL+"/api/45/project/%s/jobs/import", config.RundeckProjectName), bytes.NewBuffer(payload))

	if err != nil {
		return err
		// panic(err)
	}

	// Set headers for the request
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Rundeck-Auth-Token", config.RundeckToken)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		// panic(err)
		return err
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
		return err
	}

	// Check the response status code
	if resp.StatusCode != 200 {
		fmt.Printf("Failed to create schedule. Status code: %d\n", resp.StatusCode)
		fmt.Printf("Response: %s\n", responseBody)
		return errors.New("error creating scheduler")
	}

	return nil
}

func generateJobId(id int64, title string) string {
	idString := strconv.FormatInt(id, 10)
	titleString := strings.ReplaceAll(title, " ", "_")
	// Get the current time
	// currentTime := time.Now()
	// // Format the time as DDMMYYYY
	// formattedDate := currentTime.Format("02012006")
	// job := idString + ":" + titleString + "_" + formattedDate
	job := idString + ":" + titleString
	return job
}

func splitSchedule(input string) (string, string, string) {
	timeComponents := strings.Split(input, ":")

	if len(timeComponents) == 3 {
		return timeComponents[0], timeComponents[1], timeComponents[2] // Set default

	} else if len(timeComponents) == 2 {

		return timeComponents[0], timeComponents[1], "00" // Set default
	} else {
		return "09", "00", "00"
	}
}

func generateCronExpression(input string) string {
	// Split the input string into hours, minutes, and seconds
	timeComponents := strings.Split(input, ":")

	if len(timeComponents) != 2 {
		return "" // Invalid input
	}

	hours := timeComponents[0]
	minutes := timeComponents[1]
	// seconds := timeComponents[2]

	// Generate the cron expression
	cronExpression := fmt.Sprintf("%s %s * * *", minutes, hours)

	return cronExpression
}
