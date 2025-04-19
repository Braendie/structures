package structures

import (
	"errors"
	"fmt"
)

type Queue struct {
	items *LinkedList
}

func NewQueue() *Queue {
	return &Queue{
		items: &LinkedList{},
	}
}

func (q *Queue) IsEmpty() bool {
	return q.items.len == 0
}

func (q *Queue) Enqueue(val interface{}) {
	q.items.AddToLast(val)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	result := q.items.head.val
	q.items.DeleteValue(result)
	return result, nil
}
