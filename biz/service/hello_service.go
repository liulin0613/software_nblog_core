package service

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"nblog.org.cn/software_nblog_core/biz/consts"
	"nblog.org.cn/software_nblog_core/biz/err_code"
	core "nblog.org.cn/software_nblog_core/biz/model/software_nblog_core"
	"nblog.org.cn/software_nblog_core/mylogger"
	"nblog.org.cn/software_nblog_core/utils"
)

func Greet(ctx context.Context, req *core.HelloReq) *core.HelloResp {
	mylogger.CtxInfof(ctx, "your name is : %v", req.GetName())
	if req.Name != "ll" {
		return &core.HelloResp{
			RespBody: "",
			Error:    utils.ConvErr(err_code.UnknownError),
			LogID:    thrift.StringPtr(fmt.Sprintf("%s", ctx.Value(consts.LOGIDKEY))),
		}
	}
	return &core.HelloResp{
		RespBody: "hello," + req.Name,
		Error:    utils.ConvErr(err_code.Success.WithCustomErrMsg("i am msg")),
		LogID:    thrift.StringPtr(fmt.Sprintf("%s", ctx.Value(consts.LOGIDKEY))),
	}
}
