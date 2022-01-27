package surl

import "github.com/devchrischen/url-shortener/entities/edb"

func (s *Service) InsertHash(hash *edb.Hash) error {
	return s.db.Create(hash).Error
}

func (s *Service) GetHash(hash *edb.Hash, hashValue string) error {
	return s.db.Where("value = ?", hashValue).Take(hash).Error
}

func (s *Service) DeleteHash(ID uint64) error {
	return s.db.Where("ID = ?", ID).Delete(edb.Hash{}).Error
}
