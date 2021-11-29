package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lmiedzinski/cutt-ovh-backend/docs"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/usecase"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/logger"
)

// @title       cutt.ovh Backend API
// @description cutt.ovh Backend
// @version     1.0
// @host        localhost:9000
// @BasePath    /v1
func NewRouter(handler *gin.Engine, logger logger.Interface, uc usecase.Redirect) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// Healthcheck
	handler.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	// Routers
	h := handler.Group("/v1")
	{
		newRedirectRoutes(h, uc, logger)
	}
}
