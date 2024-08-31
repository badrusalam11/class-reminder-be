package graduationuc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"fmt"
	"strconv"
)

func Show() ([]model.GraduationShowResponse, error) {
	// select data from tbl_user_notif
	data, err := database.GetGraduation()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)

	var jsonData []model.GraduationShowResponse
	for _, item := range data {
		// Convert byte arrays to strings
		idString := string(item["id"].([]uint8))
		id, _ := strconv.Atoi(idString)
		id64 := int64(id)
		name := string(item["name"].([]uint8))
		nim := string(item["nim"].([]uint8))
		major := string(item["major"].([]uint8))
		is_registered := string(item["is_registered"].([]uint8))
		is_registered_string := "No"
		is_registered_bool := false
		if is_registered == "1" {
			is_registered_string = "Yes"
			is_registered_bool = true
		}

		// Create a map with string values
		data := model.GraduationShowResponse{
			Id:                   id64,
			Name:                 name,
			Nim:                  nim,
			Major:                major,
			Is_registered:        is_registered_bool,
			Is_registered_string: is_registered_string,
		}
		jsonData = append(jsonData, data)
	}

	return jsonData, nil

}
