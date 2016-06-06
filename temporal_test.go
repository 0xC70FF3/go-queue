package queue

import (
	"testing"
	"container/heap"
)

func TestTemporalQueue(t *testing.T) {
	items := map[string]int{
		"banana": 3,
		"apple": 2,
		"pear": 4,
	}

	pq := make(temporalQueue, len(items))
	i := 0
	for value, timestamp := range items {
		pq[i] = &item{value: value, timestamp: int64(timestamp), index: i}
		i++
	}
	heap.Init(&pq)

	x := &item{value: "orange", timestamp: 1}
	heap.Push(&pq, x)

	pq.update(x, 5)

	x = pq.peek().(*item)
	if !(x.timestamp == 2 && x.value == "apple") {
		t.Error("Expected value to be %d:%s was %d; %s", 2, "apple", x.timestamp, x.value)
	}
	x = heap.Pop(&pq).(*item)
	if !(x.timestamp == 2 && x.value == "apple") {
		t.Error("Expected value to be %d:%s was %d; %s", 2, "apple", x.timestamp, x.value)
	}
	x = pq.peek().(*item)
	if !(x.timestamp == 3 && x.value == "banana") {
		t.Error("Expected value to be %d:%s was %d; %s", 3, "banana", x.timestamp, x.value)
	}
	x = heap.Pop(&pq).(*item)
	if !(x.timestamp == 3 && x.value == "banana") {
		t.Error("Expected value to be %d:%s was %d; %s", 3, "banana", x.timestamp, x.value)
	}
	x = pq.peek().(*item)
	if !(x.timestamp == 4 && x.value == "pear") {
		t.Error("Expected value to be %d:%s was %d; %s", 4, "pear", x.timestamp, x.value)
	}
	x = heap.Pop(&pq).(*item)
	if !(x.timestamp == 4 && x.value == "pear") {
		t.Error("Expected value to be %d:%s was %d; %s", 4, "pear", x.timestamp, x.value)
	}
	x = pq.peek().(*item)
	if !(x.timestamp == 5 && x.value == "orange") {
		t.Error("Expected value to be %d:%s was %d; %s", 5, "orange", x.timestamp, x.value)
	}
	x = heap.Pop(&pq).(*item)
	if !(x.timestamp == 5 && x.value == "orange") {
		t.Error("Expected value to be %d:%s was %d; %s", 5, "orange", x.timestamp, x.value)
	}
	if pq.Len() > 0 {
		t.Error("Expected length to be 0")
	}
}
