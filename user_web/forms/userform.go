package forms

type PasswordLoginForm struct{
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha string `form:"captcha" json:"captcha" binding:"required"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}
