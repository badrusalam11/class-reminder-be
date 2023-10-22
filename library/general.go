package library

import (
	model "class-reminder-be/model"
	"encoding/json"
	"time"
)

func SetResponse(code string, description string, dataStruct interface{}) (model.Response, []byte) {
	response := model.Response{
		Code:        code,
		Description: description,
		Data:        dataStruct,
	}
	responseJSON, _ := json.Marshal(&response)

	return response, responseJSON
}

func CurrTimestamp() string {
	// Get the current time
	currentTime := time.Now()

	// Format it as "2006-01-02 15:04:05"
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	return formattedTime
}
