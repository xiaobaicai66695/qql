/**
 * @author: dn-jinmin/dn-jinmin
 * @doc:
 */

package handler

import (
	"qql/apps/im/ws/internal/handler/conversation"
	"qql/apps/im/ws/internal/handler/push"
	"qql/apps/im/ws/internal/handler/user"
	"qql/apps/im/ws/internal/svc"
	"qql/apps/im/ws/websocket"
)

func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoutes([]websocket.Route{
		{
			Method:  "user.online",
			Handler: user.OnLine(svc),
		},
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "conversation.markChat",
			Handler: conversation.MarkRead(svc),
		},
		{
			Method:  "push",
			Handler: push.Push(svc),
		},
	})
}
