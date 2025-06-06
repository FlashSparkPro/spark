package logging

import (
	"context"
	"encoding/hex"
	"log/slog"
)

type loggerContextKey string

const loggerKey = loggerContextKey("slog")

type Attr struct {
	Key   string
	Value any
}

// Inject the logger into the context. This should ONLY be called from the start of a request
// or worker context (e.g. in a top-level gRPC interceptor).
func Inject(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func WithIdentityPubkey(ctx context.Context, pubkey []byte) (context.Context, *slog.Logger) {
	return WithAttr(ctx, Attr{Key: "identity_public_key", Value: Pubkey{Pubkey: pubkey}})
}

func WithAttr(ctx context.Context, attr Attr) (context.Context, *slog.Logger) {
	logger := GetLoggerFromContext(ctx).With(slog.Any(attr.Key, attr.Value))
	return Inject(ctx, logger), logger
}

func WithAttrs(ctx context.Context, attrs []Attr) (context.Context, *slog.Logger) {
	logger := GetLoggerFromContext(ctx).With(toAnyArray(attrs)...)
	return Inject(ctx, logger), logger
}

func toAnyArray[T any](array []T) []any {
	anyArray := make([]any, len(array))
	for i, attr := range array {
		anyArray[i] = attr
	}
	return anyArray
}

// Get an instance of slog.Logger from the current context. If no logger is found, returns a
// default logger.
func GetLoggerFromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}

type Pubkey struct{ Pubkey []byte }

func (l Pubkey) LogValue() slog.Value {
	return slog.AnyValue(hex.EncodeToString(l.Pubkey))
}
