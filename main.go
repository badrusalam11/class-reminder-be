package main

import (
	"class-reminder-be/config"
	database "class-reminder-be/database/helper"
	authHandler "class-reminder-be/handler/auth"
	courseHandler "class-reminder-be/handler/course"
	dashboardHandler "class-reminder-be/handler/dashboard"
	eventHandler "class-reminder-be/handler/event"
	graduationHandler "class-reminder-be/handler/graduation"
	notifHandler "class-reminder-be/handler/notif"
	paymentReminderHandler "class-reminder-be/handler/paymentReminder"
	thesisHandler "class-reminder-be/handler/thesis"
	userHandler "class-reminder-be/handler/user"
	middleware "class-reminder-be/middleware"

	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	config.Init()
	err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
}

func main() {
	// http.HandleFunc("/notif/send", notifHandler.Send) // may hit by rundeck
	// http.HandleFunc("/event/create", eventHandler.Create)
	apiCaller("/notif/send", notifHandler.Send) // may hit by rundeck
	apiCallerWithSession(config.V1Route+"event/create", eventHandler.Create)
	apiCallerWithSession(config.V1Route+"event/edit", eventHandler.Edit)
	apiCallerWithSession(config.V1Route+"notif/blast", notifHandler.Blast)
	apiCallerWithSession(config.V1Route+"notif/blast/history", notifHandler.BlastHistory)

	apiCallerWithSession(config.V1Route+"user/register", userHandler.Register)
	apiCallerWithSession(config.V1Route+"user/show", userHandler.Show)
	apiCallerWithSession(config.V1Route+"user/show/detail", userHandler.Detail)
	apiCallerWithSession(config.V1Route+"user/edit", userHandler.Edit)
	apiCallerWithSession(config.V1Route+"user/delete", userHandler.Delete)

	apiCallerWithSession(config.V1Route+"course/list", courseHandler.List)
	apiCallerWithSession(config.V1Route+"course/create", courseHandler.Create)
	apiCallerWithSession(config.V1Route+"course/edit", courseHandler.Edit)
	apiCallerWithSession(config.V1Route+"course/delete", courseHandler.Delete)
	apiCallerWithSession(config.V1Route+"course/show", courseHandler.Show)
	apiCallerWithSession(config.V1Route+"course/show/detail", courseHandler.Detail)
	apiCallerWithSession(config.V1Route+"course/log/show", courseHandler.Log)

	apiCallerWithSession(config.V1Route+"payment-reminder/show", paymentReminderHandler.Show)
	apiCallerWithSession(config.V1Route+"payment-reminder/job/detail", paymentReminderHandler.JobDetail)
	apiCallerWithSession(config.V1Route+"payment-reminder/job/trigger", paymentReminderHandler.JobTrigger)

	apiCallerWithSession(config.V1Route+"graduation/show", graduationHandler.Show)
	apiCaller(config.V1Route+"graduation/blast", graduationHandler.Blast)
	apiCallerWithSession(config.V1Route+"graduation/send", graduationHandler.Send)

	apiCallerWithSession(config.V1Route+"thesis/show", thesisHandler.Show)
	apiCallerWithSession(config.V1Route+"thesis/blast", thesisHandler.Blast)
	apiCallerWithSession(config.V1Route+"thesis/send", thesisHandler.Send)

	apiCallerWithSession(config.V1Route+"dashboard/show", dashboardHandler.Show)
	// API route with CORS middleware
	// http.Handle(config.V1Route+"login", middleware.CorsEnabled(http.HandlerFunc(authHandler.Login)))
	// http.Handle(config.V1Route+"notif/register", middleware.CorsEnabled(http.HandlerFunc(notifHandler.Register)))
	apiCaller(config.V1Route+"login", authHandler.Login)
	apiCaller(config.V1Route+"notif/register", notifHandler.Register)
	http.ListenAndServe(":"+config.AppPort, nil)
}

func apiCaller(route string, handler http.HandlerFunc) {
	http.Handle(route, middleware.CorsEnabled(http.HandlerFunc(handler)))
}

func apiCallerWithSession(route string, handler http.HandlerFunc) {
	http.Handle(route, middleware.CorsEnabled(middleware.AuthMiddleware(http.HandlerFunc(handler))))
}
