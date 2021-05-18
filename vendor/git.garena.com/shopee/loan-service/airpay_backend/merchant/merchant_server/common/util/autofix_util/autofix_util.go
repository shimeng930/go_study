package autofix_util

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type autofixFlag struct{}

const (
	metadataKeyMSAutofix = "ms-autofix"
)

func SetAutofixFlag(ctx context.Context) context.Context {
	// set in service flag
	ctx = context.WithValue(ctx, autofixFlag{}, true)

	// set cross service flag using grpc metadata
	ctx = metadata.AppendToOutgoingContext(ctx, metadataKeyMSAutofix, "1")
	return ctx
}

func IsAutofixContext(ctx context.Context) bool {
	// check in service flag
	if ctx.Value(autofixFlag{}) != nil {
		return true
	}

	// check cross service flag
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		vals := md.Get(metadataKeyMSAutofix)
		if len(vals) > 0 && vals[0] == "1" {
			return true
		}
	}
	return false
}
