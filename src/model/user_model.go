package model

type LoginReqModel struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type LoginResModel struct {
	//LoginStatus bool   `json:"login_status"` // 登录状态：成功(true) or 失败(false)
	Token string `json:"token"`
}

type RegisterReqModel struct {
	Username string `form:"username" binding:"required,numeric"`  // numeric 必须是数字
	Password string `form:"password" binding:"required,alphanum"` // alphanum 必须是数字字母组合
	Name     string `form:"name" binding:"required"`
	Age      int    `form:"age" binding:"required,gte=0,lte=100,numeric"`
	Gender   string `form:"gender" binding:"required,oneof=男 女"`
}
