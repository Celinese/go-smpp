package model_test

/* import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fiorix/go-smpp/v2/Config"
	"github.com/gin-gonic/gin"
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

// TODO: INSERT INTO Received
func insert(db *sql.DB, p Received) error {
	query := "INSERT INTO product(product_name, product_price) VALUES (?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return err
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, p.Address, p.Phone, p.Message)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return err
	}
	log.Printf("%d products created ", rows)
	return nil
}

// ============================================================== //
// ============================================================== //
//       TODO: Function Insert SMS Receiver TO DATABASE ...       //
// ============================================================== //
// ============================================================== //
func InsertDB(address, phone, ucs2Message string) error {

	// ! [ LOAD ENV SETTINGS FILE] //

	now := time.Now()
	//db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/connxsmpp")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	//Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	defer Config.DB.Close()

	//TODO: Prepare the insert statement [Function Insert]

	_, err = db.Exec("INSERT INTO received (address, phone, message, date) VALUES (?, ?, ?, ?)", address, phone, ucs2Message, now)
	if err != nil {
		return err
	}
	return nil
}

// ============================================================== //
// ============================================================== //
//           TODO: Function Insert SMS SEND TO DATABASE ...       //
// ============================================================== //
// ============================================================== //
func SendInsert(Sender, Phone_Number, Message string) error {

	// ! [ LOAD ENV SETTINGS FILE] //

	now := time.Now()
	//db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/connxsmpp")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	//Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	defer Config.DB.Close()

	//TODO: Prepare the insert statement [Function Insert]

	_, err = db.Exec("INSERT INTO sendsms (Sender, Phone_Number, Message, date) VALUES (?, ?, ?, ?)", Sender, Phone_Number, Message, now)
	if err != nil {
		return err
	}
	return nil
}

//GetAllUsers Fetch all user data
func GetsmsSend(Sendsms *[]Sendsms) (err error) {
	//%s:%s@tcp(%s:%s)/%s
	//db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/connxsmpp?charset=utf8&parseTime=True&loc=Local")
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
	err = db.Find(Sendsms).Error
	if err != nil {
		return err
	}

	return nil
}

// ============================================================== //
// ============================================================== //
//           TODO: Function GET ITEMS JOIN TABLE ...              //
// ============================================================== //
// ============================================================== //

func GetItemtest(c *gin.Context) {
	phone := c.Param("phone")

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer db.Close()

	var received []Received
	//SELECT S.phone,S.content,S.datetime FROM logs_receive S JOIN logs_send C ON S.phone = C.phone WHERE UserID = '$UserID'
	//err = db.QueryRow("SELECT i.name, c.name FROM items i JOIN categories c ON i.category_id = c.id WHERE i.id = ?", id).Scan(&item.Name, &item.Category)
	//err = db.QueryRow("SELECT S.phone, S.message, S.date FROM receivd S JOIN sendsms c ON S.phone = c.phone_number WHERE S.phone = ?", phone).Scan(&i.phone, &i.message, &i.date) received.phone = sendsms.phone_number and
	err = db.Table("received").Joins("JOIN sendsms ON received.phone = sendsms.phone_number").Where("received.phone = ?", phone).Select("DISTINCT received.phone,received.id, received.message, received.date").Find(&received).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, received)
}
*/
