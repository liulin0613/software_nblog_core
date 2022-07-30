package err_code

import "fmt"

type NBlogError struct {
	errorNo  int64
	errorMsg string
}

// Is 判断是否为某个具体错误
func (e NBlogError) Is(err NBlogError) bool {
	return err.errorNo == e.errorNo
}

// Error 错误信息
func (e NBlogError) Error() string {
	return fmt.Sprintf("err_no = %d err_msg = %s", e.errorNo, e.errorMsg)
}

// ErrorNo 错误码
func (e NBlogError) ErrorNo() int64 {
	return e.errorNo
}

// ErrorMsg 错误描述
func (e NBlogError) ErrorMsg() string {
	return e.errorMsg
}

// WithCustomErrMsg 自定义某个错误的错误描述
func (e NBlogError) WithCustomErrMsg(errMsg string) NBlogError {
	return NBlogError{
		errorNo:  e.ErrorNo(),
		errorMsg: errMsg,
	}
}

// GetFromErrorNo 通过错误码获取错误
func GetFromErrorNo(errorNO int64) NBlogError {
	if err, ok := errorMap[errorNO]; ok {
		return err
	}
	return UnDefinedError
}
