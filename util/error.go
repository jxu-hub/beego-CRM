package util

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ErrCodeMsgMap = map[int]string{
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	500: "Internal Server Error",
	502: "用户名或者密码错误",
	504: "请求解析失败",
	505: "密码加密失败",
	//继续添加其他的错误
	506: "获取参数失败",
	507: "客户信息插入失败",
	508: "联系人信息插入失败",
	509: "客户姓名插入失败",
	510: "获取参数失败",
	511: "查询失败",
	512: "添加失败",
	513: "创建订单失败",
	514: "删除数据失败",
	515: "流失操作失败",
	516: "更新失败",
	517: "服务创建失败",
}

func NewError(errCode int) ErrorResponse {
	msg, ok := ErrCodeMsgMap[errCode]
	if !ok {
		msg = "Unknown Error"
	}
	return ErrorResponse{
		Code:    errCode,
		Message: msg,
	}
}
