package service

import ("github.com/lordlabakdas/sara/src/db/entities"
		"github.com/lordlabakdas/sara/src/config")


func GetBook(name string, author string) []string {
	var book entities.Book
	db:= config.Connect()
    db.First(&book, "name = ?", name, "author = ?", author)

}


func NewBook() error {
	return c.SendString("New Book")
}