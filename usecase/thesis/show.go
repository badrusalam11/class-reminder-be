package thesisuc

import (
	"class-reminder-be/database"
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
		supervisor := string(item["supervisor"].([]uint8))
		last_attendance_date := string(item["last_attendance_date"].([]uint8))

		// Create a map with string values
		data := model.ThesisShowResponse{
			Id:                   id64,
			Name:                 name,
			Nim:                  nim,
			Supervisor:           supervisor,
			Last_attendance_date: last_attendance_date,
		}
		jsonData = append(jsonData, data)
	}

	return jsonData, nil

}
