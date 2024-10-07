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
		is_registered := string(item["is_registered"].([]uint8))
		fmt.Println("is_registered", is_registered)
		is_regis_graduation_int, _ := strconv.Atoi(is_registered)
		// Print the type of each variable
		logbook := string(item["logbook"].([]uint8))
		logbookInt, _ := strconv.Atoi(logbook)
		last_attendance_date := string(item["last_attendance_date"].([]uint8))
		// is_attend_this_week, err := library.IsDateInCurrentWeek(last_attendance_date)
		is_reminder := false
		if err != nil {
			return nil, err
		}
		if logbookInt < 8 {
			is_reminder = true
		}
		is_regis_graduation := false
		is_regis_graduation_string := "No"
		if is_regis_graduation_int == 1 {
			is_regis_graduation = true
			is_regis_graduation_string = "Yes"
		}
		// Create a map with string values
		data := model.ThesisShowResponse{
			Id:                         id64,
			Name:                       name,
			Nim:                        nim,
			Logbook:                    logbookInt,
			Last_attendance_date:       last_attendance_date,
			Major:                      major,
			Is_regis_graduation:        is_regis_graduation,
			Is_regis_graduation_string: is_regis_graduation_string,
			Is_reminder:                is_reminder,
		}
		jsonData = append(jsonData, data)
	}

	return jsonData, nil

}
