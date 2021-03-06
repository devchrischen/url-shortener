package edb

import "time"

type Hash struct {
	ID        uint64    `gorm:"column:ID;primary_key"`
	Value     string    `gorm:"column:Value"`
	CreatedAt time.Time `gorm:"column:CreatedAt"`
	UpdatedAt time.Time `gorm:"column:UpdatedAt"`
}

// TableName sets the insert table name for this struct type
func (b *Hash) TableName() string {
	return "Hash"
}
