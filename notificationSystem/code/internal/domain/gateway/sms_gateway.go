package gateway

import "notificationsystem/internal/domain/notification"

type SMSGateway struct{}

func NewSMSGateway() SMSGateway {
	return SMSGateway{}
}

func (this SMSGateway) Process(notification notification.Notification) error {
	return nil
}