package queue

import (
	"container/heap"
)

type item struct {
	key       interface{}
	value     interface{}
	timestamp int64
	index     int
}

type temporalQueue []*item

func (pq temporalQueue) Len() int { return len(pq) }

func (pq temporalQueue) Less(i, j int) bool {
	return pq[i].timestamp < pq[j].timestamp
}

func (pq temporalQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *temporalQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *temporalQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *temporalQueue) peek() interface{} {
	return (*pq)[0]
}

func (pq *temporalQueue) update(item *item, timestamp int64) {
	item.timestamp = timestamp
	heap.Fix(pq, item.index)
}
