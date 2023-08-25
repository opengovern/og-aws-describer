package describer

import (
	"context"
	"go.uber.org/zap"

	"github.com/kaytu-io/kaytu-util/pkg/describe/enums"
)

var (
	key            describeContextKey = "describe_ctx"
	triggerTypeKey string             = "trigger_type"
)

type describeContextKey string

type DescribeContext struct {
	AccountID   string
	Region      string
	KaytuRegion string
	Partition   string
}

func WithDescribeContext(ctx context.Context, describeCtx DescribeContext) context.Context {
	return context.WithValue(ctx, key, describeCtx)
}

func GetDescribeContext(ctx context.Context) DescribeContext {
	describe, ok := ctx.Value(key).(DescribeContext)
	if !ok {
		panic("context key not found")
	}
	return describe
}

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value("logger").(*zap.Logger)
	if !ok {
		return zap.NewNop()
	}
	return logger
}

func WithTriggerType(ctx context.Context, tt enums.DescribeTriggerType) context.Context {
	return context.WithValue(ctx, triggerTypeKey, tt)
}

func GetTriggerTypeFromContext(ctx context.Context) enums.DescribeTriggerType {
	tt, ok := ctx.Value(triggerTypeKey).(enums.DescribeTriggerType)
	if !ok {
		return ""
	}
	return tt
}
