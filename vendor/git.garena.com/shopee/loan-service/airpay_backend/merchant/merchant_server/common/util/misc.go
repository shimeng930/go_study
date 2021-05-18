package common_util

import (
	"context"

	"git.garena.com/shopee/loan-service/airpay_backend/merchant/merchant_server/common/util/autofix_util"
	"git.garena.com/shopee/loan-service/airpay_backend/public/common/metadata"
	"github.com/opentracing/opentracing-go"
	grpc_meta "google.golang.org/grpc/metadata"
)

const (
	uid      = "uid"
	remoteIp = "remote_ip"
)

func NewTracingContext(ctx context.Context) context.Context {
	newCtx := context.Background()
	incomingMD := grpc_meta.MD{}

	uidFromCtx := metadata.GetUidFromCtx(ctx)
	if uidFromCtx != "" {
		newCtx = metadata.SetUidIntoContext(newCtx, uidFromCtx)
		incomingMD.Set(uid, uidFromCtx)
	}

	remoteIpFromCtx := metadata.GetRemoteIP(ctx)
	if len(remoteIpFromCtx) > 0 {
		newCtx = metadata.SetRemoteipIntoCtx(newCtx, remoteIpFromCtx)
		incomingMD.Set(remoteIp, remoteIpFromCtx)
	}

	newCtx = grpc_meta.NewIncomingContext(newCtx, incomingMD)

	if autofix_util.IsAutofixContext(ctx) {
		newCtx = autofix_util.SetAutofixFlag(newCtx)
	}

	spanContext := opentracing.SpanFromContext(ctx)
	if spanContext == nil {
		return newCtx
	}
	return opentracing.ContextWithSpan(newCtx, spanContext)
}
