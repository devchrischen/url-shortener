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

func (s *Service) InsertUrlData(hashValue, originalUrl string) error {
	tx := s.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	defer tx.Rollback()
	insertService := &Service{
		db: tx,
	}
	hash := edb.Hash{
		Value: hashValue,
	}
	if err := insertService.InsertHash(&hash); err != nil {
		return err
	}
	url := edb.OriginalUrl{
		HashID: hash.ID,
		Url:    originalUrl,
	}
	if err := insertService.InsertUrl(&url); err != nil {
		return err
	}
	return tx.Commit().Error
}
