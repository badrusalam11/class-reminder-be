package paymentReminderuc

import (
	"class-reminder-be/database"
	"class-reminder-be/model"
	"class-reminder-be/repository"
	"fmt"
)

func JobTrigger(request model.JobTriggerRequest) error {
	// jsonData := model.JobDetailResponse{}
	// select data from job
	item, err := database.GetUserPaymentByNim(request.Nim)
	if err != nil {
		return err
	}
	fmt.Println(item)
	// select trx type
	trxTypeDB, err := database.GetTrxTypeFromDB(request.IdEvent)
	if err != nil {
		return err
	}
	trxTypeuint := trxTypeDB["trx_type"].([]uint8)
	trxType := string(trxTypeuint)
	no_hp := string(item["no_hp"].([]uint8))

	var noHpArr = []string{no_hp}
	var arrayofMaps []map[string]interface{}
	var itemArr = append(arrayofMaps, item)

	_, err = repository.SendToWhatsapp(noHpArr, itemArr, trxType)
	if err != nil {
		return err
	}
	return nil

}
