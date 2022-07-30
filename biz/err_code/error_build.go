package err_code

import "fmt"

var errorMap map[int64]NBlogError

func buildError(errorNo int64, errorMsg string) NBlogError {
	if errorMap == nil {
		errorMap = make(map[int64]NBlogError)
	}
	if _, ok := errorMap[errorNo]; ok {
		panic(fmt.Sprintf("duplicated error code definition errorNo = %v", errorNo))
	}
	err := NBlogError{
		errorNo:  errorNo,
		errorMsg: errorMsg,
	}
	errorMap[errorNo] = err

	return err
}
