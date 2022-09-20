package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/go-playground/validator.v9"
)

type ConfigType struct {
	MongodbUrl   string `validate:"required"`
	Mail         string `validate:"required"`
	MailPassword string `validate:"required"`
}

var ConfigValues ConfigType

func GetConfig() *ConfigType {
	if os.Getenv("APP_ENV") != "prod" || os.Getenv("APP_ENV") != "stg" || os.Getenv("APP_ENV") != "beta" {
		godotenv.Load("./.env")
	}

	ConfigVariables := ConfigType{
		MongodbUrl:   os.Getenv("MONGODB_URL"),
		Mail:         os.Getenv("MAIL"),
		MailPassword: os.Getenv("MAIL_PASSWORD"),
	}

	if ConfigVariables.MongodbUrl == "" {
		ConfigVariables.MongodbUrl = "mongodb://127.0.0.1:27017/"
	}

	validate := validator.New()
	err := validate.Struct(ConfigVariables)

	if err != nil {
		log.Fatal("Environmental variables have not been set!")
	}
	ConfigValues = ConfigVariables

	return &ConfigVariables
}
