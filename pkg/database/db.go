package database

import (
	"context"
	"net/http"
    "fmt"
	"github.com/aayushrangwala/watermark-service/internal"
	"github.com/jinzhu/gorm"
)

type dbService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &dbService{db: db}
}

func (d *dbService) Add(_ context.Context, doc *internal.Document) (string, error) {
	fmt.Print("doc-value")
	fmt.Print(doc)
	if err := d.db.Create(doc).Error; err != nil {
		return "", err
	}
	return doc.Title, nil
}

func (d *dbService) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
	var documents []internal.Document
	query := d.db
    
	for _, filter := range filters {
		if filter.Value != "" {
			fmt.Printf("filtering by Key and Value")
			fmt.Print(filter.Key)
			fmt.Print(filter.Value)
			query = query.Where(filter.Key+" = ?", filter.Value)
		} else {
			query = query.Order(filter.Key)
		}
	}

	if err := query.Find(&documents).Error; err != nil {
		return nil, err
	}
	return documents, nil
}

func (d *dbService) Update(_ context.Context, title string, doc *internal.Document) (int, error) {
	if err := d.db.Model(&internal.Document{}).Where("title = ?", title).Updates(doc).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (d *dbService) Remove(_ context.Context, title string) (int, error) {
	if err := d.db.Where("title = ?", title).Delete(&internal.Document{}).Error; err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (d *dbService) ServiceStatus(_ context.Context) (int, error) {
	return http.StatusOK, nil
}
