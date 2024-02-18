package helper

import (
	"class-reminder-be/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var Db *sql.DB

func ConnectDB() error {
	fmt.Println("database init")
	err := InitDB()
	if err != nil {
		fmt.Println("Error initializing the database:", err)
		return err
	}
	// defer CloseDB()
	fmt.Println(err)
	fmt.Println("database", Db)
	fmt.Println("database close")
	return nil
}

func InitDB() error {
	// Replace these with your actual database connection details
	username := config.DBUsername
	password := config.DBPassword
	host := config.DBHost
	port := config.DBPort
	dbName := config.DBName

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)

	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	fmt.Println("Db", Db)
	if err != nil {
		return err
	}

	// Test the database connection
	err = Db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func MappingMessage(content string, additionalDataJSON string, data map[string]interface{}) (string, error) {
	// Parse additional data JSON
	var additionalData map[string]string
	if err := json.Unmarshal([]byte(additionalDataJSON), &additionalData); err != nil {
		return "", err
	}

	// Replace placeholders in the content string
	for key := range additionalData {
		placeholder := "$" + key

		// Convert the interface{} value to a string
		strValue, _ := toString(data[key])
		// strValue := string(data[key].([]uint8))

		content = strings.ReplaceAll(content, placeholder, strValue)
	}

	return content, nil
}

// toString converts an interface{} value to a string
func toString(value interface{}) (string, bool) {
	switch v := value.(type) {
	case string:
		return v, true
	case time.Time:
		return v.Format(time.RFC3339), true // Format time as a string
	case fmt.Stringer:
		return v.String(), true
	case sql.NullTime:
		if v.Valid {
			return v.Time.Format(time.RFC3339), true // Format time as a string if it is valid
		}
		return "", false
	case int64:
		return fmt.Sprint(v), true // Convert int64 to string using fmt.Sprint
	default:
		// Handle unknown types
		return string(v.([]uint8)), true
	}
}

// CloseDB closes the database connection
// func CloseDB() {
// 	if Db != nil {
// 		Db.Close()
// 	}
// }
