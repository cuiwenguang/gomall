package api

import (
	"gomall/api/form"
	"gomall/pkg/e"
	"gomall/pkg/web"
	"gomall/service"
	"gopkg.in/go-playground/validator.v9"
)


func Login(ctx *web.Context)  {
	form := &form.LoginForm{}
	if err := ctx.BindJSON(form); err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	token, code := service.LoginByEmail(form.Email, form.Password)
	ctx.ResponseData(code, token)

}

func Register(ctx *web.Context)  {
	form := &form.RegisterForm{}
	if err := ctx.BindJSON(form); err != nil{
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	code := service.RegisterByEmail(form.Email,form.Password)
	ctx.Response(code)
}