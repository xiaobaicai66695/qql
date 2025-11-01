package msgTransfer

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"qql/apps/task/handler/svc"
)

type MsgChatTransfer struct {
	logx.Logger
	svc *svc.ServiceContext
}

func NewMsgChatTransfer(svc *svc.ServiceContext) *MsgChatTransfer {
	return &MsgChatTransfer{
		Logger: logx.WithContext(context.Background()),
		svc:    svc,
	}
}

func (m *MsgChatTransfer) Consume(context context.Context, key, value string) error {
	fmt.Println(key, value)
	return nil
}
