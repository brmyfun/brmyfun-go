package form

// LoginDefault 默认登录表单
type LoginDefault struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
