package middleware

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"nblog.org.cn/software_nblog_core/biz/err_code"
	core "nblog.org.cn/software_nblog_core/biz/model/software_nblog_core"
	"nblog.org.cn/software_nblog_core/mylogger"
)

type CommonErr struct {
	Error core.Error
}

func HertzErrCodeMw() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)
		cr := &CommonErr{}
		err := json.Unmarshal(c.Response.BodyBytes(), cr)
		if err != nil {
			mylogger.CtxErrorf(ctx, "check error code meet exception : unmarshal err : %v", err)
		}
		curErr := err_code.GetFromErrorNo(cr.Error.ErrNo)
		// 未知 errNo
		if curErr.Is(err_code.UnDefinedError) {
			mylogger.CtxErrorf(ctx, "http response body contain undefined error code, err : %v", cr.Error.String())
		} else {
			// 未知 errTips
			if cr.Error.ErrTips != curErr.ErrorMsg() {
				mylogger.CtxErrorf(ctx, "http response body err msg is undefined, err : %v", cr.Error.String())
			}
		}
	}
}
