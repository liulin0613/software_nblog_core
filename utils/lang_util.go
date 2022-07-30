package utils

import (
	"nblog.org.cn/software_nblog_core/biz/consts"
	"regexp"
)

func GetLogFormat(format string) string {
	return consts.LOGFORMAT + format
}

func MustMatch(matchString string, format string) bool {
	mustRegexp := regexp.MustCompile(format)
	return mustRegexp.MatchString(matchString)
}
