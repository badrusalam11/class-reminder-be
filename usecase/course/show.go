package courseuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/model"
	"fmt"
	"strconv"
)

func Show() ([]model.CourseShowResponse, error) {
	// select data from tbl_user_notif
	data, err := database.GetCourse()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)

	var jsonData []model.CourseShowResponse
	for _, item := range data {
		// Convert byte arrays to strings
		idString := string(item["id"].([]uint8))
		id, _ := strconv.Atoi(idString)
		id64 := int64(id)
		title := string(item["title"].([]uint8))
		description := string(item["description"].([]uint8))
		schedule := string(item["schedule"].([]uint8))
		job_every := library.GetDayFromPrefix(string(item["job_every"].([]uint8)))
		event_day := library.GetDayFromPrefix(string(item["event_day"].([]uint8)))
		event_day_prefix := string(item["event_day"].([]uint8))

		// Create a map with string values
		data := model.CourseShowResponse{
			Id:              id64,
			Title:           title,
			Description:     description,
			Schedule:        schedule,
			CourseDay:       event_day,
			ReminderDay:     job_every,
			CourseDayPrefix: event_day_prefix,
		}
		// data := map[string]interface{}{
		// 	"title": title,
		// 	"id":    id,
		// }

		jsonData = append(jsonData, data)
	}

	return jsonData, nil

}
