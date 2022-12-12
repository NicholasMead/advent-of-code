package queue

import "testing"

func TestCostOrderedQueue(t *testing.T) {
	queue := CreateCostOrdered[string, int]()

	t.Log(queue)
	queue.Push("c", 3)
	t.Log(queue)
	queue.Push("b", 2)
	t.Log(queue)
	queue.Push("d", 4)
	t.Log(queue)
	queue.Push("a", 1)
	t.Log(queue)

	for i := 0; i < 4; i++ {
		item, cost, next := queue.Pop()

		if !next {
			t.Errorf("%d: Expected item", i)
		}
		if item != string(rune('a'+i)) {
			t.Errorf("%d: Expected item %v got %v", i, rune('a'+i), item)
		}
		if cost != i+1 {
			t.Errorf("%d: Expected cost %d for %v got %d", i, i+1, item, cost)
		}
	}

	if item, cost, next := queue.Pop(); next {
		t.Errorf("Expected end of queue, got (%v:%d)", item, cost)
	}
}
