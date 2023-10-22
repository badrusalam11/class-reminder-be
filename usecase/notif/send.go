package notifuc

import (
	database "class-reminder-be/database"
	"class-reminder-be/repository"
	"fmt"
	"strings"
)

func Send(eventId int) (string, error) {
	// get event from tbl_event
	// check id_event in tbl_event_type if is_specific user or not
	event, err := database.GetEventFromDB(eventId)
	if err != nil {
		fmt.Print("send", err)
	}
	fmt.Println(event["is_specific_user"])
	is_specific := event["is_specific_user"].(int64)
	var notifArr []string
	// if specific user, check tbl_user_event
	// call tbl_user_notif
	if is_specific == 1 {
		userEvent, err := database.GetUserEventFromDB(eventId)
		if err != nil {
			fmt.Print("send", err)
		}
		fmt.Println(userEvent[0]["notif_id"])
		for i := 0; i < len(userEvent); i++ {
			// Assuming userEvent["notif_id"] is a byte slice ([]uint8)
			notifIDsBytes := userEvent[i]["notif_id"].([]uint8)
			notifIDsString := string(notifIDsBytes)
			fmt.Println(notifIDsString)

			// Now you can split the string into a slice
			notifIDs := strings.Split(notifIDsString, ",")
			notifArr = append(notifArr, notifIDs...)

		}
		fmt.Println(notifArr)

		// else, blast to all user in tbl_user_notif
	} else {
		userNotif, err := database.GetUserNotifFromDB()
		if err != nil {
			fmt.Print("send", err)
		}
		for i := 0; i < len(userNotif); i++ {
			// Assuming userEvent["notif_id"] is a byte slice ([]uint8)
			notifIDsBytes := userNotif[i]["notif_id"].([]uint8)
			notifIDsString := string(notifIDsBytes)
			fmt.Println(notifIDsString)

			// Now you can split the string into a slice
			notifIDs := strings.Split(notifIDsString, ",")
			notifArr = append(notifArr, notifIDs...)

		}
	}
	// call to firebase
	response, err := repository.SendToFirebase(notifArr)
	if err != nil {
		return "", err
	}
	return response, err
}
