package repository

import (
	"bytes"
	"class-reminder-be/config"
	"compress/gzip"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

func SendToFirebase(notifArr []string) (string, error) {
	var request FirebaseRequest
	request.RegistrationIds = notifArr
	request.Notification.Title = "Jadwal event"
	request.Notification.Body = "Sudah saatnya masuk"
	request.Notification.Subtitle = ""
	reqBody, _ := jsoniter.Marshal(request)
	req, err := http.NewRequest("POST", config.FirebaseURL, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	header := GenerateHeader(config.FirebaseKey)
	req.Header = header
	// // Set headers for the outgoing request
	// for key, values := range reqHeaders {
	// 	for _, value := range values {
	// 	}
	// }

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

	return string(responseBody), nil
}

func GenerateHeader(key string) http.Header {

	headers := http.Header{
		"Authorization": []string{key},
		"Content-Type":  []string{"application/json"},
	}

	return headers
}
