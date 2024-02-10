package notifuc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"class-reminder-be/repository"
)

func Blast(registerData model.BlastRequest) (string, error) {
	// select data from tbl_user_notif
	data, err := database.GetNumberForBlast()
	if err != nil {
		return "01", err
	}
	// do blast
	user_success, err := repository.BlastToWhatsapp(data, registerData.Message)
	if err != nil {
		return "EW", err // error whatsapp
	}
	err = database.InsertBlastHistory(registerData.Message, user_success)
	return "", err

}
