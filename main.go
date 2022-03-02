package main

import (
	"fmt"
	"gorm-postgres/application"
	"gorm-postgres/migrations"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	migrator, err := migrations.NewMigrator()
	if err != nil {
		os.Exit(1)
	}

	migrator.DropDatabase()
	migrator.CreateDatabase()

	application, err := application.New()
	if err != nil {
		// TODO:
		fmt.Println(err.Error())
		os.Exit(1)
	}

	application.Listen()
}
