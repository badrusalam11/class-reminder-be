package courseuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/model"
	"fmt"
)

func Detail(course model.CourseDeleteRequest) (result model.CourseShowResponse, err error) {
	// select data from tbl_user_notif
	item, err := database.GetCourseById(course.Id)
	if err != nil {
		return result, err
	}
	fmt.Println(item)

	// Convert byte arrays to strings
	id := item["id"].(int64)
	title := string(item["title"].([]uint8))
	description := string(item["description"].([]uint8))
	schedule := string(item["schedule"].([]uint8))
	job_every := library.GetDayFromPrefix(string(item["job_every"].([]uint8)))
	event_day := library.GetDayFromPrefix(string(item["event_day"].([]uint8)))
	event_day_prefix := string(item["event_day"].([]uint8))

	// Create a map with string values
	data := model.CourseShowResponse{
		Id:              id,
		Title:           title,
		Description:     description,
		Schedule:        schedule,
		CourseDay:       event_day,
		ReminderDay:     job_every,
		CourseDayPrefix: event_day_prefix,
	}

	return data, nil

}
