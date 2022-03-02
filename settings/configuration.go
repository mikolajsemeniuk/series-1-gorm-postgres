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
	configurationOnce      sync.Once
	err                    error
)

type Configuration interface {
	GetServerBasePath() string
	GetServerPort() string
	GetConnectionString() string
	GetDatabaseConnectionString() string
	GetDatabaseHost() string
	GetDatabaseUsername() string
	GetDatabasePassword() string
	GetDatabaseName() string
	GetDatabasePort() string
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

func (configuration *configuration) GetConnectionString() string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		configuration.GetDatabaseHost(),
		configuration.GetDatabaseUsername(),
		configuration.GetDatabasePassword(),
		configuration.GetDatabasePort())

	return connectionString
}

func (configuration *configuration) GetDatabaseConnectionString() string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		configuration.GetDatabaseHost(),
		configuration.GetDatabaseUsername(),
		configuration.GetDatabasePassword(),
		configuration.GetDatabaseName(),
		configuration.GetDatabasePort())
	return connectionString
}

func (configuration *configuration) GetDatabaseHost() string {
	return configuration.viper.GetString("database.host")
}

func (configuration *configuration) GetDatabaseUsername() string {
	return configuration.viper.GetString("database.username")
}

func (configuration *configuration) GetDatabasePassword() string {
	return configuration.viper.GetString("database.password")
}

func (configuration *configuration) GetDatabaseName() string {
	return configuration.viper.GetString("database.databasename")
}

func (configuration *configuration) GetDatabasePort() string {
	return configuration.viper.GetString("database.port")
}

func NewConfiguration() (Configuration, error) {
	configurationOnce.Do(func() {
		environment := getEnvironment()

		viper := viper.New()
		viper.SetConfigType("yaml")
		viper.SetConfigName(environment)
		viper.AddConfigPath("../settings")
		viper.AddConfigPath("settings/")
		err = viper.ReadInConfig()

		configurationSingleton = configuration{
			viper: viper,
		}
	})
	return &configurationSingleton, err
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
