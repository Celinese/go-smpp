package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fiorix/go-smpp/v2/Config"
	_ "github.com/go-sql-driver/mysql"
)

/* ================ [Function Submit Insert] ======================== */

func Submitfrominsert(Sender, Phone_To, Message_To string, userName string, userId uint /* , userId int */) error {

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

	defer Config.DB.Close()

	//TODO: Prepare the insert statement [Function Insert]

	_, err = db.Exec("INSERT INTO sendsms (Sender, Phone_To, Message_To, User_name, User_Id, date_insert) VALUES (?, ?, ?, ?, ?, ?)", Sender, Phone_To, Message_To, userName, userId, now)
	if err != nil {
		return err
	}
	return nil
}

/* ================ [Function Submit Insert] ======================== */

func QuerySendLog(userId uint) ([]Sendsms, error) {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))

	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM sendsms WHERE user_id = ?", userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := []Sendsms{}
	for rows.Next() {
		var log Sendsms
		if err := rows.Scan(&log.Id, &log.Sender, &log.Phone_to, &log.Message_to, &log.User_name, &log.Date_insert); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
