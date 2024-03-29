package courseuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/model"
	"class-reminder-be/repository"
	"fmt"
)

func Create(course model.CourseCrateRequest) error {
	job_day := library.GetJobDay(course.Day)
	// select data from tbl_user_notif
	id, err := database.InsertCourseToDB(course.Title, course.Description, course.Schedule, course.Day, job_day)
	if err != nil {
		fmt.Println(err)
		return err
	}
	jobId, err := repository.CreateJobToRundeck(id, course.Title, course.Schedule, job_day)
	if err != nil {
		fmt.Println(err)
		return err
	}

	jobName := library.GenerateJobName(id, course.Title)

	err = database.InsertToTableJob(jobName, jobId, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
