package forms

type PasswordLoginForm struct{
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Captcha string `form:"captcha" json:"captcha" binding:"required"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=20"`
}
