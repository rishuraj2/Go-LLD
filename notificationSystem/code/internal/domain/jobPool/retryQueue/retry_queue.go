package retryqueue

import (
	"container/heap"
	"errors"
	"notificationsystem/internal/domain/notification"
	"sync"
)

type RetryQueue struct {
	queue PriorityQueue
	mu    sync.Mutex
}

var (
	retryQueueInstance *RetryQueue
	retryQueueOnce     sync.Once
)

var (
	errQueueFull = errors.New("job queue is full")
)

func NewRetryQueue() *RetryQueue {
	retryQueueOnce.Do(func() {
		retryQueueInstance = &RetryQueue{
			queue: NewPriorityQueue(),
		}
		heap.Init(&retryQueueInstance.queue)
	})

	return retryQueueInstance
}

func (this *RetryQueue) Enqueue(notification notification.Notification) error {
	select {
	case this.queue <- notification:
		return nil

	default:
		return errQueueFull
	}
}

func (this *RetryQueue) Listen() <-chan notification.Notification {
	return this.queue
}
