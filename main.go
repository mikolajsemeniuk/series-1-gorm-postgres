package main

import (
	"gorm-postgres/application"
)

func main() {
	application.New().Listen()
}
