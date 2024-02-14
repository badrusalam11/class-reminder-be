package repository

import (
	"bytes"
	"class-reminder-be/config"
	"encoding/json"
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

type Job struct {
	Index     int    `json:"index"`
	Href      string `json:"href"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Group     string `json:"group"`
	Project   string `json:"project"`
	Permalink string `json:"permalink"`
}

type Response struct {
	Succeeded []Job `json:"succeeded"`
	Failed    []Job `json:"failed"`
	Skipped   []Job `json:"skipped"`
}

func CreateJobToRundeck(id int64, title string, schedule string, jobEvery string, optionalParams ...string) (jobId string, err error) {
	uuid := ""
	dupeOption := ""

	// Set optional parameters if provided
	if len(optionalParams) > 0 {
		uuid = optionalParams[0]
	}
	if len(optionalParams) > 1 {
		dupeOption = optionalParams[1]
	}

	jobName := generateJobName(id, title)
	// scheduleCron := generateCronExpression(schedule)
	hr, min, sec := splitSchedule(schedule)
	fmt.Println(jobName)
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
			"uuid": "%s",
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
	 `, uuid, jobName, id, hr, min, sec, jobEvery))
	fmt.Println("payload", string(payload))
	// Create a new HTTP client
	client := &http.Client{}

	// Create a request to create a schedule in Rundeck
	req, err := http.NewRequest("POST", fmt.Sprintf(config.RundeckURL+"/api/45/project/%s/jobs/import?dupeOption=%s", config.RundeckProjectName, dupeOption), bytes.NewBuffer(payload))

	if err != nil {
		return "", err
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
		return "", err
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
		return "", err
	}
	// Check the response status code
	if resp.StatusCode != 200 {
		fmt.Printf("Failed to create schedule. Status code: %d\n", resp.StatusCode)
		fmt.Printf("Response: %s\n", responseBody)
		return "", errors.New("error creating scheduler")
	}

	// Define a variable of the struct type
	var response Response

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	jobId = response.Succeeded[0].ID
	return jobId, err
}

func DeleteJobRundeck(uuid string) (err error) {
	payload := []byte("")
	fmt.Println("payload", string(payload))
	// Create a new HTTP client
	client := &http.Client{}

	// Create a request to create a schedule in Rundeck
	req, err := http.NewRequest("DELETE", fmt.Sprintf(config.RundeckURL+"/api/45/job/%s", uuid), bytes.NewBuffer(payload))

	if err != nil {
		return err
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
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// panic(err)
		return err
	}
	// // Check the response status code
	// if resp.StatusCode != 200 {
	// 	fmt.Printf("Failed to create schedule. Status code: %d\n", resp.StatusCode)
	// 	fmt.Printf("Response: %s\n", responseBody)
	// 	return errors.New("error creating scheduler")
	// }

	return nil
}

func generateJobName(id int64, title string) string {
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
