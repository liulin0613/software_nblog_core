package utils

import (
	"nblog.org.cn/software_nblog_core/biz/err_code"
	core "nblog.org.cn/software_nblog_core/biz/model/software_nblog_core"
)

func ConvErr(e err_code.NBlogError) *core.Error {
	return &core.Error{
		ErrNo:   e.ErrorNo(),
		ErrTips: e.ErrorMsg(),
	}
}
