package config

import (
	"fmt"
	"os"
)

type OpenIDCred struct {
	ClientID     string
	ClientSecret string
}

type DatabaseService struct {
	Port string
	Host string
}

type SmsConfig struct {
	APIKey  string
	APIUser string
	Sender  string
}

type Config struct {
	OpenId   OpenIDCred
	Database DatabaseService
	Sms      SmsConfig
}

func LoadConfig() (*Config, error) {
	clientID, clientSecret := os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("missing required environment variables CLIENT_ID and/or CLIENT_SECRET")
	}

	dbPort, dbHost := os.Getenv("DB_PORT"), os.Getenv("DB_HOST")
	if dbPort == "" || dbHost == "" {
		return nil, fmt.Errorf("missing required environment variables DB_PORT and/or DB_HOST")
	}

	smsAPIKey, smsAPIUser, smsSender := os.Getenv("SMS_API_KEY"), os.Getenv("SMS_API_USER"), os.Getenv("SMS_SENDER")
	if smsAPIKey == "" || smsAPIUser == "" || smsSender == "" {
		return nil, fmt.Errorf("missing required environment variables SMS_API_KEY, SMS_API_USER and/or SMS_SENDER")
	}

	cfg := &Config{
		OpenId: OpenIDCred{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
		Database: DatabaseService{
			Port: dbPort,
			Host: dbHost,
		},
		Sms: SmsConfig{
			APIKey:  smsAPIKey,
			APIUser: smsAPIUser,
			Sender:  smsSender,
		},
	}

	return cfg, nil
}
