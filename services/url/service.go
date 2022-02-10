package surl

import (
	"github.com/jinzhu/gorm"

	"github.com/devchrischen/url-shortener/entities/edb"
)

type Service struct {
	db *gorm.DB
}

func New(conn *gorm.DB) *Service {
	return &Service{
		db: conn,
	}
}

func (s *Service) CreateUrlData(hashValue, originalUrl string) error {
	tx := s.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	createService := &Service{
		db: tx,
	}
	hash := edb.Hash{
		Value: hashValue,
	}
	if err := createService.InsertHash(&hash); err != nil {
		tx.Rollback()
		return err
	}
	url := edb.OriginalUrl{
		HashID: hash.ID,
		Url:    originalUrl,
	}
	if err := createService.InsertUrl(&url); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
