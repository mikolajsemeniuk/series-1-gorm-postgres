package data

import (
	"gorm-postgres/settings"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	databaseSingleton *gorm.DB
	databaseOnce      sync.Once
	err               error
)

func NewDatabase() (*gorm.DB, error) {
	databaseOnce.Do(func() {
		var configuration settings.Configuration
		configuration, err = settings.NewConfiguration()
		if err != nil {
			return
		}

		connectionString := configuration.GetDatabaseConnectionString()
		config := &gorm.Config{}
		databaseSingleton, err = gorm.Open(postgres.Open(connectionString), config)
	})
	return databaseSingleton, err
}
