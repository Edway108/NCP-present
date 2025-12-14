package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)



func setupLogger() *zap.Logger {
	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    1, // MB
		MaxBackups: 3, // số file log cũ
		MaxAge:     7, // ngày
		Compress:   true,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writer,
		zapcore.InfoLevel,
	)

	return zap.New(core)
}

func main() {
	logger := setupLogger()
	defer logger.Sync()

	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var body struct {
			Username string `json:"username"`
		}

		if err := c.ShouldBindJSON(&body); err != nil {
			logger.Error("Invalid request body",
				zap.Error(err),
			)

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid body",
			})
			return
		}

		logger.Info("User login",
			zap.String("username", body.Username),
			zap.String("ip", c.ClientIP()),
		)

		c.JSON(http.StatusOK, gin.H{
			"message": "login success",
		})
	})

	logger.Info("Server started", zap.String("port", "8080"))
	r.Run(":8080")
}
