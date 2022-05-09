package response

import (
	"net/http"

	"github.com/bitrainforest/filmeta-hic/core/errno"
	"github.com/gin-gonic/gin/render"

	"github.com/gin-gonic/gin"
)

type (
	Response errno.Error

	GinContextFunc func(ctx *gin.Context) Response
)

func NewResponse(code int, errMsg string) Response {
	return errno.NewError(code, errMsg)
}

func Json(e GinContextFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e(ctx).RenderJson(ctx)
	}
}

func JsonCustom(fn func(ctx *gin.Context) interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, render.JSON{Data: fn(ctx)})
	}
}

func Text(e GinContextFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		e(ctx).RenderText(ctx)
	}
}
