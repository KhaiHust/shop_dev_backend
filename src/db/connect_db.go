package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"shop_dev/config"
)

type IInitDatabase interface {
}

func NewInitDatabase(config *config.Config) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.DatabaseName,
	)
	db := mysql.Open(url)
	gormDB, err := gorm.Open(db, &gorm.Config{})
	if err != nil {
		log.Fatal(fmt.Sprintf("Can not connect db with err %v"), err)
		return nil, err
	}
	return gormDB, nil
}
