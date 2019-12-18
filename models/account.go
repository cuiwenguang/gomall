package models

type Account struct {
	Model
	Username string `gorm:"type:varchar(50)"`
	Password string `gorm:"type:varchar(50);unique"`
	Email    string `gorm:"unique_index;type:varchar(50)"`
	Mobile   string `gorm:"unique_index;type:varchar(20)"`
	UserType string `gorm:"type:char(10);unique"`
}

func GetAccountByEmail(email string) *Account {
	account := &Account{}
	db.Where("email = ?", email).First(account)
	return account
}

func (a *Account) ExistEmail() bool {
	result := 0
	db.Select("id").Where("email = ?", a.Email).Find(a).Count(&result)
	return result > 0
}

// Add 创建用户
func (a *Account) Add() error {
	if err := db.Create(a).Error; err != nil {
		return nil
	}
	return nil
}
