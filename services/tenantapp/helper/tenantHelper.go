package helper

import (
	"context"
	"fmt"
	"tenantapp/logger"
	"tenantapp/models"
	"time"
)

func AddTenantHelper(ctx context.Context, tenant_obj models.Tenant) (string, error) {
	logger := logger.LoggingInit()
	logger.Info("Validating jwt token.")
	select {
	case <-time.After(1 * time.Nanosecond):
		fmt.Println("Adding tenant.", tenant_obj)
		return "NewTenant", nil
	case <-ctx.Done():
		logger.Info("Request deadline exceeded.")
		return "Failure", ctx.Err()
	}
}
