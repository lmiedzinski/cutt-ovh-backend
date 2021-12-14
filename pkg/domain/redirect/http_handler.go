package redirect

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lmiedzinski/cutt-ovh-backend/internal/logger"
)

type RedirectHttpHandler struct {
	service *RedirectService
	logger  logger.Interface
}

func NewRedirectHttpHandler(s *RedirectService, l logger.Interface) *RedirectHttpHandler {
	return &RedirectHttpHandler{service: s, logger: l}
}

func (h *RedirectHttpHandler) AddRedirectRoutes(rg *gin.RouterGroup) {

	hg := rg.Group("/redirect")
	{
		hg.POST("/", h.postRedirect)
		hg.GET("/:slug/execute", h.getRedirectExecute)
		hg.GET("/:slug", h.getRedirectBySlug)
	}
}

// @Summary     Add redirect
// @Description Adds new short url redirect
// @ID          postRedirect
// @Tags  	    redirect
// @Accept      json
// @Produce     json
// @Param       request body postRedirectRequest true "Set up url to shorten"
// @Success     200 {object} postRedirectResponse
// @Failure     500 {object} redirectErrorResponse
// @Router      /redirect [post]
func (h *RedirectHttpHandler) postRedirect(c *gin.Context) {
	var request postRedirectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		h.logger.Error(err)
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	redirect, err := h.service.createRedirect(
		c.Request.Context(),
		request.Url,
	)
	if err != nil {
		h.logger.Error(err)
		errorResponse(c, http.StatusInternalServerError, "cannot create short url")
		return
	}

	c.JSON(http.StatusOK, postRedirectResponse{Url: redirect.Url, Slug: redirect.Slug})
}

// @Summary     Get redirect
// @Description Gets shortened url info by slug
// @ID          getRedirectBySlug
// @Tags  	    redirect
// @Accept      json
// @Produce     json
// @Param       slug  path string true "Slug"
// @Success     200 {object} getRedirectBySlugResponse
// @Failure     500 {object} redirectErrorResponse
// @Failure     404 {object} redirectErrorResponse
// @Router      /redirect/{slug} [get]
func (h *RedirectHttpHandler) getRedirectBySlug(c *gin.Context) {
	slug := c.Param("slug")
	r, err := h.service.getRedirect(c.Request.Context(), slug)
	if err != nil {
		h.logger.Error(err)
		errorResponse(c, http.StatusInternalServerError, "cannot get shortened url info")
		return
	}
	if r == (redirect{}) {
		errorResponse(c, http.StatusNotFound, "shortened url info not found for slug")
		return
	}
	c.JSON(http.StatusOK, getRedirectBySlugResponse{Url: r.Url, Slug: r.Slug})
}

// @Summary     Get execute redirect
// @Description Executes redirect to original url
// @ID          getRedirectExecute
// @Tags  	    redirect
// @Accept      json
// @Param       slug  path string true "Slug"
// @Success     301
// @Failure     500 {object} redirectErrorResponse
// @Failure     404 {object} redirectErrorResponse
// @Router      /redirect/{slug}/execute [get]
func (h *RedirectHttpHandler) getRedirectExecute(c *gin.Context) {
	slug := c.Param("slug")
	r, err := h.service.getRedirect(c.Request.Context(), slug)
	if err != nil {
		h.logger.Error(err)
		errorResponse(c, http.StatusInternalServerError, "cannot get shortened url info")
		return
	}
	if r == (redirect{}) {
		errorResponse(c, http.StatusNotFound, "shortened url info not found for slug")
		return
	}
	c.Redirect(http.StatusMovedPermanently, r.Url)
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, redirectErrorResponse{Code: code, Message: msg})
}
