package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"nblog.org.cn/software_nblog_core/biz/consts"
	"nblog.org.cn/software_nblog_core/mylogger"
	"nblog.org.cn/software_nblog_core/utils"
)

func HertzLogMw() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		logID := utils.GetUniqLogID()
		ctx = context.WithValue(ctx, consts.LOGIDKEY, logID)
		mylogger.CtxInfof(ctx, "method: %v", string(c.Request.Path()))
		c.Next(ctx)
	}
}
