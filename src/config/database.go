package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/lordlabakdas/sara/src/db/entities"
)

var envFile string = ".env"

func LoadEnv(file *string) (*string, *string, *string, *string, *string) {
	fmt.Println("Loading Environment Variables")
	err := godotenv.Load(*file)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var dbname string = os.Getenv("DB_NAME")
	var password string = os.Getenv("DB_PASSWORD")
	var port string = os.Getenv("DB_PORT")
	var dbHost string = os.Getenv("DB_HOST")
	var dbUser string = os.Getenv("DB_USER")
	return &dbname, &password, &port, &dbHost, &dbUser
}

func Connect() *gorm.DB {
	dbname, password, port, dbHost, dbUser := LoadEnv(&envFile)
	var databaseURI string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *dbUser, *password, *dbHost, *port, *dbname)
	db, err := gorm.Open("postgres", databaseURI)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Book{})

	return db
}
