package form

// EmailLoginForm 邮箱登录 vo
type LoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterForm struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	Password2 string `json:"password_2" validate:"eqfield=Password"`
}


