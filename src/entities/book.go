package entities

import "gorm.io/gorm"

type Book struct {
    gorm.Model
    Name      string `json:"name"`
    YearPublished     string `json:"year_published"`
    Publisher       int    `json:"publisher"`
    NumberCopies int   `json:"number_copies" gorm:"default:true"`
}
