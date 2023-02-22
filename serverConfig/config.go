package serverConfig

import (
	"github.com/spf13/viper"
)

// mapstructure: for viper
type Config struct {
	DBConnection     string `mapstructure:"DB_CONNECTION"`
	EncryptionSecret string `mapstructure:"ENCRYPTION_SECRET"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	Port             string `mapstructure:"PORT"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./serverConfig/")
	viper.SetConfigName("development")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
