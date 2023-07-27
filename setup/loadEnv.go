package setup

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DBMS_HOST"`
	DBPort     string `mapstructure:"DBMS_PORT"`
	DBUserName string `mapstructure:"DBMS_USERNAME"`
	DBPassword string `mapstructure:"DBMS_PASSWORD"`

	DBName string `mapstructure:"DB_NAME"`

	ServerPort string `mapstructure:"SERVER_PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	Expiration time.Duration `mapstructure:"SESSION_EXPIRES_IN"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
