package logger

import (
	"go.uber.org/zap"
)

// (zap + Gin-контекст)
// Здесь важно: мы передаём логгер в хендлеры через замыкание или dependency injection.
// Для Дня 1 хватит простого варианта с замыканием.

func NewLogger(level string) (*zap.Logger, error) {
	var l *zap.Logger
	var err error

	switch level {
	case "debug":
		l, err = zap.NewDevelopment()
	default:
		l, err = zap.NewProduction()
	}
	if err != nil {
		return nil, err
	}

	return l, nil
}
