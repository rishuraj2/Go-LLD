package retryqueue

import "notificationsystem/internal/domain/notification"

type PriorityQueue struct {
	queue []notification.Notification
}

func NewPriorityQueue() PriorityQueue {
	return PriorityQueue{
		queue: make([]notification.Notification, 0),
	}
}

func (this *PriorityQueue) Len() int {
	return len(this.queue)
}

// learn about this
func (this *PriorityQueue) Less(i, j int) bool {
	return this.queue[i].GetConfig().GetNextTryAt().Before(this.queue[j].GetConfig().GetNextTryAt())
}

func (this *PriorityQueue) Swap(i, j int) {
	this.queue[i], this.queue[j] = this.queue[j], this.queue[i]
}

func (this *PriorityQueue) Push(x any) {
	n := x.(notification.Notification)
	this.queue = append(this.queue, n)
}

func (this *PriorityQueue) Pop() any {
	old := this.queue
	n := len(old)
	item := old[n-1]
	this.queue = old[0 : n-1]
	return item
}
