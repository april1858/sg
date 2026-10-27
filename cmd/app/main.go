package main

import (
	"fmt"
	"log"

	"github.com/april1858/sg/config"
	"github.com/april1858/sg/internal/handler"
	"github.com/april1858/sg/internal/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Точка входа: cmd/app/main.go (Gin + логгер + конфиг)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	l, err := logger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatalf("error creating logger: %v", err)
	}
	defer l.Sync()

	l.Info("application starting",
		zap.String("port", cfg.AppPort),
		zap.String("log_level", cfg.LogLevel),
	)

	r := gin.Default()
	api := r.Group("/api") // все ручки будут по префиксу /api/...
	handler.RegisterSubscriptionHandlers(api, l)

	addr := fmt.Sprintf(":%s", cfg.AppPort)

	l.Info("listening on", zap.String("addr", addr))

	if err := r.Run(addr); err != nil {
		l.Error("server error", zap.Error(err)) // корректный вызов l.Error(err)
	}
}
