package surl

import (
	"github.com/jinzhu/gorm"

	"github.com/devchrischen/url-shortener/entities/edb"
)

func (s *Service) InsertUrl(url *edb.OriginalUrl) error {
	return s.db.Create(url).Error
}

func (s *Service) GetUrl(url *edb.OriginalUrl, HashID uint64) error {
	return s.db.Where("HashId = ?", HashID).Take(url).Error
}

func (s *Service) DeleteUrl(HashID uint64) error {
	return s.db.Where("HashId = ?", HashID).Delete(edb.OriginalUrl{}).Error
}

func (s *Service) CheckUrlExist(urlStr string) (string, bool, error) {
	urlInstance := &edb.OriginalUrl{}
	err := s.db.Where("Url = ?", urlStr).Take(&urlInstance).Error
	var notFound bool
	if err != nil {
		if notFound = gorm.IsRecordNotFoundError(err); notFound {
			// url record notFoundError
			return "", false, nil
		} else {
			// other db error
			return "", false, err
		}
	}
	hashInstance := &edb.Hash{}
	err = s.db.Where("ID = ?", urlInstance.HashID).Take(&hashInstance).Error
	if err != nil {
		// hash record notFoundError or other db error
		return "", true, err
	}
	return hashInstance.Value, true, nil

}
