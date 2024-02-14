package eventuc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"class-reminder-be/repository"
	"fmt"
	"strings"

	"encoding/json"
)

func Create(event model.EventCreateRequest) (string, error) {
	jobEveryMap := make(map[string]bool)
	// compose job every
	for i := 0; i < len(event.JobEvery); i++ {
		jobEveryMap[event.JobEvery[i]] = true
	}
	jobJson, _ := json.Marshal(jobEveryMap)
	jobJsonString := string(jobJson)
	// insert to tbl_event
	id, err := database.InsertEventToDB(event, jobJsonString)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(id)

	// create job to rundeck
	jobString := strings.Join(event.JobEvery, ",")
	jobId, err := repository.CreateJobToRundeck(id, event.Title, event.Schedule, jobString)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(jobId)
	return "success", nil
}
