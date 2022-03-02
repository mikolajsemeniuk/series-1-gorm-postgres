package repositories

import (
	"gorm-postgres/data"

	"gorm.io/gorm"
)

type Order interface{}

type order struct {
	database *gorm.DB
}

func NewOrder() (Order, error) {
	database, err := data.NewDatabase()
	if err != nil {
		return nil, err
	}

	order := &order{
		database: database,
	}

	return order, nil
}
