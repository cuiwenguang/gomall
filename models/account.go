package models

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

// EmailLoginForm 邮箱登录表单
type LoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// 邮箱注册表单
type RegisterForm struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Password2 string `json:"password_2" validate:"eqfield=Password"`
}

// 账号信息Vo
type AccountVo struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	UserType    string `json:"user_type"`
	CompanyID   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
}
