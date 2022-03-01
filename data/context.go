package data

import (
	"fmt"
	"gorm-postgres/settings"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	databaseSingleton *gorm.DB
	once              sync.Once
)

func NewDatabase() *gorm.DB {
	once.Do(func() {
		configuration := settings.NewConfiguration()

		connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			configuration.GetDatabaseHost(),
			configuration.GetDatabaseUsername(),
			configuration.GetDatabasePassword(),
			configuration.GetDatabaseName(),
			configuration.GetDatabasePort())

		config := &gorm.Config{}
		databaseSingleton, _ = gorm.Open(postgres.Open(connectionString), config)
	})
	return databaseSingleton
}
