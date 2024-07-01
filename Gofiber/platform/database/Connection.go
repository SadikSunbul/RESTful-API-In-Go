package database

import (
	"gorm.io/gorm"
	"sync"
)

var once sync.Once
var connection *gorm.DB

func Connection() *gorm.DB {
	once.Do(func() {
		connection = initialize()
	})

	return connection
}

func initialize() *gorm.DB {
	return initializePostgres()
}
