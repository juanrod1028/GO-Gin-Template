package health

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHandler(logger *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info("Check Handler Init")
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
