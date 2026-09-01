package sender

import (
	"notificationsystem/internal/domain/gateway"
	"notificationsystem/internal/domain/notification"
)

type PushSender struct {
	gateway *gateway.PushGateway
}

func NewPushSender(gateway *gateway.PushGateway) *PushSender {
	return &PushSender{
		gateway: gateway,
	}
}

func (this *PushSender) Send(notification notification.Notification) error {
	return this.gateway.Process(notification)
}
