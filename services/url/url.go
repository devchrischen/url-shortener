package surl

import "github.com/devchrischen/url-shortener/entities/edb"

func (s *Service) InsertUrl(url *edb.OriginalUrl) error {
	return s.db.Create(url).Error
}

func (s *Service) GetUrl(url *edb.OriginalUrl, HashID uint64) error {
	return s.db.Where("HashId = ?", HashID).Take(url).Error
}

func (s *Service) DeleteUrl(HashID uint64) error {
	return s.db.Where("HashId = ?", HashID).Delete(edb.OriginalUrl{}).Error
}
