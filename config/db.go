package Config

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}
type Data struct {
	ID          int
	PhoneNumber string
	Message     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func BuildDBConfig() *DBConfig {

	errorENV := godotenv.Load()
	if errorENV != nil {
		panic("Failed to load env file")
	}

	dbhost := os.Getenv("DB_HOST")
	//dbport := os.Getenv("DB_PORT")
	dbuser := os.Getenv("DB_USER")
	dbpwd := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dbConfig := DBConfig{
		Host:     dbhost,
		User:     dbuser,
		Password: dbpwd,
		DBName:   dbName,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.DBName,
	)
}
