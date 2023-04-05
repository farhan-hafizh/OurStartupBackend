package config

import (
	"github.com/spf13/viper"
)

// mapstructure: for viper
type Config struct {
	DBConnection     string `mapstructure:"DB_CONNECTION"`
	EncryptionSecret string `mapstructure:"ENCRYPTION_SECRET"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	Port             string `mapstructure:"PORT"`
	MidClient        string `mapstructure:"MIDTRANS_CLIENT_KEY"`
	MidServer        string `mapstructure:"MIDTRANS_SERVER_KEY"`
}

func LoadConfig(mode string) (config Config, err error) {

	viper.AddConfigPath("./config/")

	if mode == "development" {
		viper.SetConfigName("development")
	} else if mode == "production" {
		viper.SetConfigName("development")
	}

	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
