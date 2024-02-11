package useruc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"fmt"
)

func Delete(data model.EditUserRequest) error {
	// Update tbl_user_notif
	err := database.DeleteStudent(data.Nim)
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	return nil

}
