package model

type LoginReqModel struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResModel struct {
	//LoginStatus bool   `json:"login_status"` // 登录状态：成功(true) or 失败(false)
	Token string `json:"token"`
}

type RegisterReqModel struct {
	Username string `json:"username" binding:"required,numeric"`  // numeric 必须是数字
	Password string `json:"password" binding:"required,alphanum"` // alphanum 必须是数字字母组合
	Name     string `json:"name" binding:"required"`
	Age      int    `json:"age" binding:"required,gte=0,lte=100,numeric"`
	Gender   string `json:"gender" binding:"required,oneof=男 女"`
}
