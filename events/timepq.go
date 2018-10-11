package events

import "github.com/gustavocovas/goscsim"

// timePriorityQueue is the private type that implements
// heap.Interface
type timePriorityQueue []*goscsim.Event

func (q timePriorityQueue) Len() int { return len(q) }

func (q timePriorityQueue) Less(i, j int) bool {
	return q[i].Time < q[j].Time
}

func (q timePriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *timePriorityQueue) Push(x interface{}) {
	item := x.(*goscsim.Event)
	*q = append(*q, item)
}

func (q *timePriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0 : n-1]
	return item
}
