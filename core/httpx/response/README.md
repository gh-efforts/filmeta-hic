## 返回

### 实例代码
```go
func (h *Handler) Example(c *gin.Context) httpx.Response {
	var req Req
	if err := c.ShouldBind(&req); err != nil {
		return errno.ParamValidationErr.WithReason(validate.TransErr(err))
	}

	return errno.OK.WithData(data)
}
```
```
router.POST("/json", response.Json(h.Example))
router.GET("/text", response.Text(func(ctx *gin.Context) response.Response {
    return errno.OK.WithData("text")
}))
```