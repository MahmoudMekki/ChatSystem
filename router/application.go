package router

import (
	"github.com/MahmoudMekki/ChatSystem/cmd/application"
	"github.com/MahmoudMekki/ChatSystem/validators"
)

func (r *routerImp) setAppRoutes() {
	appEndpoints := r.engine.Group("/api/v1/apps")
	appEndpoints.POST("", validators.ValidateCreateApplication(), application.CreateApp)
}
