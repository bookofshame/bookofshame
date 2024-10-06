package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Name               string
	Version            string
	Url                string
	Port               string
	Ip                 string
	Env                string
	TursoDbUrl         string
	TursoDbAuthToken   string
	SmtpHost           string
	SmtpPort           string
	SmtpUsername       string
	SmtpPassword       string
	SmtpFromEmail      string
	SmsHost            string
	SmsApiKey          string
	SmsSenderId        string
	R2AccessKey        string
	R2SecretAccessKey  string
	R2Host             string
	R2BucketName       string
	R2AccountId        string
	JwtSecret          string
	ReCaptchaSecretKey string
	ReCaptchaHost      string
	// R2Token              string
}

func New() Config {
	if err := godotenv.Load(); err != nil {
		log.Panic("error loading .env file.")
	}

	cfg := Config{
		Name:               os.Getenv("APP_NAME"),
		Version:            os.Getenv("APP_VERSION"),
		Url:                os.Getenv("APP_URL"),
		Port:               os.Getenv("APP_PORT"),
		Ip:                 os.Getenv("APP_IP"),
		Env:                os.Getenv("APP_ENV"),
		TursoDbUrl:         os.Getenv("TURSO_DB_URL"),
		TursoDbAuthToken:   os.Getenv("TURSO_DB_AUTH_TOKEN"),
		SmtpHost:           os.Getenv("SMTP_HOST"),
		SmtpPort:           os.Getenv("SMTP_PORT"),
		SmtpUsername:       os.Getenv("SMTP_USERNAME"),
		SmtpPassword:       os.Getenv("SMTP_PASSWORD"),
		SmtpFromEmail:      os.Getenv("SMTP_FROM_EMAIL"),
		SmsHost:            os.Getenv("SMS_HOST"),
		SmsApiKey:          os.Getenv("SMS_API_KEY"),
		SmsSenderId:        os.Getenv("SMS_SENDER_ID"),
		JwtSecret:          os.Getenv("JWT_SECRET"),
		R2AccountId:        os.Getenv("R2_ACCOUNT_ID"),
		R2AccessKey:        os.Getenv("R2_ACCESS_KEY"),
		R2SecretAccessKey:  os.Getenv("R2_SECRET_ACCESS_KEY"),
		R2Host:             os.Getenv("R2_HOST"),
		R2BucketName:       os.Getenv("R2_BUCKET_NAME"),
		ReCaptchaHost:      os.Getenv("RECAPTCHA_HOST"),
		ReCaptchaSecretKey: os.Getenv("RECAPTCHA_SECRET_KEY"),
	}

	return cfg
}
