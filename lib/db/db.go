package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/devchrischen/url-shortener/config"
)

var DB *gorm.DB

func Init() {
	var err error
	dbConfig := config.Config.DB
	dsn := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?%v",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Protocol,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
		dbConfig.Params,
	)
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

}
