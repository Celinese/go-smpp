package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fiorix/go-smpp/v2/Config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//GetAllUsers Fetch all user data
func GetAllSMS(received *[]Received) (err error) {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		return err
	}
	defer db.Close()

	// Use GORM to retrieve all SMS data from the database
	err = db.Find(received).Error
	if err != nil {
		return err
	}

	return nil
}

// ============================================================== //
// ============================================================== //
//       TODO: Function Insert SMS Receiver TO DATABASE ...       //
// ============================================================== //
// ============================================================== //
func InsertDB(address, phone, ucs2Message string) error {

	now := time.Now()
	//db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/connxsmpp")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))

	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	defer Config.DB.Close()

	//TODO: Prepare the insert statement [Function Insert]

	_, err = db.Exec("INSERT INTO received (sender, phone_customer, message_customer, date_received, hand_on) VALUES (?, ?, ?, ?, ?)", address, phone, ucs2Message, now, "API")
	if err != nil {
		return err
	}
	return nil
}
