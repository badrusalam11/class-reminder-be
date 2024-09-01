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

func GetDateYMD(dateString string) string {
	// Parse the input timestamp string
	layout := "2006-01-02 15:04:05"
	timestamp, err := time.Parse(layout, dateString)
	if err != nil {
		return ""
	}

	// Format the timestamp as per the desired format
	formattedTime := timestamp.Format("2006-01-02")
	return formattedTime
}

func GetLastMonth() time.Time {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()

	// Get the first day of the next month
	firstDayOfNextMonth := time.Date(currentYear, currentMonth+1, 1, 0, 0, 0, 0, now.Location())

	// Subtract one day to get the last day of the current month
	lastDayOfCurrentMonth := firstDayOfNextMonth.AddDate(0, 0, -1)
	return lastDayOfCurrentMonth
}

// isDateInCurrentWeek checks if the given date is within the current week.
func IsDateInCurrentWeek(dateStr string) (bool, error) {
	// Define the layout of the input date string
	layout := "2006-01-02 15:04:05"
	// Parse the string into a time.Time object
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		return false, err
	}
	// Get the current date and time
	now := time.Now()
	// Calculate the start of the week (Monday)
	startOfWeek := now.AddDate(0, 0, -int(now.Weekday())+1) // Adjust to start on Monday
	// Calculate the end of the week (Sunday)
	endOfWeek := startOfWeek.AddDate(0, 0, 6)
	// Check if the given date is within the start and end of the week
	isInWeek := date.After(startOfWeek) && date.Before(endOfWeek) || date.Equal(startOfWeek) || date.Equal(endOfWeek)
	return isInWeek, nil
}
