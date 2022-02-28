package settings

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
	"github.com/spf13/viper"
)

type Configuration interface {
	GetServerBasePath() string
	GetServerPort() string
}

type configuration struct {
	viper *viper.Viper
}

func (configuration *configuration) GetServerBasePath() string {
	return configuration.viper.GetString("server.basepath")
}

func (configuration *configuration) GetServerPort() string {
	return configuration.viper.GetString("server.port")
}

func NewConfiguration() Configuration {
	environment := os.Getenv("ENVIROMENT")
	if environment == "" {
		environment = "development"
		message := "\nENVIROMENT variable is not set, setting Enviroment to: development\n"
		fmt.Print(color.Ize(color.Yellow, message))
	}

	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.SetConfigName(environment)
	viper.AddConfigPath("../settings")
	viper.AddConfigPath("settings/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(color.Ize(color.Red, err.Error()))
	}

	return &configuration{
		viper: viper,
	}
}
