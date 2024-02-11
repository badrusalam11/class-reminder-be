package useruc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"fmt"
)

func Register(registerData model.RegisterUserRequest) error {
	// Update tbl_user_notif
	err := database.RegisterStudent(registerData.Name, registerData.Nim, registerData.Phone, registerData.Major, registerData.Class)
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	return nil

}
