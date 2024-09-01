package thesisuc

import (
	"class-reminder-be/database"
	"class-reminder-be/library"
	"class-reminder-be/repository"
	"fmt"
	"strings"
)

func Blast() error {
	// Select data from tbl_user_notif
	data, err := database.GetUserThesis()
	if err != nil {
		return err
	}

	var notifArr []string
	userNotAttend := make([]map[string]interface{}, 0)

	for i := 0; i < len(data); i++ {
		lastAttendanceDate := data[i]["last_attendance_date"].([]uint8)
		isAttendThisWeek, _ := library.IsDateInCurrentWeek(string(lastAttendanceDate))
		if isAttendThisWeek {
			continue
		}

		// Assuming "no_hp" contains the phone numbers in a byte slice ([]uint8)
		noHPBytes := data[i]["no_hp"].([]uint8)
		noHPString := string(noHPBytes)

		// Split the string into a slice of notification IDs
		notifIDs := strings.Split(noHPString, ",")
		notifArr = append(notifArr, notifIDs...)

		// Append the current user's data to the userNotAttend slice
		userNotAttend = append(userNotAttend, data[i])
	}

	fmt.Println(userNotAttend)

	// Blast to WhatsApp
	trxType := "Thesis"
	_, err = repository.SendToWhatsapp(notifArr, userNotAttend, trxType)
	if err != nil {
		return err
	}
	return nil
}
