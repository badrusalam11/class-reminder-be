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
	trxTypeDB, err := database.GetTrxTypeFromDB(eventId)
	trxTypeuint := trxTypeDB["trx_type"].([]uint8)
	trxType := string(trxTypeuint)
	if err != nil {
		fmt.Print("send", err)
	}
	is_specific := event["is_specific_user"].(int64)
	var notifArr []string
	var userEvent []map[string]interface{}
	// if specific user, check tbl_user_event
	// call tbl_user_notif
	if is_specific == 1 {
		userEvent, err = database.GetUserEventFromDB(eventId)
		fmt.Println(userEvent)
		if err != nil {
			fmt.Print("send", err)
			return "failed", err
		}
		// fmt.Println(userEvent[0]["notif_id"])
		for i := 0; i < len(userEvent); i++ {
			// Assuming userEvent["notif_id"] is a byte slice ([]uint8)
			notifIDsBytes := userEvent[i]["no_hp"].([]uint8)

			notifIDsString := string(notifIDsBytes)
			fmt.Println(notifIDsString)

			// Now you can split the string into a slice
			notifIDs := strings.Split(notifIDsString, ",")
			notifArr = append(notifArr, notifIDs...)

		}
		fmt.Println(notifArr)

		// else, blast to all user in tbl_user_notif
	} else {
		fmt.Println("masuk else")
		userEvent, err = database.GetUserNotifFromDB()
		if err != nil {
			fmt.Print("send", err)
			return "failed", err
		}
		for i := 0; i < len(userEvent); i++ {
			// Assuming userEvent["notif_id"] is a byte slice ([]uint8)
			notifIDsBytes := userEvent[i]["no_hp"].([]uint8)
			userEvent[i]["event"] = event["title"]
			notifIDsString := string(notifIDsBytes)
			fmt.Println(notifIDsString)

			// Now you can split the string into a slice
			notifIDs := strings.Split(notifIDsString, ",")
			notifArr = append(notifArr, notifIDs...)

		}
	}

	// call to whatsapp
	// response, err := repository.SendToFirebase(notifArr)
	count, err := repository.SendToWhatsapp(notifArr, userEvent, trxType)
	if err != nil {
		return "", err
	}
	fmt.Println("count user", count)
	// insert to trx log
	err = database.InsertToTrxLog(eventId, count, trxType)
	response := "success"
	return response, err
}
