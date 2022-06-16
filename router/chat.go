package router

import (
	"github.com/MahmoudMekki/ChatSystem/cmd/chats"
	"github.com/MahmoudMekki/ChatSystem/validators"
)

func (r *routerImp) setChatRoutes() {
	appEndpoints := r.engine.Group("/api/v1/apps/:token/chats")
	appEndpoints.POST("", validators.ValidateCreateChat(), chats.CreateChat)
	appEndpoints.GET("/:number", validators.ValidateGetChat(), chats.GetChat)

}
