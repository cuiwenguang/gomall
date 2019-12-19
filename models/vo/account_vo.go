package vo

type AccountVo struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Mobile      string `json:"mobile"`
	UserType    string `json:"user_type"`
	CompanyID   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
}
