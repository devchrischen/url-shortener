package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/devchrischen/url-shortener/entities/edb"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("mysql", "chrischen:funnow@tcp(localhost:3306)/Url_Shortener?charset=utf8mb4,utf8&parseTime=True")
	if err != nil {
		panic(err)
	}

	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&edb.OriginalUrl{}, &edb.Hash{})
}
