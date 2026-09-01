package notification

import "time"

type NotificationConfig struct {
	retries       int
	retryInterval int
	nextTryAt     time.Time
}

func NewNotificationConfig(retries, retryInterval int) NotificationConfig {
	return NotificationConfig{
		retries:       retries,
		retryInterval: retryInterval,
	}
}

func (this NotificationConfig) GetRetries() int {
	return this.retries
}

func (this NotificationConfig) GetRetryInterval() int {
	return this.retryInterval
}

func (this NotificationConfig) GetNextTryAt() time.Time {
	return this.nextTryAt
}

func (this *NotificationConfig) SetNextTryAt(at time.Time) {
	this.nextTryAt = at
}
