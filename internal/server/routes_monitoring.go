package server

import (
	"github.com/gin-contrib/expvar"
	"github.com/gin-gonic/gin"

	"github.com/photoprism/photoprism/internal/config"
)

// registerMonitoringRoutes configures debugging and monitoring endpoints.
func registerMonitoringRoutes(router *gin.Engine, conf *config.Config) {

	// Serves golang metrics.
	if conf.EnableExpvar() {
		router.GET("/debug/vars", expvar.Handler())
	}
}
