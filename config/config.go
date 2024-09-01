package config

import (
	"github.com/spf13/viper"
)

var (
	DBHost             string
	DBPort             string
	DBUsername         string
	DBPassword         string
	DBName             string
	FirebaseKey        string
	FirebaseURL        string
	RundeckURL         string
	RundeckUsername    string
	RundeckPassword    string
	RundeckProjectName string
	RundeckToken       string
	RundeckTimeout     int
	WhatsappURL        string
	WhatsappApiKey     string
	WhatsappAppKey     string
	WhatsappAuthKey    string
	WhatsappSender     string
	AppBaseUrl         string
	AppPort            string
)

func GetString(key string) string {
	return viper.GetString(key)
}

func Init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	// Database
	DBHost = GetString(`database.host`)
	DBPort = GetString(`database.port`)
	DBUsername = GetString(`database.username`)
	DBPassword = GetString(`database.password`)
	DBName = GetString(`database.dbName`)
	FirebaseKey = GetString(`firebase.key`)
	FirebaseURL = GetString(`firebase.url`)
	RundeckURL = GetString(`rundeck.url`)
	RundeckUsername = GetString(`rundeck.username`)
	RundeckPassword = GetString(`rundeck.password`)
	RundeckProjectName = GetString(`rundeck.project_name`)
	RundeckToken = GetString(`rundeck.token`)
	RundeckTimeout = viper.GetInt(`rundeck.project_name`)
	WhatsappURL = viper.GetString(`whatsapp.url`)
	WhatsappApiKey = viper.GetString(`whatsapp.api_key`)
	WhatsappAppKey = viper.GetString(`whatsapp.app_key`)
	WhatsappAuthKey = viper.GetString(`whatsapp.auth_key`)
	WhatsappSender = viper.GetString(`whatsapp.sender`)
	AppBaseUrl = viper.GetString(`app.base_url`)
	AppPort = viper.GetString(`app.port`)
}
