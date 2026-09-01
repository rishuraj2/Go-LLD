package notification

import "notificationsystem/internal/domain/enum"

type NotificationBuilder struct {
	id               string
	notificationType enum.NotificationType
	recipient        string
	subject          string
	content          string
	config           NotificationConfig
}

func NewNotificationBuilder(id string, notificationType enum.NotificationType, recipient string, content string) *NotificationBuilder {
	return &NotificationBuilder{
		id:               id,
		notificationType: notificationType,
		recipient:        recipient,
		content:          content,
	}
}

func (this *NotificationBuilder) Subject(subject string) *NotificationBuilder {
	this.subject = subject
	return this
}

func (this *NotificationBuilder) Config(config NotificationConfig) *NotificationBuilder {
	this.config = config
	return this
}

func (this *NotificationBuilder) Build() Notification {
	return Notification{
		id:               this.id,
		notificationType: this.notificationType,
		recipient:        this.recipient,
		subject:          this.subject,
		content:          this.content,
		config:           this.config,
	}
}
