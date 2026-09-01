package sender

import (
	"notificationsystem/internal/domain/gateway"
	"notificationsystem/internal/domain/notification"
)

type SMSSender struct {
	gateway *gateway.SMSGateway
}

func NewSMSSender(gateway *gateway.SMSGateway) *SMSSender {
	return &SMSSender{
		gateway: gateway,
	}
}

func (this *SMSSender) Send(notification notification.Notification) error {
	return this.gateway.Process(notification)
}
