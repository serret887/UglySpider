package models

import (
	"github.com/jinzhu/gorm"
)

// Cars is the DB model that we have in the
// first page of the ccraighlist page
type Cars struct {
	gorm.Model
	Name  string
	place string
	price float32
}
