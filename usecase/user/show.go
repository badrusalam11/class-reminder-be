package useruc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	// "fmt"
)

func Show() ([]model.Result, error) {
	// select data from tbl_user_notif
	data, err := database.GetStudentInfo()
	if err != nil {
		return nil, err
	}

	resultMap := make(map[string]*model.Result)
	for _, item := range data {
		// Convert byte arrays to strings
		name := string(item["name"].([]byte))
		nim := string(item["nim"].([]byte))
		noHP := string(item["no_hp"].([]byte))
		major := string(item["major"].([]byte))
		key := nim
		// key := major + "_" + name + "_" + nim + "_" + noHP

		// If the entry does not exist, create a new entry
		newResult := model.Result{
			Major: major,
			Name:  name,
			NIM:   nim,
			NoHP:  noHP,
		}
		resultMap[key] = &newResult
	}

	// Convert resultMap values to the final resultArray
	var resultArray []model.Result
	for _, result := range resultMap {
		resultArray = append(resultArray, *result)
	}

	return resultArray, nil
}
