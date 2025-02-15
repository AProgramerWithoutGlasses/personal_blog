package response

type Code int //错误码

const SuccessCode Code = 200

const (
	// ServerError 1 开头为服务级错误码
	_            Code = iota + 10000 // 让 iota 从 10001 开始
	ErrServer                        // 10001	服务内部错误
	ErrParamBind                     // 10002	请求参数错误
	ErrToken                         // 10003 token错误
)

const (
	// ILLegaLDatasetName 2 开头为业务级错误码

	// 其中数据集管理为 201 开头
	_                  Code = iota + 20100
	IllegalDatasetName      // 20101 无效的数据集名称
	ParamNameError          // 20102 参数 name 错误
)

const (
	// ILLegaLPhoneNum 用户管理模块：202 开头
	_                  Code = iota + 20200
	ErrUserExisted          //用户已存在
	ErrUserNotExist         //用户不存在
	ErrInvalidPassword      //密码错误
)

// 定义errorCode对应的文本信息
var errorMsg = map[Code]string{
	SuccessCode:        "操作成功",
	ErrServer:          "服务内部错误",
	ErrParamBind:       "请求参数错误",
	ErrUserExisted:     "用户已存在",
	ErrUserNotExist:    "用户不存在",
	ErrInvalidPassword: "密码错误",
	IllegalDatasetName: "无效的数据集名称",
	ParamNameError:     "参数name错误",
	ErrToken:           "token解析错误",
}

// 用于获取code对应的提示信息
func (c Code) Msg() string {
	return errorMsg[c]
}

// NewCustomError 新建自定义error实例化
//func NewCustomError(code  Code) error {
//	// 初次调用得用Wrap方法，进行实例化
//	return errors.Wrap(&Response{
//		Code: code,
//		Msg:  code.String(),
//	}, "")
//}
