package eventuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/model"
	"class-reminder-be/repository"
	"fmt"
)

func Create(event model.EventCreateRequest) (string, error) {
	// jobEveryMap := make(map[string]bool)
	// // compose job every
	// for i := 0; i < len(event.JobEvery); i++ {
	// 	jobEveryMap[event.JobEvery[i]] = true
	// }
	// jobJson, _ := json.Marshal(jobEveryMap)
	// jobJsonString := string(jobJson)
	// insert to tbl_event
	id, err := database.InsertEventToDB(event, event.JobEvery)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(id)

	// create job to rundeck
	// jobString := strings.Join(event.JobEvery, ",")
	jobId := ""
	if event.IdEventType != 3 {
		jobId, err = repository.CreateJobToRundeck(id, event.Title, event.Schedule, event.JobEvery)
	} else {
		jobId, err = repository.CreateJobToRundeck(id, event.Title, event.Schedule, "", "", "", event.JobEvery)
	}
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(jobId)
	jobName := library.GenerateJobName(id, event.Title)
	err = database.InsertToTableJob(jobName, jobId, id)
	if err != nil {
		return "", err
	}
	return "success", nil
}
