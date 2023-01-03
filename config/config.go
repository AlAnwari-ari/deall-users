package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APP_VERSION,
	APP_NAME,
	KEY_ENCRYPTION,
	IP_ADDRESS,
	HTTP_PORT,
	RUN_MODE,
	DEFAULT_TIME_FORMAT,
	TIME_LOCATION,

	DB_HOST,
	DB_PORT,
	DB_USER,
	DB_NAME,
	DB_PASSWORD,
	DB_TIMEOUT,

	JWT_ACCESS_KEY,
	JWT_REFRESH_KEY string
	JWT_ACCESS_EXPIRATION,
	JWT_REFRESH_EXPIRATION int
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("sad .env file found")
	}

	APP_VERSION = os.Getenv("APP_VERSION")
	APP_NAME = os.Getenv("APP_NAME")
	KEY_ENCRYPTION = os.Getenv("KEY_ENCRYPTION")
	IP_ADDRESS = os.Getenv("IP_ADDRESS")
	HTTP_PORT = os.Getenv("HTTP_PORT")
	RUN_MODE = os.Getenv("RUN_MODE")
	DEFAULT_TIME_FORMAT = os.Getenv("DEFAULT_TIME_FORMAT")
	TIME_LOCATION = os.Getenv("TIME_LOCATION")

	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_NAME = os.Getenv("DB_NAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_TIMEOUT = os.Getenv("DB_TIMEOUT")

	JWT_ACCESS_KEY = os.Getenv("JWT_ACCESS_KEY")
	JWT_REFRESH_KEY = os.Getenv("JWT_REFRESH_KEY")
	JWT_ACCESS_EXPIRATION, _ = strconv.Atoi(os.Getenv("JWT_ACCESS_EXPIRATION"))
	JWT_REFRESH_EXPIRATION, _ = strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRATION"))
}
