package models

type Account struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	UserType string `json:"user_type"`
}

func GetAccountByEmail(email string) *Account{
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
func (a *Account) Add() error  {
	if err:= db.Create(a).Error; err != nil{
		return nil
	}
	return nil
}

