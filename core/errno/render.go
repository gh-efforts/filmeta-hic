package errno

import (
	"github.com/gin-gonic/gin"
)

type (
	Render interface {
		RenderJson(ctx *gin.Context)
		//RenderCustomJson just response Json
		RenderCustomJson(ctx *gin.Context)
		// RenderText just response for text/plain
		RenderText(ctx *gin.Context)
	}
)

func (e errno) RenderJson(ctx *gin.Context) {
	ctx.JSON(e.GetHttpStatusCode(), e.resetNowTime())
	ctx.Abort()
}

func (e errno) RenderCustomJson(ctx *gin.Context) {
	ctx.JSON(e.GetHttpStatusCode(), e.GetData())
	ctx.Abort()
}

func (e errno) RenderText(ctx *gin.Context) {
	ctx.Data(e.GetHttpStatusCode(), "text/plain", e.GetRawData())
	ctx.Abort()
}
