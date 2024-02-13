package main

import (
	"class-reminder-be/config"
	database "class-reminder-be/database/helper"
	authHandler "class-reminder-be/handler/auth"
	courseHandler "class-reminder-be/handler/course"
	eventHandler "class-reminder-be/handler/event"
	notifHandler "class-reminder-be/handler/notif"
	userHandler "class-reminder-be/handler/user"
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

	apiCallerWithSession(config.V1Route+"user/register", userHandler.Register)
	apiCallerWithSession(config.V1Route+"user/show", userHandler.Show)
	apiCallerWithSession(config.V1Route+"user/show/detail", userHandler.Detail)
	apiCallerWithSession(config.V1Route+"user/edit", userHandler.Edit)
	apiCallerWithSession(config.V1Route+"user/delete", userHandler.Delete)

	apiCallerWithSession(config.V1Route+"course/list", courseHandler.List)
	apiCallerWithSession(config.V1Route+"course/create", courseHandler.Create)

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
