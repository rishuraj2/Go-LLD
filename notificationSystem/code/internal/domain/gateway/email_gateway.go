package gateway

import "notificationsystem/internal/domain/notification"

type EmailGateway struct{}

func NewEmailGateway() EmailGateway {
	return EmailGateway{}
}

func (this EmailGateway) Process(notification notification.Notification) error {
	return nil
}
