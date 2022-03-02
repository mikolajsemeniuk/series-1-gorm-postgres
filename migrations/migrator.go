package migrations

import (
	"database/sql"
	"fmt"
	"gorm-postgres/data"
	"gorm-postgres/entities"
	"gorm-postgres/settings"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type Migrator interface {
	CreateDatabase() error
	DropDatabase() error
	Migrate() error
}

type migrator struct {
	configuration settings.Configuration
	database      *gorm.DB
}

func (migrator *migrator) CreateDatabase() error {
	connectionString := migrator.configuration.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	command := fmt.Sprintf("CREATE DATABASE %s", migrator.configuration.GetDatabaseName())
	_, err = db.Exec(command)
	if err != nil {
		return err
	}

	return err
}

func (migrator *migrator) DropDatabase() error {
	connectionString := migrator.configuration.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	command := fmt.Sprintf(`
		REVOKE CONNECT ON DATABASE %s FROM public;
		SELECT pg_terminate_backend(pg_stat_activity.pid)
		FROM pg_stat_activity
		WHERE pg_stat_activity.datname = '%s';
		DROP DATABASE IF EXISTS %s;`,
		migrator.configuration.GetDatabaseName(),
		migrator.configuration.GetDatabaseName(),
		migrator.configuration.GetDatabaseName())
	_, err = db.Exec(command)
	if err != nil {
		return err
	}

	return err
}

func (migrator *migrator) Migrate() error {
	err := migrator.database.AutoMigrate(&entities.Order{})
	return err
}

func NewMigrator() (Migrator, error) {
	configuration, err := settings.NewConfiguration()
	if err != nil {
		return nil, err
	}

	database, err := data.NewDatabase()
	if err != nil {
		return nil, err
	}

	migrator := &migrator{
		configuration: configuration,
		database:      database,
	}

	return migrator, nil
}
