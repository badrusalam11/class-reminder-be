package paymentReminderuc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"fmt"
	"strconv"
)

func JobDetail() (model.JobDetailResponse, error) {
	jsonData := model.JobDetailResponse{}
	// select data from job
	item, err := database.JobDetail()
	if err != nil {
		jsonData = model.JobDetailResponse{
			IdEvent:    0,
			Title:      "",
			Schedule:   "",
			IsJobExist: false,
			JobEvery:   "",
			JobName:    "",
			JobId:      "",
			ButtonText: "Create Job Scheduler",
		}
		return jsonData, nil
	}
	fmt.Println(item)

	// Convert byte arrays to strings
	idString := string(item["id"].([]uint8))
	id, _ := strconv.Atoi(idString)
	title := string(item["title"].([]uint8))
	schedule := string(item["schedule"].([]uint8))
	job_every := string(item["job_every"].([]uint8))
	job_name := string(item["job_name"].([]uint8))
	job_id := string(item["job_id"].([]uint8))

	// Create a map with string values
	jsonData = model.JobDetailResponse{
		IdEvent:    id,
		Title:      title,
		Schedule:   schedule,
		IsJobExist: true,
		JobEvery:   job_every,
		JobName:    job_name,
		JobId:      job_id,
		ButtonText: "Edit Job Scheduler",
	}

	return jsonData, nil

}
