package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Services  map[string]string `json:"services"`
}

// Health handles the health check endpoint
func (h *HealthHandler) Health(c echo.Context) error {
	health := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
		Services:  make(map[string]string),
	}

	// Check database connection
	sqlDB, err := h.db.DB()
	if err != nil {
		health.Status = "error"
		health.Services["database"] = "error: " + err.Error()
	} else if err := sqlDB.Ping(); err != nil {
		health.Status = "error"
		health.Services["database"] = "error: " + err.Error()
	} else {
		health.Services["database"] = "ok"
	}

	if health.Status == "error" {
		return c.JSON(http.StatusServiceUnavailable, health)
	}
	return c.JSON(http.StatusOK, health)
}

// Ready handles the readiness check endpoint
func (h *HealthHandler) Ready(c echo.Context) error {
	ready := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
		Services:  make(map[string]string),
	}

	// Check database connection
	sqlDB, err := h.db.DB()
	if err != nil {
		ready.Status = "error"
		ready.Services["database"] = "error: " + err.Error()
	} else {
		// Check if database is responsive
		if err := sqlDB.Ping(); err != nil {
			ready.Status = "error"
			ready.Services["database"] = "error: " + err.Error()
		} else {
			ready.Services["database"] = "ok"
		}

		// Check database connection stats
		stats := sqlDB.Stats()
		if stats.OpenConnections > stats.MaxOpenConnections*90/100 { // 90% of max connections
			ready.Status = "error"
			ready.Services["database_connections"] = "error: high connection count"
		} else {
			ready.Services["database_connections"] = "ok"
		}
	}

	if ready.Status == "error" {
		return c.JSON(http.StatusServiceUnavailable, ready)
	}
	return c.JSON(http.StatusOK, ready)
}
