package thesisuc

import (
	"class-reminder-be/database"
	// "class-reminder-be/library"
	"class-reminder-be/model"
	"fmt"
	"strconv"
)

func Show() ([]model.ThesisShowResponse, error) {
	fmt.Println("masuk show")
	// select data from tbl_user_notif
	data, err := database.GetThesis()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)

	var jsonData []model.ThesisShowResponse
	for _, item := range data {
		// Convert byte arrays to strings
		idString := string(item["id"].([]uint8))
		id, _ := strconv.Atoi(idString)
		id64 := int64(id)
		name := string(item["name"].([]uint8))
		nim := string(item["nim"].([]uint8))
		major := string(item["major"].([]uint8))
		// Print the type of each variable
		logbook := string(item["logbook"].([]uint8))
		logbookInt, _ := strconv.Atoi(logbook)
		last_attendance_date := string(item["last_attendance_date"].([]uint8))
		// is_attend_this_week, err := library.IsDateInCurrentWeek(last_attendance_date)
		is_reminder := false
		if err != nil {
			return nil, err
		}
		is_eligable_grad := "Yes"
		if logbookInt < 8 {
			is_reminder = true
			is_eligable_grad = "No"
		}
		// Create a map with string values
		data := model.ThesisShowResponse{
			Id:                   id64,
			Name:                 name,
			Nim:                  nim,
			Logbook:              logbookInt,
			Last_attendance_date: last_attendance_date,
			Major:                major,
			Is_eligable_grad:     is_eligable_grad,
			Is_reminder:          is_reminder,
		}
		jsonData = append(jsonData, data)
	}

	return jsonData, nil

}
