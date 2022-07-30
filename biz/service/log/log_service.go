package log

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"nblog.org.cn/software_nblog_core/biz/consts"
	"nblog.org.cn/software_nblog_core/biz/err_code"
	core "nblog.org.cn/software_nblog_core/biz/model/software_nblog_core"
	"nblog.org.cn/software_nblog_core/mylogger"
	"nblog.org.cn/software_nblog_core/utils"
	"os"
	"strings"
)

func QueryLog(ctx context.Context, req *core.QueryLogReq) *core.QueryLogResp {
	mylogger.CtxInfof(ctx, "QueryLog para : %v", req.String())
	resp := &core.QueryLogResp{}
	resp.LogID = thrift.StringPtr(fmt.Sprintf("%s", ctx.Value(consts.LOGIDKEY)))
	err := paraCheck(req)
	if err != nil {
		mylogger.CtxInfof(ctx, "QueryLog err : %v", err)
		resp.Error = utils.ConvErr(err_code.ParaError)
		return resp
	}
	fileName := "./logs/" + req.Day + ".txt"
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		mylogger.CtxInfof(ctx, "file open fileName : %v , err : %v", fileName, err)
		resp.Error = utils.ConvErr(err_code.Success)
		resp.LogDetail = []string{"无记录"}
		return resp
	}
	defer file.Close()
	logStr := fmt.Sprintf(consts.LOGFORMAT, req.GetLogID())
	var logDetails []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, logStr) {
			logDetails = append(logDetails, strings.Replace(line, logStr, "", 1))
		}
	}
	if len(logDetails) == 0 {
		logDetails = append(logDetails, "无记录")
	}
	resp.LogDetail = logDetails
	resp.Error = utils.ConvErr(err_code.Success)
	return resp
}

func paraCheck(req *core.QueryLogReq) error {
	if req.GetDay() == "" || req.LogID == "" {
		return errors.New("day or log id is empty")
	}

	if !utils.MustMatch(req.GetDay(), consts.TimeRegex) {
		return errors.New("day format error")
	}

	return nil
}
