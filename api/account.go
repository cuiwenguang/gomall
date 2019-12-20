package api

import (
	"gomall/models/vo"
	"gomall/pkg/e"
	"gomall/pkg/web"
	"gomall/service"
	"gopkg.in/go-playground/validator.v9"
)

func Login(ctx *web.Context) {
	form := &vo.LoginForm{}
	if err := ctx.BindJSON(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewAccountService(ctx.RequestContext)
	token, code := srv.LoginByEmail(form.Email, form.Password)
	ctx.ResponseData(code, token)

}

func Register(ctx *web.Context) {
	form := &vo.RegisterForm{}
	if err := ctx.BindJSON(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	validate := validator.New()
	if err := validate.Struct(form); err != nil {
		ctx.Response(e.BAD_REQUEST)
		return
	}
	srv := service.NewAccountService(ctx.RequestContext)
	code := srv.RegisterByEmail(form.Email, form.Password)
	ctx.Response(code)
}
