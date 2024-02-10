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
	http.HandleFunc("/notif/send", notifHandler.Send) // may hit by rundeck
	// http.HandleFunc("/event/create", eventHandler.Create)
	apiCallerWithSession(config.V1Route+"event/create", eventHandler.Create)
	apiCallerWithSession(config.V1Route+"notif/blast", notifHandler.Blast)
	apiCallerWithSession(config.V1Route+"notif/blast/history", notifHandler.BlastHistory)

	// API route with CORS middleware
	// http.Handle(config.V1Route+"login", middleware.CorsEnabled(http.HandlerFunc(authHandler.Login)))
	// http.Handle(config.V1Route+"notif/register", middleware.CorsEnabled(http.HandlerFunc(notifHandler.Register)))
	apiCaller(config.V1Route+"login", authHandler.Login)
	apiCaller(config.V1Route+"notif/register", notifHandler.Register)
	http.ListenAndServe(":9090", nil)
}

func apiCaller(route string, handler http.HandlerFunc) {
	http.Handle(route, middleware.CorsEnabled(http.HandlerFunc(handler)))
}

func apiCallerWithSession(route string, handler http.HandlerFunc) {
	http.Handle(route, middleware.CorsEnabled(middleware.AuthMiddleware(http.HandlerFunc(handler))))
}
