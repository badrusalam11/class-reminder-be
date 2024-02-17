package library

import (
	model "class-reminder-be/model"
	"encoding/json"
	"strconv"
	"strings"
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

func GetJobDay(day string) string {
	jobMap := make(map[string]string)
	jobMap["MON"] = "SUN"
	jobMap["TUE"] = "MON"
	jobMap["WED"] = "TUE"
	jobMap["THU"] = "WED"
	jobMap["FRI"] = "THU"
	jobMap["SAT"] = "FRI"
	jobMap["SUN"] = "SAT"
	return jobMap[day]
}

func GenerateJobName(id int64, title string) string {
	idString := strconv.FormatInt(id, 10)
	titleString := strings.ReplaceAll(title, " ", "_")
	// Get the current time
	// currentTime := time.Now()
	// // Format the time as DDMMYYYY
	// formattedDate := currentTime.Format("02012006")
	// job := idString + ":" + titleString + "_" + formattedDate
	job := idString + ":" + titleString
	return job
}

func GetDayFromPrefix(prefix string) string {
	dayMap := make(map[string]string)
	dayMap["MON"] = "Monday"
	dayMap["TUE"] = "Tuesday"
	dayMap["WED"] = "Wednesday"
	dayMap["THU"] = "Thursday"
	dayMap["FRI"] = "Friday"
	dayMap["SAT"] = "Saturday"
	dayMap["SUN"] = "Sunday"
	return dayMap[prefix]
}

func GetTrxDate(dateString string) string {
	// Parse the input timestamp string
	layout := "2006-01-02 15:04:05"
	timestamp, err := time.Parse(layout, dateString)
	if err != nil {
		return ""
	}

	// Format the timestamp as per the desired format
	formattedTime := timestamp.Format("02/01/2006 15:04:05")
	return formattedTime
}
