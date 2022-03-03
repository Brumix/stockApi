package repository

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"stockApi/errors"
	"strconv"
)

func InitConnection() *gorm.DB {

	host := os.Getenv("HOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port, _ := strconv.Atoi(os.Getenv("DBPORT"))

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Europe/Lisbon", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errors.ErrorRepository("Error connecting with database", err)

	return db
}

func init() {
	migrations()
}

func migrations() {
	db := InitConnection()
	defer func() {
		var dbC, _ = db.DB()
		errors.ErrorRepository("[REPOSITORY] ERROR closing the database", dbC.Close())
	}()
	//errors.CheckError("[REPOSITORY] Fail executing USER Migrations", db.AutoMigrate(&models.User{}))

	log.Debug("[REPOSITORY] AutoMigration Complete!!")
}
