package sender

import (
	"notificationsystem/internal/domain/gateway"
	"notificationsystem/internal/domain/notification"
)

type EmailSender struct {
	gateway *gateway.EmailGateway
}

func NewEmailSender(gateway *gateway.EmailGateway) *EmailSender {
	return &EmailSender{
		gateway: gateway,
	}
}

func (this *EmailSender) Send(notification notification.Notification) error {
	return this.gateway.Process(notification)
}
