package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/usecase"
	"github.com/lmiedzinski/cutt-ovh-backend/pkg/logger"
)

type redirectRoutes struct {
	ru     usecase.Redirect
	logger logger.Interface
}

func newRedirectRoutes(handler *gin.RouterGroup, ru usecase.Redirect, logger logger.Interface) {
	routes := &redirectRoutes{ru, logger}

	h := handler.Group("/redirect")
	{
		h.POST("/", routes.postRedirect)
	}
}

type postRedirectRequest struct {
	Url string `json:"url"`
}

type postRedirectResponse struct {
	Url  string `json:"url"`
	Slug string `json:"slug"`
}

// @Summary     Add redirect
// @Description Adds new short url redirect
// @ID          postRedirect
// @Tags  	    redirect
// @Accept      json
// @Produce     json
// @Param       request body postRedirectRequest true "Set up url to shorten"
// @Success     200 {object} postRedirectResponse
// @Failure     500 {object} response
// @Router      /redirect [post]
func (r *redirectRoutes) postRedirect(c *gin.Context) {
	var request postRedirectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.logger.Error(err)
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	redirect, err := r.ru.CreateRedirect(
		c.Request.Context(),
		request.Url,
	)
	if err != nil {
		r.logger.Error(err)
		errorResponse(c, http.StatusInternalServerError, "cannot create short url")
		return
	}

	c.JSON(http.StatusOK, redirect)
}
