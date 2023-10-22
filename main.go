package main

import (
	database "class-reminder-be/database/helper"
	authHandler "class-reminder-be/handler/auth"
	eventHandler "class-reminder-be/handler/event"
	notifHandler "class-reminder-be/handler/notif"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	database.ConnectDB()
}

func main() {
	http.HandleFunc("/notif/send", notifHandler.Send)
	http.HandleFunc("/event/create", eventHandler.Create)
	http.HandleFunc("/login", authHandler.Login)
	http.ListenAndServe(":9090", nil)
}
