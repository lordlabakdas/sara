package config

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"github.com/lordlabakdas/sara/src/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

var DATABASE_URI string = "root:root@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"


type person struct {
	name string
	age  int
  }

func LoadEnv(env_file *string) (*string, *string, *string) {
	fmt.Println("Loading Environment Variables")
	err := godotenv.Load(*env_file)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var username string = os.Getenv("USERNAME")
	var password string = os.Getenv("PASSWORD")
	var url string= os.Getenv("URL")
	return &username, &password, &url
}

func Connect() error {
	env_file := "local.env"
	username, password, url := LoadEnv(&env_file)
	var DATABASE_URI string = ("root:root@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local", username, password, url)
	var err error
    Database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })

    if err != nil {
        panic(err)
    }

    Database.AutoMigrate(&entities.Book{})

    return nil
}
