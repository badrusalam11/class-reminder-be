package courseuc

import (
	"class-reminder-be/database"
)

func List() ([]map[string]interface{}, error) {
	// select data from tbl_user_notif
	data, err := database.GetCourse()
	if err != nil {
		return nil, err
	}

	return data, nil

}
