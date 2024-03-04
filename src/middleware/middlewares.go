package middlewares

import (
	"time"

	ginZap "github.com/gin-contrib/zap"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterMiddlewares(router *gin.Engine) {
	logger, _ := zap.NewProduction()
	router.Use(LoggerMiddleware())
	router.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	router.Use(ginZap.RecoveryWithZap(logger, true))
}
