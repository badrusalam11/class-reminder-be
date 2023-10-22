package helper

import (
	"class-reminder-be/config"
	"database/sql"
	"fmt"
)

var Db *sql.DB

func ConnectDB() {
	fmt.Println("database init")
	err := InitDB()
	if err != nil {
		fmt.Println("Error initializing the database:", err)
		return
	}
	// defer CloseDB()
	fmt.Println(err)
	fmt.Println("database", Db)
	fmt.Println("database close")
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

// CloseDB closes the database connection
// func CloseDB() {
// 	if Db != nil {
// 		Db.Close()
// 	}
// }
