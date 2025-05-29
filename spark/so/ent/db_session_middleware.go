package ent

import (
	"context"
	"errors"

	"github.com/lightsparkdev/spark/common/logging"
	"google.golang.org/grpc"
)

// ErrNoRollback is an error indicating that we should not rollback the DB transaction.
var ErrNoRollback = errors.New("no rollback performed")

// DbSessionMiddleware is a middleware to manage database sessions for each gRPC call.
func DbSessionMiddleware(dbClient *Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if info != nil && info.FullMethod == "/grpc.health.v1.Health/Check" {
			return handler(ctx, req)
		}

		// Start a transaction or session
		tx, err := dbClient.Tx(ctx)
		if err != nil {
			return nil, err
		}

		// Attach the transaction to the context
		ctx = Inject(ctx, tx)
		// Ensure rollback on panic
		defer func() {
			if r := recover(); r != nil {
				_ = tx.Rollback()
				panic(r)
			}
		}()

		logger := logging.GetLoggerFromContext(ctx)

		// Call the handler (the actual RPC method)
		resp, err := handler(ctx, req)

		// Handle transaction commit/rollback
		if err != nil && !errors.Is(err, ErrNoRollback) {
			if dberr := tx.Rollback(); dberr != nil {
				logger.Error("Failed to rollback transaction", "error", dberr)
			}
			return nil, err
		}

		if dberr := tx.Commit(); dberr != nil {
			logger.Error("Failed to commit transaction", "error", dberr)
			return nil, dberr
		}

		if errors.Is(err, ErrNoRollback) {
			logger.Debug("Skipping rollback", "error", err)
			return nil, err
		}

		return resp, nil
	}
}
