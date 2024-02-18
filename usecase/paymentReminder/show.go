package paymentReminderuc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"fmt"
	"strconv"
)

func Show() ([]model.PaymentReminderResponse, error) {
	// select data from tbl_user_notif
	data, err := database.GetUserPayment()
	if err != nil {
		return nil, err
	}
	fmt.Println(data)

	var jsonData []model.PaymentReminderResponse
	for _, item := range data {
		// Convert byte arrays to strings
		idString := string(item["id"].([]uint8))
		id, _ := strconv.Atoi(idString)
		name := string(item["name"].([]uint8))
		nim := string(item["nim"].([]uint8))
		bill := string(item["bill"].([]uint8))
		tuition_fee, _ := strconv.Atoi(bill)
		va_account := string(item["va_account"].([]uint8))
		last_payment_date := string(item["last_payment_date"].([]uint8))
		due_date := string(item["due_date"].([]uint8))

		// Create a map with string values
		data := model.PaymentReminderResponse{
			Id:              id,
			Name:            name,
			Nim:             nim,
			TuitionFee:      tuition_fee,
			VaAccount:       va_account,
			LastPaymentDate: last_payment_date,
			DueDate:         due_date,
		}
		// data := map[string]interface{}{
		// 	"title": title,
		// 	"id":    id,
		// }

		jsonData = append(jsonData, data)
	}

	return jsonData, nil

}
