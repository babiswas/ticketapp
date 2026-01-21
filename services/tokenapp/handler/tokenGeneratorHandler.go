package handler

import (
	"context"
	"net/http"
	"time"
	"tokenapp/helper"
	"tokenapp/logger"

	"github.com/gin-gonic/gin"
)

func TokenGenerator(ctx *gin.Context) {
	logger := logger.LoggingInit()
	logger.Info("Executing token generator handler.")

	c, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	var user_info_obj struct {
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	if err := ctx.ShouldBindJSON(&user_info_obj); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to generate token.Check payload.",
		})
		return
	}

	logger.Info("Generating token using token handler:")
	token, err := helper.TokenGenHelper(c, user_info_obj.Username, user_info_obj.Role)
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
	logger.Info("Sucessfully generated jwt token.")
	data := map[string]string{"token": token}
	ctx.AsciiJSON(http.StatusOK, data)
}
