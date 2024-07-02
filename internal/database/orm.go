package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Document struct {
	gorm.Model
	TicketID  string `gorm:"type:varchar(100);unique_index"`
	Content   string `gorm:"type:varchar(100)"`
	Title     string `gorm:"type:varchar(100)"`
	Author    string `gorm:"type:varchar(100)"`
	Topic     string `gorm:"type:varchar(100)"`
	Watermark string `gorm:"type:varchar(100)"`
}

func Init(host, port, user, dbname, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.DB().Ping(); err != nil {
		return nil, err
	}

	return db, nil
}