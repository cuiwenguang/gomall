package entity

import "github.com/jinzhu/gorm"

type Account struct {
	gorm.Model
	Username  string `gorm:"type:varchar(50)"`
	Password  string `gorm:"type:varchar(50);unique"`
	Email     string `gorm:"unique_index;type:varchar(50)"`
	Mobile    string `gorm:"unique_index;type:varchar(20)"`
	UserType  string `gorm:"type:char(10);unique"`
	CompanyID int
}
