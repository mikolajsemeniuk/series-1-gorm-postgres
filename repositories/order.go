package repositories

import (
	"gorm-postgres/data"

	"gorm.io/gorm"
)

type Order interface{}

type order struct {
	database *gorm.DB
}

func NewOrder() Order {
	database := data.NewDatabase()

	return &order{
		database: database,
	}
}
