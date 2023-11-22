package context

import "github.com/gin-gonic/gin"

type GinContext struct {
	Ctx *gin.Context
}

func (g *GinContext) ShouldBindUri(obj interface{}) error {
	return g.Ctx.ShouldBindUri(obj)
}

func (g *GinContext) ShouldBindJSON(obj interface{}) error {
	return g.Ctx.ShouldBindJSON(obj)
}

func (g *GinContext) JSON(code int, obj interface{}) {
	g.Ctx.JSON(code, obj)
}

func (g *GinContext) Param(key string) string {
	return g.Ctx.Param(key)
}
