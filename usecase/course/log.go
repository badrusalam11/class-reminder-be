package courseuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/model"
	"fmt"
	"strconv"
)

func Log() ([]model.CourseLogResponse, error) {
	// select data from tbl_user_notif
	data, err := database.GetTrxLog()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)

	var jsonData []model.CourseLogResponse
	for _, item := range data {
		// Convert byte arrays to strings
		idString := string(item["id"].([]uint8))
		id, _ := strconv.Atoi(idString)
		// id64 := int64(id)
		title := string(item["title"].([]uint8))
		user_success := string(item["user_success"].([]uint8))
		trx_type := string(item["trx_type"].([]uint8))
		trx_date := library.GetTrxDate(string(item["trx_date"].([]uint8)))

		// Create a map with string values
		data := model.CourseLogResponse{
			Id:          id,
			Title:       title,
			UserSuccess: user_success,
			TrxType:     trx_type,
			TrxDate:     trx_date,
		}
		// data := map[string]interface{}{
		// 	"title": title,
		// 	"id":    id,
		// }

		jsonData = append(jsonData, data)
	}

	return jsonData, nil

}
