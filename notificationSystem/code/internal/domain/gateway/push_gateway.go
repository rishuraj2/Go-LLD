package gateway

import "notificationsystem/internal/domain/notification"

type PushGateway struct{}

func NewPushGateway() PushGateway {
	return PushGateway{}
}

func (this PushGateway) Process(notification notification.Notification) error {
	return nil
}