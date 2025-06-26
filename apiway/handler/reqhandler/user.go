package reqhandler

type LoginForm struct {
	Mobile  string `form:"mobile" binding:"required"`
	SendSms string `form:"send_sms" binding:"required"`
}
