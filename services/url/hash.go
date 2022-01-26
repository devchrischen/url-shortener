package surl

import "github.com/devchrischen/url-shortener/entities/edb"

func (s *Service) InsertHash(hash *edb.Hash) error {
	return s.db.Create(hash).Error
}

func (s *Service) GetHashRecord(hash *edb.Hash, hashValue string) error {
	return s.db.Where("value = ?", hashValue).Take(hash).Error
}
