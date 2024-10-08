package useruc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/model"
	"fmt"
)

func Detail(request model.DetailUserRequest) (*model.Result, error) {
	// select data from database
	nim := request.Nim
	fmt.Println(nim)
	data, err := database.GetDetailStudentInfo(nim)
	fmt.Println(data)
	if err != nil {
		fmt.Println("error database", err)
		return nil, err
	}

	resultMap := make(map[string]*model.Result)

	for _, item := range data {
		// Convert byte arrays to strings
		name := string(item["name"].([]byte))
		nim := string(item["nim"].([]byte))
		noHP := string(item["no_hp"].([]byte))
		major := string(item["major"].([]byte))
		va_account := string(item["va_account"].([]byte))
		tuition_fee := int(item["bill"].(int64))
		last_payment_date := library.GetDateYMD(string(item["last_payment_date"].([]uint8)))
		is_regis_graduation := int(item["is_regis_graduation"].(int64))
		is_done_thesis := int(item["is_done_thesis"].(int64))
		logbook := int(item["logbook"].(int64))

		key := nim

		newResult := model.Result{
			Major:             major,
			Name:              name,
			NIM:               nim,
			NoHP:              noHP,
			TuitionFee:        tuition_fee,
			VaAccount:         va_account,
			LastPaymentDate:   last_payment_date,
			IsRegisGraduation: is_regis_graduation,
			IsDoneThesis:      is_done_thesis,
			Logbook:           logbook,
		}
		resultMap[key] = &newResult
	}

	// Retrieve the first result from resultMap (assuming there's at least one result)
	var result *model.Result
	for _, r := range resultMap {
		result = r
		break
	}

	return result, nil
}
