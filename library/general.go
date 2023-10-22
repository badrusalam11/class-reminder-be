package library

import (
	model "class-reminder-be/model"
	"encoding/json"
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
