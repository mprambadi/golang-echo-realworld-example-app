package model

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Status   bool
	AuthorID uint
}
