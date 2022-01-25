package edb

import "time"

type OriginalUrl struct {
	ID        uint64    `gorm:"column:ID;primary_key"`
	HashID    uint64    `gorm:"column:HashID"`
	Url       string    `gorm:"column:Url"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}
