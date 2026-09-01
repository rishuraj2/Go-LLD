package jobpool

import retryqueue "notificationsystem/internal/domain/jobPool/retryQueue"

type JobPool struct {
	jobQueue   *JobQueue
	retryQueue *retryqueue.RetryQueue
}
