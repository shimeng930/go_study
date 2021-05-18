package common_util

import (
	"context"
	_log "git.garena.com/shopee/loan-service/airpay_backend/public/common/log"
	"strings"
)

const maxStackLen = 2000

func CtxLog(ctx context.Context) *_log.OnceLog {
	return _log.ExtractOnceLog(ctx)
}

// StackLog remove newline, reserve maxStackLen's characters:w
func StackLog(stack []byte) string {
	strStack := string(stack)
	if len(strStack) > maxStackLen {
		strStack = strStack[:maxStackLen]
	}
	return strings.ReplaceAll(strStack, "\n", "->")
}
