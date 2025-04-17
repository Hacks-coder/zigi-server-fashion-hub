package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConnection *gorm.DB

func connectToDatabase() {

	err := godotenv.Load("./.env")
	if err != nil {
		println("Error in loading env", err)
	}

	DBUSER := os.Getenv("DBUSER")
	DBHOST := os.Getenv("DBHOST")
	DBPORT := os.Getenv("DBPORT")
	DBNAME := os.Getenv("DBNAME")
	DBPASSWORD := os.Getenv("DBPASSWORD")

	connection := DBUSER + ":" + DBPASSWORD + "@tcp(" + DBHOST + ":" + DBPORT + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	log.Panicln("Connected successfully")

	DBConnection = db
}
