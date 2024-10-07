package thesisuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/repository"
	"fmt"
)

func Send(nim string) error {
	// select data from tbl_user_notif
	data, err := database.GetUserThesisSpecific(nim)
	if err != nil {
		return err
	}
	notifIDsString := ""
	date := library.GetLastMonth() // Corrected: Call the function to get time.Time
	// Convert time.Time to string in the desired format
	dateString := date.Format("02 January 2006")
	// Convert string to []uint8
	dateBytes := []uint8(dateString)
	for i := 0; i < len(data); i++ {
		// Assuming userEvent["notif_id"] is a byte slice ([]uint8)
		notifIDsBytes := data["no_hp"].([]uint8)
		data["date"] = dateBytes // Assign marshaled date to data
		// Convert []uint8 to string
		notifIDsString = string(notifIDsBytes)
	}

	fmt.Println(data)

	// Blast to WhatsApp
	trxType := "Thesis"
	_, err = repository.SendToWhatsappSpecific(notifIDsString, data, trxType)
	if err != nil {
		return err
	}
	return nil
}
