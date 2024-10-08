package useruc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"fmt"
)

func Edit(editData model.EditUserRequest) error {
	// Update tbl_user_notif
	err := database.EditStudent(
		editData.Name,
		editData.Nim,
		editData.Phone,
		editData.Major,
		editData.TuitionFee,
		editData.LastPaymentDate,
		editData.Logbook,
		editData.IsRegisGraduation,
	)
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	return nil

}
