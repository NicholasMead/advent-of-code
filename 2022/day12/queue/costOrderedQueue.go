package queue

import (
	"fmt"
	"strings"
)

type CostOrderedQueue[Item any, Cost int | float64] interface {
	Push(Item, Cost)
	Pop() (Item, Cost, bool)
	
}

func CreateCostOrdered[Item any, Cost float64 | int]() CostOrderedQueue[Item, Cost] {
	return &costOrderedQueue[Item, Cost]{
		items: []costedItem[Item, Cost]{},
	}
}

type costedItem[Item any, Cost float64 | int] struct {
	Item Item
	cost Cost
}

type costOrderedQueue[Item any, Cost float64 | int] struct {
	items []costedItem[Item, Cost]
}

// pop implements Queue
func (q *costOrderedQueue[Item, Cost]) Pop() (Item, Cost, bool) {
	if len(q.items) < 1 {
		return *new(Item), 0, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item.Item, item.cost, true
}

// push implements Queue
func (q *costOrderedQueue[Item, Cost]) Push(item Item, cost Cost) {
	for idx, i := range q.items {
		if cost < i.cost {
			items := append([]costedItem[Item, Cost]{}, q.items[0:idx]...)
			items = append(items, costedItem[Item, Cost]{item, cost})
			q.items = append(items, q.items[idx:]...)
			return
		}
	}
	q.items = append(q.items, costedItem[Item, Cost]{item, cost})
}

func (q costOrderedQueue[Item, Cost]) String() string {
	items := make([]string, len(q.items))

	for idx, itm := range q.items {
		items[idx] = fmt.Sprintf("(%v: %v)", itm.Item, itm.cost)
	}

	return "[" + strings.Join(items, ", ") + "]"
}
