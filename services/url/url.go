package surl

import "github.com/devchrischen/url-shortener/entities/edb"

func (s *Service) InsertUrl(url *edb.OriginalUrl) error {
	return s.db.Create(url).Error
}
