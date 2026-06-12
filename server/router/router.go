package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zu-bae/v-connect/server/config"
	"github.com/zu-bae/v-connect/server/handler"
)

func New(cfg *config.Config) *gin.Engine {
	if !cfg.IsDev() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	{
		api := r.Group("/api")
		{
			api.GET("/health", handler.Health)
		}
	}

	r.NoRoute(spaHandler(cfg.ClientDir))

	return r
}

func spaHandler(clientDir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Guard: unknown /api paths get JSON 404s, never index.html.
		if strings.HasPrefix(c.Request.URL.Path, "/api/") {
			c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
			return
		}

		// Clean("/"+path) pins the path to the clientDir (no ../ escapes)
		reqPath := filepath.Clean("/" + c.Request.URL.Path)
		filePath := filepath.Join(clientDir, reqPath)

		if info, err := os.Stat(filePath); err == nil && !info.IsDir() {
			c.File(filePath)
		}

		// SPA Fallback
		index := filepath.Join(clientDir, "index.html")
		if _, err := os.Stat(index); err != nil {
			// Build folder missing, fail loudly instead of silently returning 404s.
			c.JSON(http.StatusInternalServerError, gin.H{"error": "index.html not found"})
			return
		}
		c.File(index)
	}
}
