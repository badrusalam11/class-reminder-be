package notifuc

import (
	"class-reminder-be/database"
)

func BlastHistory() ([]map[string]interface{}, error) {
	// select data from tbl_user_notif
	data, err := database.GetBlastHistory()
	if err != nil {
		return nil, err
	}

	return data, nil

}
