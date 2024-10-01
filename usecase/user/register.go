package useruc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/model"
	"fmt"
)

func Register(registerData model.RegisterUserRequest) error {
	virtualAccount := library.GenerateVA()
	// Update tbl_user_notif
	err := database.RegisterStudent(
		registerData.Name,
		registerData.Nim,
		registerData.Phone,
		registerData.Major,
		registerData.TuitionFee,
		virtualAccount,
		registerData.LastPaymentDate,
		registerData.IsRegisGraduation,
		registerData.IsDoneThesis,
	)
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	return nil

}
