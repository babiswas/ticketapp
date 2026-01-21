package helper

import (
	"context"
	"time"
	"tokenapp/jwtutil"
	"tokenapp/logger"
)

func TokenValidationHelper(ctx context.Context, token string) (bool, error) {
	logger := logger.LoggingInit()
	logger.Info("Validating jwt token.")
	select {
	case <-time.After(1 * time.Nanosecond):
		status := jwtutil.IsAuthenticated(token)
		return status, nil
	case <-ctx.Done():
		logger.Info("Request deadline exceeded.")
		return false, ctx.Err()
	}
}
