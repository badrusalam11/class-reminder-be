package notifuc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
)

func Register(registerData model.RegisterNotifRequest) (string, error) {
	// Update tbl_user_notif
	code, err := database.UpdateNotifId(registerData.Username, registerData.NotifId)
	if err != nil {
		return "01", err
	}
	if code == "02" {
		return code, nil
	}
	return "", nil

}
