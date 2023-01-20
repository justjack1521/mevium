package httperr

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InternalError(err error, slug string, ctx *gin.Context) {
	httpRespondWithError(err, slug, http.StatusInternalServerError, ctx)
}

func UnauthorisedError(err error, slug string, ctx *gin.Context) {
	httpRespondWithError(err, slug, http.StatusUnauthorized, ctx)
}

func BadRequest(err error, slug string, ctx *gin.Context) {
	httpRespondWithError(err, slug, http.StatusBadRequest, ctx)
}

type ErrorResponse struct {
	Slug string `json:"message"`
}

func httpRespondWithError(err error, slug string, status int, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(status, ErrorResponse{
		Slug: slug,
	})
}
