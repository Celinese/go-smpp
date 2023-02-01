package model

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

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
