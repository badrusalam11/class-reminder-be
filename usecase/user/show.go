package useruc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"strconv"
	// "fmt"
)

func Show() ([]model.Result, error) {
	// select data from tbl_user_notif
	data, err := database.GetStudentInfo()
	if err != nil {
		return nil, err
	}

	resultMap := make(map[string]*model.Result)
	classArr := make([]int, 0)
	for _, item := range data {
		// Convert byte arrays to strings
		name := string(item["name"].([]byte))
		nim := string(item["nim"].([]byte))
		noHP := string(item["no_hp"].([]byte))
		major := string(item["major"].([]byte))
		classIDStr := string(item["class_id"].([]uint8))
		classID, _ := strconv.Atoi(classIDStr)
		classTitle := string(item["class_title"].([]byte))
		key := nim
		// key := major + "_" + name + "_" + nim + "_" + noHP
		if existingResult, exists := resultMap[key]; exists {
			// If the entry already exists, add the class detail to the existing entry
			existingResult.Class = append(existingResult.Class, model.ClassDetail{
				ID:    classID,
				Title: classTitle,
			})
			// existingResult.ClassStr = existingResult.ClassStr + "," + classTitle
		} else {
			// If the entry does not exist, create a new entry
			newResult := model.Result{
				Major: major,
				Name:  name,
				NIM:   nim,
				NoHP:  noHP,
				Class: []model.ClassDetail{{
					ID:    classID,
					Title: classTitle,
				}},
				ClassArr: append(classArr, classID),
			}
			resultMap[key] = &newResult
		}
	}

	// Convert resultMap values to the final resultArray
	var resultArray []model.Result
	for _, result := range resultMap {
		resultArray = append(resultArray, *result)
	}

	return resultArray, nil
}
