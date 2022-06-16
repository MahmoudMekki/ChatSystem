package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	SetRouter() *gin.Engine
}

func NewRouter(router *gin.Engine) Router {

	return &routerImp{
		engine: router,
	}
}

type routerImp struct {
	engine *gin.Engine
}

func (r *routerImp) SetRouter() *gin.Engine {
	r.setAppRoutes()
	return r.engine
}
