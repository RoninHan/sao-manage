package forms

import "time"

type PasswordLoginForm struct {
	UserName  string `form:"name" json:"name"`
	Mobile    string `form:"mobile" json:"mobile"` //手机号码格式有规范可寻， 自定义validator
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha" ` // 验证码 binding:"required,min=5,max=5"
	CaptchaId string `form:"captcha_id" json:"captcha_id" `       // 验证码id binding:"required"
}



type UserListForm struct {
	UserGuid string
	Password string
	UserName string
	Head_url string
	Birthday string
	Address string
	Gender string
	mobile string
	Current int
	PageSize int
}

type UserForm struct {
	UserGuid string
	Password string
	UserName string
	HeadUrl  string
	Birthday *time.Time
	Address  string
	Gender   string
	Role     int
	Mobile   string
}