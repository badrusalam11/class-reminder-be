package graduationuc

import (
	"class-reminder-be/database"
	"class-reminder-be/repository"
	"fmt"
	"strings"
)

func Send() error {
	// select data from tbl_user_notif
	data, err := database.GetNotRegisGraduation()
	if err != nil {
		return err
	}
	fmt.Println(data)
	var notifArr []string
	for i := 0; i < len(data); i++ {
		// Assuming userEvent["notif_id"] is a byte slice ([]uint8)
		notifIDsBytes := data[i]["no_hp"].([]uint8)

		notifIDsString := string(notifIDsBytes)
		fmt.Println(notifIDsString)

		// Now you can split the string into a slice
		notifIDs := strings.Split(notifIDsString, ",")
		notifArr = append(notifArr, notifIDs...)

	}

	// blast to whatsapp
	message := "hai"
	_, err = repository.BlastToWhatsapp(data, message)
	if err != nil {
		return err
	}
	return nil
}
