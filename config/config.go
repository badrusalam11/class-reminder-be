package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
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
	// First try to load from .env
	err := godotenv.Load()
	if err == nil {
		// Use OS environment variables
		viper.AutomaticEnv()
		fmt.Println("✅ Loaded configuration from .env")
		loadFromEnv()
		return
	}

	// Fall back to config.json
	viper.SetConfigFile("config.json")
	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("❌ failed to load config.json: %w", err))
	}

	fmt.Println("✅ Loaded configuration from config.json")
	loadFromViper()
}

func loadFromEnv() {
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBUsername = os.Getenv("DB_USERNAME")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")

	FirebaseKey = os.Getenv("FIREBASE_KEY")
	FirebaseURL = os.Getenv("FIREBASE_URL")

	RundeckURL = os.Getenv("RUNDECK_URL")
	RundeckUsername = os.Getenv("RUNDECK_USERNAME")
	RundeckPassword = os.Getenv("RUNDECK_PASSWORD")
	RundeckProjectName = os.Getenv("RUNDECK_PROJECT_NAME")
	RundeckToken = os.Getenv("RUNDECK_TOKEN")
	RundeckTimeout = getEnvAsInt("RUNDECK_TIMEOUT", 30)

	WhatsappURL = os.Getenv("WHATSAPP_URL")
	WhatsappApiKey = os.Getenv("WHATSAPP_API_KEY")
	WhatsappAppKey = os.Getenv("WHATSAPP_APP_KEY")
	WhatsappAuthKey = os.Getenv("WHATSAPP_AUTH_KEY")
	WhatsappSender = os.Getenv("WHATSAPP_SENDER")

	AppBaseUrl = os.Getenv("APP_BASE_URL")
	AppPort = os.Getenv("APP_PORT")
}

func loadFromViper() {
	DBHost = GetString("database.host")
	DBPort = GetString("database.port")
	DBUsername = GetString("database.username")
	DBPassword = GetString("database.password")
	DBName = GetString("database.dbName")

	FirebaseKey = GetString("firebase.key")
	FirebaseURL = GetString("firebase.url")

	RundeckURL = GetString("rundeck.url")
	RundeckUsername = GetString("rundeck.username")
	RundeckPassword = GetString("rundeck.password")
	RundeckProjectName = GetString("rundeck.project_name")
	RundeckToken = GetString("rundeck.token")
	RundeckTimeout = viper.GetInt("rundeck.timeout")

	WhatsappURL = GetString("whatsapp.url")
	WhatsappApiKey = GetString("whatsapp.api_key")
	WhatsappAppKey = GetString("whatsapp.app_key")
	WhatsappAuthKey = GetString("whatsapp.auth_key")
	WhatsappSender = GetString("whatsapp.sender")

	AppBaseUrl = GetString("app.base_url")
	AppPort = GetString("app.port")
}

// getEnvAsInt returns env as int with fallback default
func getEnvAsInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal
	}
	var val int
	_, err := fmt.Sscanf(valStr, "%d", &val)
	if err != nil {
		return defaultVal
	}
	return val
}
