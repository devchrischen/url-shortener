package surl

import "github.com/devchrischen/url-shortener/entities/edb"

func (s *Service) InsertUrl(url *edb.OriginalUrl) error {
	return s.db.Create(url).Error
}

func (s *Service) GetUrlRecord(url *edb.OriginalUrl, HashID uint64) error {
	return s.db.Where("HashId = ?", HashID).Take(url).Error
}
