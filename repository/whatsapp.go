package repository

import (
	"bytes"
	"class-reminder-be/config"
	"class-reminder-be/database"
	"class-reminder-be/database/helper"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func SendToWhatsapp(notifArr []string, data []map[string]interface{}, trxType string) (count int, err error) {
	var message string
	var mapData map[string]interface{}
	count = 0
	//get content from db
	content, err := database.GetContentFromDB(trxType)
	content_string := string(content["content"].([]uint8))
	additional_data := string(content["additional_data"].([]uint8))
	if err != nil {
		fmt.Println(err)
		return count, err
	}
	for i := 0; i < len(notifArr); i++ {
		message, err = helper.MappingMessage(content_string, additional_data, data[i])
		if err != nil {
			fmt.Println(err)
			return count, err
		}
		response, _ := whatsappApiCaller(message, string(notifArr[i]))

		// Unmarshal the JSON string into the map
		json.Unmarshal([]byte(response), &mapData)
		fmt.Println("data", data)
		if mapData["message_status"] == "Success" {
			count++
		}
	}
	return count, nil
}

func SendToWhatsappSpecific(notifArr string, data map[string]interface{}, trxType string) (count int, err error) {
	var message string
	var mapData map[string]interface{}
	count = 0
	//get content from db
	content, err := database.GetContentFromDB(trxType)
	content_string := string(content["content"].([]uint8))
	additional_data := string(content["additional_data"].([]uint8))
	if err != nil {
		fmt.Println(err)
		return count, err
	}

	message, err = helper.MappingMessage(content_string, additional_data, data)
	if err != nil {
		fmt.Println(err)
		return count, err
	}
	response, _ := whatsappApiCaller(message, string(notifArr))
	fmt.Println(response)
	// Unmarshal the JSON string into the map
	json.Unmarshal([]byte(response), &mapData)
	fmt.Println("data", data)
	if mapData["message_status"] == "Success" {
		count++
	}

	return count, nil
}

func BlastToWhatsapp(notifArr []map[string]interface{}, message string) (count int, err error) {
	var data map[string]interface{}
	for i := 0; i < len(notifArr); i++ {
		response, _ := whatsappApiCaller(message, string(notifArr[i]["no_hp"].([]uint8)))

		// Unmarshal the JSON string into the map
		json.Unmarshal([]byte(response), &data)
		fmt.Println(data)
		if data["message_status"] == "Success" {
			count++
		}
	}
	return count, nil
}

func whatsappApiCaller(message string, no_hp string) (response string, err error) {
	var request WhatsappNewRequest
	request.AppKey = config.WhatsappAppKey
	request.Authkey = config.WhatsappAuthKey
	request.To = no_hp
	request.Sandbox = "false"
	request.Message = message
	reqBody, _ := jsoniter.Marshal(request)
	fmt.Println(request)
	req, err := http.NewRequest("POST", config.WhatsappURL, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}

	// Set headers if needed
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and send the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check if the response is compressed (e.g., gzip)
	var responseBody []byte
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return "", err
		}
		defer reader.Close()
		responseBody, err = ioutil.ReadAll(reader)
		if err != nil {
			return "", err
		}
	default:
		responseBody, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
	}
	fmt.Println("response", string(responseBody))
	return string(responseBody), nil
}
