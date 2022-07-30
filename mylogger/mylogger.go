package mylogger

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"nblog.org.cn/software_nblog_core/biz/consts"
	"nblog.org.cn/software_nblog_core/utils"
)

func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	logID := ctx.Value(consts.LOGIDKEY)
	hlog.CtxErrorf(ctx, utils.GetLogFormat(format), logID, v)
}

func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	logID := ctx.Value(consts.LOGIDKEY)
	hlog.CtxInfof(ctx, utils.GetLogFormat(format), logID, v)
}

func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	logID := ctx.Value(consts.LOGIDKEY)
	hlog.CtxDebugf(ctx, utils.GetLogFormat(format), logID, v)
}

func CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	logID := ctx.Value(consts.LOGIDKEY)
	hlog.CtxFatalf(ctx, utils.GetLogFormat(format), logID, v)
}
