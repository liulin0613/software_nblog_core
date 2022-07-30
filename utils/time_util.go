package utils

import (
	"nblog.org.cn/software_nblog_core/biz/consts"
	"time"
)

func GetToDay() string {
	return time.Now().Format(consts.TimeFormat)
}
