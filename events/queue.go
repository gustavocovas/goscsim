package events

import (
	"container/heap"

	"github.com/gustavocovas/goscsim"
)

// Queue is a queue for goscsim events
type Queue struct {
	pq timePriorityQueue
}

// New creates a new Queue
func New() *Queue {
	return &Queue{
		pq: make(timePriorityQueue, 0),
	}
}

// Push enqueues e
func (q *Queue) Push(e *goscsim.Event) {
	heap.Push(&q.pq, e)
}

// Pop returns the next event, based on the Time field
func (q *Queue) Pop() *goscsim.Event {
	return heap.Pop(&q.pq).(*goscsim.Event)
}

// Len returns the current length of the queue
func (q *Queue) Len() int {
	return q.pq.Len()
}
