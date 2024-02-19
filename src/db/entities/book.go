package entities

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name          string `json:"name"`
	Author        string `json:"author"`
	YearPublished string `json:"year_published"`
	Publisher     int    `json:"publisher"`
	NumberCopies  int    `json:"number_copies"`
}
