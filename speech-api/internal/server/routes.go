package server

import (
	"net/http"
	"speech-api/internal/handler"
	"speech-api/internal/pkg/httpclient"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes(httpClient httpclient.HttpClient, stream handler.Stream) http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/health", handler.HealthCheck)
	r.GET("/stream", stream.Listen)

	return r
}
