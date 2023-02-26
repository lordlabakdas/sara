package main

import (
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lordlabakdas/sara/src/config"
	"github.com/lordlabakdas/sara/src/handlers"
)

func routes(app *fiber.App) {
	app.Get("/get-books", handlers.GetBooks)
	app.Get("/get-user", handlers.GetUser)
	app.Post("/new-user", handlers.NewUser)
	app.Put("/update-user", handlers.UpdateUser)
	app.Delete("/delete-user", handlers.DeleteUser)
	app.Post("/new-book", handlers.NewBook)

}

var (
    WarningLog *log.Logger
    InfoLog   *log.Logger
    ErrorLog   *log.Logger
)

func main() {
    logFile, err := os.OpenFile("log.txt", os.O_CREATE | os.O_APPEND | os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)
	InfoLog = log.New(mw, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    WarningLog = log.New(mw, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
    ErrorLog = log.New(mw, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLog.Println("Sara Library Management System - Under Construction")
	app := fiber.New()
	config.Connect()
	routes(app)
	log.Fatal(app.Listen(":8000"))
}

