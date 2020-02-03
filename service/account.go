package service

import (
	"gomall/models"
	"gomall/pkg/e"
	"gomall/pkg/token"
	"gomall/pkg/util"
	"gomall/pkg/web"
)

type AccountService struct {
	Service
}

func NewAccountService(context *web.RequestContext) *AccountService {
	a := &AccountService{
		InitService(context),
	}
	return a
}

func (a *AccountService) RegisterByEmail(email, password string) int {
	account := &models.Account{
		Email:    email,
		Password: util.EncodeMD5(password),
	}
	if a.ExistEmail(email) {
		return e.ERROR_EXIST_USER
	}
	if err := a.Create(account); err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

func (a *AccountService) LoginByEmail(email, password string) (string, int) {
	account := a.GetAccountByEmail(email)
	if account == nil {
		return "", e.ERROR_USER_OR_PASSWORD
	}
	if util.EncodeMD5(password) != account.Password {
		return "", e.ERROR_USER_OR_PASSWORD
	}
	tokenStr, _ := token.Generate(a.RequestContext.Host, email, password)
	return tokenStr, e.SUCCESS
}

func (a *AccountService) GetAccountByEmail(email string) *models.Account {
	account := &models.Account{}
	a.DB.Where("email = ?", email).First(account)
	return account
}

func (a *AccountService) ExistEmail(email string) bool {
	result := 0
	a.DB.Select("id").Where("email = ?", email).Count(result)
	return result > 0
}

// Create 创建用户
func (a *AccountService) Create(account *models.Account) error {
	if err := a.DB.Create(account).Error; err != nil {
		return nil
	}
	return nil
}
