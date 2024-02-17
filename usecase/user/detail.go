package useruc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"fmt"
)

func Detail(request model.DetailUserRequest) (*model.Result, error) {
	// select data from database
	nim := request.Nim
	data, err := database.GetDetailStudentInfo(nim)
	fmt.Println(data)
	if err != nil {
		fmt.Println("error database", err)
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
		classID := int(item["class_id"].(int64))
		classTitle := string(item["class_title"].([]byte))
		va_account := string(item["va_account"].([]byte))
		tuition_fee := int(item["bill"].(int64))
		key := nim
		// key := major + "_" + name + "_" + nim + "_" + noHP
		if existingResult, exists := resultMap[key]; exists {
			// If the entry already exists, add the class detail to the existing entry
			existingResult.Class = append(existingResult.Class, model.ClassDetail{
				ID:    classID,
				Title: classTitle,
			})
			existingResult.ClassArr = append(existingResult.ClassArr, classID)
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
				ClassArr:   append(classArr, classID),
				TuitionFee: tuition_fee,
				VaAccount:  va_account,
			}
			resultMap[key] = &newResult
		}
	}

	// Retrieve the first result from resultMap (assuming there's at least one result)
	var result *model.Result
	for _, r := range resultMap {
		result = r
		break
	}

	return result, nil
}
