package helper

import (
	"context"
	"time"
	"tokenapp/jwtutil"
	"tokenapp/logger"
)

func TokenGenHelper(ctx context.Context, username, role string) (string, error) {
	logger := logger.LoggingInit()
	logger.Info("Generating jwt token.")
	select {
	case <-time.After(1 * time.Nanosecond):
		token, err := jwtutil.GenerateToken(username, role)
		if err != nil {
			return "Failed to generate token.", err
		}
		return token, nil
	case <-ctx.Done():
		logger.Info("Request deadline exceeded.")
		return "", ctx.Err()
	}
}
