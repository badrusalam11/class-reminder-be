package courseuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/model"
	"class-reminder-be/repository"
	"fmt"
)

func Edit(course model.CourseEditRequest) error {
	job_day := library.GetJobDay(course.Day)
	// update data to tbl_event
	err := database.UpdateCourseToDB(course.Id, course.Title, course.Description, course.Schedule, course.Day, job_day)
	if err != nil {
		fmt.Println(err)
		return err
	}
	// select from tbl_job
	data, err := database.GetJob(course.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	uuid := string(data["job_id"].([]uint8))
	jobId, err := repository.CreateJobToRundeck(course.Id, course.Title, course.Schedule, job_day, uuid, "update")
	if err != nil {
		fmt.Println(err)
		return err
	}

	jobName := library.GenerateJobName(course.Id, course.Title)

	err = database.UpdateJob(jobName, jobId, course.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
