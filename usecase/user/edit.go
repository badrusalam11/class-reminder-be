package useruc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"fmt"
)

func Edit(registerData model.EditUserRequest) error {
	// Update tbl_user_notif
	err := database.EditStudent(
		registerData.Name,
		registerData.Nim,
		registerData.Phone,
		registerData.Major,
		registerData.Class,
		registerData.TuitionFee,
		registerData.VaAccount,
		registerData.LastPaymentDate,
	)
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	return nil

}
