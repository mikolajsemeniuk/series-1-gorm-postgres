package settings

import (
	"fmt"
	"os"
	"sync"

	"github.com/TwiN/go-color"
	"github.com/spf13/viper"
)

var (
	configurationSingleton configuration
	once                   sync.Once
)

type Configuration interface {
	GetServerBasePath() string
	GetServerPort() string
	GetDatabaseHost() string
	GetDatabaseUsername() string
	GetDatabasePassword() string
	GetDatabaseName() string
	GetDatabasePort() string
}

type configuration struct {
	viper *viper.Viper
}

func (configuration configuration) GetServerBasePath() string {
	return configuration.viper.GetString("server.basepath")
}

func (configuration configuration) GetServerPort() string {
	return configuration.viper.GetString("server.port")
}

func (configuration configuration) GetDatabaseHost() string {
	return configuration.viper.GetString("database.host")
}

func (configuration configuration) GetDatabaseUsername() string {
	return configuration.viper.GetString("database.username")
}

func (configuration configuration) GetDatabasePassword() string {
	return configuration.viper.GetString("database.password")
}

func (configuration configuration) GetDatabaseName() string {
	return configuration.viper.GetString("database.databasename")
}

func (configuration configuration) GetDatabasePort() string {
	return configuration.viper.GetString("database.port")
}

func NewConfiguration() Configuration {
	once.Do(func() {
		environment := getEnvironment()

		viper := viper.New()
		viper.SetConfigType("yaml")
		viper.SetConfigName(environment)
		viper.AddConfigPath("../settings")
		viper.AddConfigPath("settings/")
		viper.ReadInConfig()

		configurationSingleton = configuration{
			viper: viper,
		}
	})

	return configurationSingleton
}

func getEnvironment() string {
	environment := os.Getenv("ENVIROMENT")
	if environment == "" {
		environment = "development"
		message := "\nENVIROMENT variable is not set, setting Enviroment to: development\n"
		fmt.Print(color.Ize(color.Yellow, message))
	}
	return environment
}
