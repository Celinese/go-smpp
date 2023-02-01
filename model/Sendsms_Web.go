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

func Submitfrominsert(Sender, PhoneTo, MessageTo string, userName string, userId uint /* , userId int */) error {

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

	_, err = db.Exec("INSERT INTO sendsms (Sender, Phone_To, Message_To, User_name, User_Id, date_insert, hand_on) VALUES (?, ?, ?, ?, ?, ?, ?)", Sender, PhoneTo, MessageTo, userName, userId, now, "WeB")
	if err != nil {
		return err
	}
	return nil
}

/* ================ [Function Query Send Log] ======================== */

func QuerySendLogs(userId uint) ([]Sendsms, error) {

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
		var User_id string
		var Date_insertString string
		if err := rows.Scan(&log.Id, &log.Sender, &log.Phone_to, &log.Message_to, &log.User_name, &Date_insertString, &log.User_id, &log.Hand_on); err != nil {
			return nil, err
		}

		Date_insert, err := time.Parse("2006-01-02 15:04:05", Date_insertString)
		if err != nil {
			log.Fatal(err)
		}

		log.User_id = User_id
		log.Date_insert = Date_insert
		logs = append(logs, log)
	}
	return logs, nil
}
