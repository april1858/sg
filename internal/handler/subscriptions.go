package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// (заглушки на Gin)
// Пока без БД, просто возвращаем 501 и логируем.
// Это соответствует требованиям «HTTP‑ручки» и «покрыть код логами».

func RegisterSubscriptionHandlers(rg *gin.RouterGroup, l *zap.Logger) {
	// POST /subscriptions — создание
	rg.POST("/subscriptions", func(c *gin.Context) {
		l.Info("POST /subscriptions called (stub)",
			zap.Any("payload", c.Request.Body),
		)
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
	})

	// GET /subscriptions — список (с фильтрами user_id, service_name)
	rg.GET("/subscriptions", func(c *gin.Context) {
		userID := c.Query("user_id")
		serviceName := c.Query("service_name")

		l.Info("GET /subscriptions called (stub)",
			zap.String("user_id", userID),
			zap.String("service_name", serviceName),
		)
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
	})

	// GET /subscriptions/:id — одна запись
	rg.GET("/subscriptions/:id", func(c *gin.Context) {
		id := c.Param("id")
		l.Info("GET /subscriptions/{id} called (stub)", zap.String("id", id))
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
	})

	// PUT /subscriptions/:id — обновление
	rg.PUT("/subscriptions/:id", func(c *gin.Context) {
		id := c.Param("id")
		l.Info("PUT /subscriptions/{id} called (stub)", zap.String("id", id))
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
	})

	// DELETE /subscriptions/:id — удаление
	rg.DELETE("/subscriptions/:id", func(c *gin.Context) {
		id := c.Param("id")
		l.Info("DELETE /subscriptions/{id} called (stub)", zap.String("id", id))
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet"})
	})

	// GET /subscriptions/total — агрегация по периоду и фильтрам
	rg.GET("/subscriptions/total", func(c *gin.Context) {
		userID := c.Query("user_id")
		serviceName := c.Query("service_name")
		startPeriod := c.Query("start_period") // MM-YYYY
		endPeriod := c.Query("end_period")     // MM-YYYY

		l.Info("GET /subscriptions/total called (stub)",
			zap.String("user_id", userID),
			zap.String("service_name", serviceName),
			zap.String("start_period", startPeriod),
			zap.String("end_period", endPeriod),
		)
		c.JSON(http.StatusNotImplemented, gin.H{"message": "not implemented yet", "total_cost": 0})
	})
}
