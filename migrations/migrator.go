package migrations

import (
	"gorm-postgres/data"

	"gorm.io/gorm"
)

type Migrator interface{}

type migrator struct {
	database *gorm.DB
}

func NewMigrator() (Migrator, error) {
	database := data.NewDatabase()

	migrator := &migrator{
		database: database,
	}

	return migrator, nil
}
