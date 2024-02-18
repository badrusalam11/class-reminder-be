package eventuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/model"
	"class-reminder-be/repository"
	"fmt"
)

func Edit(event model.EventEditRequest) (string, error) {
	// jobEveryMap := make(map[string]bool)
	// // compose job every
	// for i := 0; i < len(event.JobEvery); i++ {
	// 	jobEveryMap[event.JobEvery[i]] = true
	// }
	// jobJson, _ := json.Marshal(jobEveryMap)
	// jobJsonString := string(jobJson)
	// insert to tbl_event
	err := database.EditEventToDB(event, event.JobEvery)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	// select job
	data, err := database.GetJob(int64(event.Id))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	uuid := string(data["job_id"].([]uint8))

	// create job to rundeck
	// jobString := strings.Join(event.JobEvery, ",")
	jobId := ""
	if event.IdEventType != 3 {
		jobId, err = repository.CreateJobToRundeck(int64(event.Id), event.Title, event.Schedule, event.JobEvery)
	} else {
		jobId, err = repository.CreateJobToRundeck(int64(event.Id), event.Title, event.Schedule, "", uuid, "update", event.JobEvery)
	}
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(jobId)
	jobName := library.GenerateJobName(int64(event.Id), event.Title)
	err = database.UpdateJob(jobName, jobId, int64(event.Id))
	if err != nil {
		return "", err
	}
	return "success", nil
}
