package main

import (
	"net/http"
	"os"
	"time"
	"tokenapp/handler"
	"tokenapp/helper"
	"tokenapp/logger"
	Middleware "tokenapp/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	helper.LoadEnv()
}

func main() {

	logger := logger.LoggingInit()

	logger.Info("Initiallised logger.")
	logger.Info("Fetching port from enviroment variable.")
	port := ":" + os.Getenv("PORT")

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	logger.Info("Configuring the middleware.")
	router.Use(Middleware.HostValidationMiddleWare(logger), Middleware.LoggerMiddleWare(logger), gin.Recovery())

	logger.Info("Configuring the routes:")
	router.POST("/fetchToken", handler.TokenGenerator)

	logger.Info("Configuraing the server.")
	server := &http.Server{
		Addr:         port,
		Handler:      router,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	server.ListenAndServe()
}
