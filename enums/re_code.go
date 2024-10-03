package enums

// ErrorCode 接口，包含获取错误码和描述的方法
type ErrorCode interface {
	GetCode() int
	GetDesc() string
	SetDesc(string)
}

// ReCode 结构体，包含状态码和消息
type ReCode struct {
	Code    int
	Message string
}

// 构造函数
func NewReCode(code int, message string) ReCode {
	return ReCode{
		Code:    code,
		Message: message,
	}
}

// 实现 ErrorCode 接口的 GetCode 方法
func (r ReCode) GetCode() int {
	return r.Code
}

// 实现 ErrorCode 接口的 GetDesc 方法
func (r ReCode) GetDesc() string {
	return r.Message
}

// 实现 ErrorCode 接口的 SetDesc 方法
func (r *ReCode) SetDesc(desc string) {
	r.Message = desc
}

// 定义 ReCode 常量
var (
	SUCCESS               = NewReCode(200, "Success")
	NO_PERMISSION         = NewReCode(403, "No Permission")
	SERVER_ERROR          = NewReCode(500, "Internal Server Error")
	INVALID_PARAMETERS    = NewReCode(501, "Invalid Parameters")
	INVALID_TOKEN         = NewReCode(502, "Invalid Token")
	TOKEN_EXPIRED         = NewReCode(502, "The token has expired")
	FAILED                = NewReCode(503, "Failed")
	DATA_DUPLICATION      = NewReCode(505, "Data exists")
	VERIFICATION_FAILED   = NewReCode(506, "Wrong account or password.")
	REQUEST_TOO_FREQUENT  = NewReCode(509, "Requests are too frequent")
	PRIVATE_KEY_EXIST     = NewReCode(502, "exist")
	PRIVATE_KEY_NOT_EXIST = NewReCode(503, "key not exist")
)
