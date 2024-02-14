package courseuc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"class-reminder-be/repository"
	"fmt"
)

func Delete(course model.CourseDeleteRequest) error {

	// delete data from tbl_event, tbl_job, tbl_user_event
	err := database.DeleteCourse(course.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// select data from tbl_job
	data, err := database.GetJob(course.Id)
	if data == nil {
		fmt.Println("no data in tbl job")
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	uuid := string(data["job_id"].([]uint8))
	err = repository.DeleteJobRundeck(uuid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
