package handler

import (
	"context"
	"net/http"
	"time"
	"tokenapp/helper"
	"tokenapp/logger"

	"github.com/gin-gonic/gin"
)

func ValidateToken(ctx *gin.Context) {
	logger := logger.LoggingInit()
	logger.Info("Executing token validator handler.")

	c, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	var token_obj struct {
		token string
	}

	if err := ctx.ShouldBindJSON(&token_obj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to validate token.Check payload.",
		})
		return
	}

	logger.Info("Generating token using token handler:")
	status, err := helper.TokenValidationHelper(c, token_obj.token)
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
	logger.Info("Sucessfully validated the token STATUS:", status)
	data := map[string]bool{"token_status": status}
	ctx.AsciiJSON(http.StatusOK, data)
}
