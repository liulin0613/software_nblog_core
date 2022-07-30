package mylogger

import (
	"context"
	"testing"
)

func TestCtxInfof(t *testing.T) {
	CtxInfof(context.Background(), "%v", "css")
}
