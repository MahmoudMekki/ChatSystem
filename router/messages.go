package router

import (
	"github.com/MahmoudMekki/ChatSystem/cmd/messages"
	"github.com/MahmoudMekki/ChatSystem/validators"
)

func (r *routerImp) setMessagesRoutes() {
	appEndpoints := r.engine.Group("/api/v1/apps/:token/chats/:chat_number/messages")
	appEndpoints.POST("", validators.ValidateCreateMessage(), messages.CreateMsg)
	appEndpoints.GET("/:msg_number", validators.ValidateGetMessage(), messages.GetMsg)
}
