package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lmiedzinski/cutt-ovh-backend/docs"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/logger"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/domain/redirect"
)

// @title       cutt.ovh Backend API
// @description cutt.ovh Backend
// @version     1.0
// @host        localhost:9000
// @BasePath    /v1
func AddHttpRouter(handler *gin.Engine, logger logger.Interface, h *redirect.RedirectHttpHandler) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Healthcheck
	handler.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	// Routers
	hg := handler.Group("/v1")
	{
		h.AddRedirectRoutes(hg)
	}
}
