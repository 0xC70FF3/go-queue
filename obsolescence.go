package queue

import (
	"container/heap"
)

type obsolescenceQueue struct {
	priorityQueue *temporalQueue
	index         map[interface{}]*item
}

func ObsolescenceQueue() obsolescenceQueue {
	pq := make(temporalQueue, 0)
	idx := make(map[interface{}]*item)
	ipq := obsolescenceQueue{priorityQueue: &pq, index: idx}
	heap.Init(ipq.priorityQueue)
	return ipq
}

func (ipq obsolescenceQueue) Push (key interface{}, value interface{}, priority int64){
	item := &item{
		value:    value,
		key:      key,
		timestamp: priority,
	}
	heap.Push(ipq.priorityQueue, item)
	ipq.index[key] = item
}

func (ipq obsolescenceQueue) Update (key interface{}, priority int64){
	item := ipq.index[key]
	ipq.priorityQueue.update(item, priority)
}

func (ipq obsolescenceQueue) Len () int {
	return ipq.priorityQueue.Len()
}

func (ipq obsolescenceQueue) Pop () (interface{}, interface{}) {
	item := heap.Pop(ipq.priorityQueue).(*item)
	delete(ipq.index, item.key)
	return item.key, item.value
}

func (ipq obsolescenceQueue) Remove (key interface{}) (interface{}, interface{}) {
	item := ipq.index[key]
	heap.Remove(ipq.priorityQueue, item.index)
	delete(ipq.index, item.key)
	return item.key, item.value
}

func (ipq obsolescenceQueue) Clear (timestamp int64) map[interface{}]interface{} {
	removed := make(map[interface{}]interface{})
	for ipq.priorityQueue.peek().(*item).timestamp < timestamp {
		key, value := ipq.Pop()
		removed[key] = value
	}
	return removed
}
