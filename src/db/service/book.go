package service

import (
	"github.com/lordlabakdas/sara/src/config"
	"github.com/lordlabakdas/sara/src/db/entities"
)

func GetBook(name string, author string) []string {
	var book entities.Book
	db := config.Connect()
	db.First(&book, "name = ?", name, "author = ?", author)
	return []string{book.Name, book.Author, book.YearPublished}

}
