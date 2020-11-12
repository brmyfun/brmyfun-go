package form

// RegisterDefault 默认注册表单
type RegisterDefault struct {
	Telephone        string `form:"telephone" json:"telephone" binding:"required"`
	VerificationCode string `form:"verificationCode" json:"verificationCode" binding:"required"`
}
