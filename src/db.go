package main

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	/* dsn := "stan:j@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("fail to init db")
	} */
}
