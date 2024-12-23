package config

import (
	"errors"

	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		SecretKey string
		DBConfig
	}

	DBConfig struct {
		DBUser 		 string
		DBPassword   string
		DBHost       string
		DBPort       string
		DBName       string
	}
)

func NewAppConfig() (*AppConfig, error) {
	viper.AutomaticEnv()

	dbUser := viper.GetString("DB_USER")
	if dbUser == "" {
		return nil, errors.New("DB_USER is required")
	}

	dbPassword := viper.GetString("DB_PASSWORD")
	if dbPassword == "" {
		return nil, errors.New("DB_PASSWORD is required")
	}

	dbHost := viper.GetString("DB_HOST")
	if dbHost == "" {
		return nil, errors.New("DB_HOST is required")
	}

	dbPort := viper.GetString("DB_PORT")
	if dbPort == "" {
		return nil, errors.New("DB_PORT is required")
	}

	dbName := viper.GetString("DB_NAME")
	if dbName == "" {
		return nil, errors.New("DB_NAME is required")
	}


	secretKey := viper.GetString("SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("SECRET_KEY is required")
	}

	return &AppConfig{
		SecretKey: secretKey,
		DBConfig: DBConfig{
			DBUser:     dbUser,
			DBPassword: dbPassword,
			DBHost:     dbHost,
			DBPort:     dbPort,
			DBName:     dbName,
		},
	}, nil
}
