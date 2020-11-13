package form

// RegisterDefault 默认注册表单
type RegisterDefault struct {
	Username         string `form:"username" json:"username" binding:"required"`
	Password         string `form:"password" json:"password" binding:"required"`
	Email            string `form:"email" json:"email" binding:"required"`
	VerificationCode string `form:"verificationCode" json:"verificationCode" binding:"required"`
}
