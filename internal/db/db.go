package db

import (
	"encoding/gob"

	"github.com/intrigues/golang-starter-boilerplate/core"
	"github.com/intrigues/golang-starter-boilerplate/internal/config"
	"github.com/intrigues/golang-starter-boilerplate/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	Conn *gorm.DB
}

func NewDBConnection(cfg config.Config, logger core.Clog) (*DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// migrating to database
	db.AutoMigrate(&models.Users{})
	gob.Register(models.Users{})

	return &DB{Conn: db}, nil
}
