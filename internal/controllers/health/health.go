package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
}

func NewServer(router *gin.Engine) {
	server := &Server{}
	server.registerEndpoints(router)
}

func (s *Server) registerEndpoints(router *gin.Engine) {
	router.GET("/health", s.Health)
}

func (s *Server) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
