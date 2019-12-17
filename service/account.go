package service

import (
	"gomall/models"
	"gomall/pkg/e"
	"gomall/pkg/token"
	"gomall/pkg/util"
)

func RegisterByEmail(email, password string) int  {
	account :=models.Account{
		Email:    email,
		Password: util.EncodeMD5(password),
	}
	if account.ExistEmail(){
		return e.ERROR_EXIST_USER
	}
	if err := account.Add(); err != nil{
		return e.ERROR
	}
	return e.SUCCESS
}

func LoginByEmail(email, password string) (string, int)  {
	account := models.GetAccountByEmail(email)
	if account == nil {
		return "", e.ERROR_USER_OR_PASSWORD
	}
	if util.EncodeMD5(password) != account.Password {
		return "", e.ERROR_USER_OR_PASSWORD
	}
	token,_ := token.Generate("localhost", email, password)
	return token, e.SUCCESS
}
