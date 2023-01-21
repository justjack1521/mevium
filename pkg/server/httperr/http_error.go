package httperr

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/errors"
	"github.com/justjack1521/mevium/pkg/genproto/common"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Slug  string `json:"message"`
}

func ResponseError(header *common.ResponseHeader, ctx *gin.Context) {
	httpRespondWithError(int(header.StatusCode), errors.NewSlugError(header.Error, header.Message), ctx)
}

func InternalError(err error, slug string, ctx *gin.Context) {
	httpRespondWithError(http.StatusInternalServerError, errors.NewSlugError(err.Error(), slug), ctx)
}

func UnauthorisedError(err error, slug string, ctx *gin.Context) {
	httpRespondWithError(http.StatusUnauthorized, errors.NewSlugError(err.Error(), slug), ctx)
}

func BadRequest(err error, slug string, ctx *gin.Context) {
	httpRespondWithError(http.StatusBadRequest, errors.NewSlugError(err.Error(), slug), ctx)
}

func httpRespondWithError(status int, err error, ctx *gin.Context) {
	ctx.AbortWithError(status, err)
}
