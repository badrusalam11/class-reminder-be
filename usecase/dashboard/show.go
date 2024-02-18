package dashboarduc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"fmt"
	"strconv"
)

func Show() ([]model.DashboardShowResponse, error) {
	// select data from tbl_user_notif
	data, err := database.GetTransactionInfo()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)

	var jsonData []model.DashboardShowResponse
	for _, item := range data {

		// Convert byte arrays to strings
		title := string(item["title"].([]uint8))
		total := string(item["total"].([]uint8))
		totalInt, _ := strconv.Atoi(total)
		total_str := string(item["total_str"].([]uint8))

		// Create a map with string values
		data := model.DashboardShowResponse{
			Title:    title,
			Total:    totalInt,
			TotalStr: total_str,
		}
		// data := map[string]interface{}{
		// 	"title": title,
		// 	"id":    id,
		// }

		jsonData = append(jsonData, data)
	}

	return jsonData, nil

}
