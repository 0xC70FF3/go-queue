package queue

import (
	"testing"
)

func TestObsolescenceQueue(t *testing.T) {
	ipq := ObsolescenceQueue()

	ipq.Push("b", "banana", 3)
	ipq.Push("a", "apple", 2)
	ipq.Push("p", "pear", 4)

	ipq.Update("p", 1)

	ipq.Remove("a")

	key, value := ipq.Pop()
	if !(key == "p" && value == "pear") {
		t.Error("Expected value to be %s:%s was %s; %s", "p", "pear", key, value)
	}
	key, value = ipq.Pop()
	if !(key == "b" && value == "banana") {
		t.Error("Expected value to be %s:%s was %s; %s", "b", "banana", key, value)
	}
	if ipq.Len() > 0 {
		t.Error("Expected length to be 0")
	}
}


func TestObsolescenceQueue_Clear(t *testing.T) {
	ipq := ObsolescenceQueue()

	ipq.Push("b", "banana", 3)
	ipq.Push("a", "apple", 2)
	ipq.Push("p", "pear", 4)

	ipq.Clear(4)

	key, value := ipq.Pop()
	if !(key == "p" && value == "pear") {
		t.Error("Expected value to be %s:%s was %s; %s", "p", "pear", key, value)
	}
	if ipq.Len() > 0 {
		t.Error("Expected length to be 0")
	}
}