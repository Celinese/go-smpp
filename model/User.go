package model

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var err error
var db *gorm.DB

func GetUserByUsername(Email string) (*User, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName))

	if err != nil {
		return nil, err
	}
	defer db.Close()

	user := &User{}
	if err := db.Where("Email = ?", Email).First(user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
