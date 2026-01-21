package handler

import (
	"context"
	"net/http"
	"tenantapp/helper"
	"tenantapp/logger"
	"tenantapp/models"
	"time"

	"github.com/gin-gonic/gin"
)

func AddTenant(ctx *gin.Context) {
	logger := logger.LoggingInit()
	logger.Info("Adding new tenant.")

	c, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	tenant_obj := models.Tenant{}

	if err := ctx.ShouldBindJSON(&tenant_obj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add tenant.Check your input.",
		})
		return
	}

	logger.Info("Adding new tenant:")
	tenant_var, err := helper.AddTenantHelper(c, tenant_obj)
	if err != nil {
		if err == context.DeadlineExceeded {
			logger.Error("Request deadline exceeded.")
			ctx.JSON(http.StatusRequestTimeout, gin.H{
				"error": "request timed out",
			})
			return
		}

		logger.Error("Internal error occured.")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("Sucessfully added the tenant.")
	data := map[string]string{"status": "successfully added" + tenant_var}
	ctx.AsciiJSON(http.StatusOK, data)
}
