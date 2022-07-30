package err_code

//所有错误码都在这里枚举

var (
	Success = buildError(0, "success")

	// 1xxxx 公共类错误
	InternalError  = buildError(1000, "系统异常，请稍后重试")
	UnknownError   = buildError(1001, "未知异常，请稍后重试")
	ParaError      = buildError(1002, "参数处理异常")
	UnDefinedError = buildError(1003, "未定义异常")
)
