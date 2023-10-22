package main

import (
	"class-reminder-be/config"
	database "class-reminder-be/database/helper"
	authHandler "class-reminder-be/handler/auth"
	eventHandler "class-reminder-be/handler/event"
	notifHandler "class-reminder-be/handler/notif"
	middleware "class-reminder-be/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	database.ConnectDB()
}

func main() {
	http.HandleFunc("/notif/send", notifHandler.Send)
	http.HandleFunc("/event/create", eventHandler.Create)
	//api route
	http.HandleFunc(config.V1Route+"login", authHandler.Login)
	http.Handle(config.V1Route+"notif/register", middleware.AuthMiddleware(http.HandlerFunc(notifHandler.Register)))

	http.ListenAndServe(":9090", nil)
}
