package thesisuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/repository"
	"fmt"
	"strings"
)

func Blast() error {
	// select data from tbl_user_notif
	data, err := database.GetNotRegisGraduation()
	if err != nil {
		return err
	}

	var notifArr []string
	date := library.GetLastMonth() // Corrected: Call the function to get time.Time
	// Convert time.Time to string in the desired format
	dateString := date.Format("02 January 2006")
	// Convert string to []uint8
	dateBytes := []uint8(dateString)
	for i := 0; i < len(data); i++ {
		// Assuming userEvent["notif_id"] is a byte slice ([]uint8)
		notifIDsBytes := data[i]["no_hp"].([]uint8)
		data[i]["date"] = dateBytes // Assign marshaled date to data
		// Convert []uint8 to string
		notifIDsString := string(notifIDsBytes)
		fmt.Println(notifIDsString)

		// Now you can split the string into a slice
		notifIDs := strings.Split(notifIDsString, ",")
		notifArr = append(notifArr, notifIDs...)
	}

	fmt.Println(data)

	// Blast to WhatsApp
	trxType := "Graduation"
	_, err = repository.SendToWhatsapp(notifArr, data, trxType)
	if err != nil {
		return err
	}
	return nil
}
