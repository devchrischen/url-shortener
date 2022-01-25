package edb

import "time"

type Hash struct {
	ID        uint64    `gorm:"column:ID;primary_key"`
	Value     string    `gorm:"column:value"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}
