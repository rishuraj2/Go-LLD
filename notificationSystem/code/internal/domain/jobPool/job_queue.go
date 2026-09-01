package jobpool

import (
	"errors"
	"notificationsystem/internal/domain/notification"
	"sync"
)

type JobQueue struct {
	queue chan notification.Notification
}

var (
	jobQueueInstance *JobQueue
	jobQueueOnce     sync.Once
)

var (
	errQueueFull = errors.New("job queue is full")
)

func GetJobQueue() *JobQueue {
	jobQueueOnce.Do(func() {
		jobQueueInstance = &JobQueue{
			queue: make(chan notification.Notification, 10),
		}
	})

	return jobQueueInstance
}

func (this *JobQueue) Enqueue(notification notification.Notification) error {
	select {
	case this.queue <- notification:
		return nil

	default:
		return errQueueFull
	}
}

func (this *JobQueue) Listen() <-chan notification.Notification {
	return this.queue
}
