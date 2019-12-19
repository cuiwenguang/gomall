package entity

import (
	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name string `json:"name"`
}
