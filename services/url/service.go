package surl

import "github.com/jinzhu/gorm"

type Service struct {
	db *gorm.DB
}

func New(conn *gorm.DB) *Service {
	return &Service{
		db: conn,
	}
}
