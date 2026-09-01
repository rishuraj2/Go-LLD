package notification

import "notificationsystem/internal/domain/enum"

type Notification struct {
	id               string
	notificationType enum.NotificationType
	recipient        string
	subject          string
	content          string
	config           NotificationConfig
}

func (this Notification) GetID() string {
	return this.id
}

func (this Notification) GetType() enum.NotificationType {
	return this.notificationType
}

func (this Notification) GetRecipient() string {
	return this.recipient
}

func (this Notification) GetSubject() string {
	return this.subject
}

func (this Notification) GetContent() string {
	return this.content
}

func (this Notification) GetConfig() NotificationConfig {
	return this.config
}
