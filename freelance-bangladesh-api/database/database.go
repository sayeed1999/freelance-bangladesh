package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sayeed1999/freelance-bangladesh/config"
	"github.com/sayeed1999/freelance-bangladesh/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

// Connect function
func Connect() {
	dbConfig := config.GetConfig().Database

	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(dbConfig.Port, 10, 32)
	if err != nil {
		fmt.Println("Error parsing port str to int")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations on database")

	if err = db.AutoMigrate(
		&models.Talent{},
		&models.Client{},
		&models.Job{},
		&models.Bid{},
		&models.Assignment{},
		&models.Review{},
	); err != nil {
		log.Fatal(err)
	}

	DB = DBInstance{
		Db: db,
	}
}
