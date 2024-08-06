package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("FhLk:kmutT@631305.78@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	//_, err := gorm.Open(mysql.Open(user+":"+pass+"@tcp("+url+")/"+db+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("Connected!")
	return db
}
