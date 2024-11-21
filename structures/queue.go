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

func start() {
	queue := &Queue{}

	// Добавляем элементы
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	// Удаляем элементы
	val, _ := queue.Dequeue()
	fmt.Println(val) // 10

	val, _ = queue.Dequeue()
	fmt.Println(val) // 20

	// Проверяем пустоту
	fmt.Println(queue.IsEmpty()) // false

	val, _ = queue.Dequeue()
	fmt.Println(val) // 30

	fmt.Println(queue.IsEmpty()) // true
}
