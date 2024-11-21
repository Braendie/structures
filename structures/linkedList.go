package structures

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	head *linkedListNode
	last *linkedListNode
	len  int
}

type linkedListNode struct {
	val  interface{}
	next *linkedListNode
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) AddToHead(vl interface{}) {
	newHead := &linkedListNode{
		val:  vl,
		next: l.head,
	}

	l.head = newHead
	if l.len == 0 {
		l.last = newHead
	}
	l.len++
}

func (l *LinkedList) AddToLast(vl interface{}) {
	newLast := &linkedListNode{
		val: vl,
	}
	if l.len == 0 {
		l.head = newLast
		l.last = newLast
	} else {
		l.last.next = newLast
		l.last = newLast
	}
	l.len++
}

func (l *LinkedList) FindVal(vl interface{}) *linkedListNode {
	curNode := l.head
	for curNode != nil {
		if curNode.val == vl {
			return curNode
		}
		curNode = curNode.next
	}
	return nil
}

func (l *LinkedList) DeleteValue(vl interface{}) error {
	if l.len == 0 {
		return errors.New("linked list is empty")
	}

	if l.head.val == vl {
		l.head = l.head.next
		if l.len == 1 {
			l.last = nil
		}
		l.len--
		return nil
	}

	curNode := l.head
	for curNode.next != nil {
		if curNode.next.val == vl {
			if curNode.next == l.last {
				l.last = curNode
			}
			curNode.next = curNode.next.next
			l.len--
			return nil
		}
		curNode = curNode.next
	}

	return nil
}

func (l *LinkedList) Print() {
	curNode := l.head
	for curNode != nil {
		fmt.Printf("%v -> ", curNode.val)
		curNode = curNode.next
	}
	fmt.Println("nil")
}
