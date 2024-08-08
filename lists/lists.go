package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()
	Push(1, queue)
	Push(2, queue)
	Push(3, queue)
	printQueue(queue) // 123
	//Pop(queue)
	//printQueue(queue) // в ту же строку: 23
	printQueue(ReverseList(queue))
}

func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for el := l.Front(); el != nil; el = el.Next() {
		reversedList.PushFront(el.Value)
	}
	return reversedList
}

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	if queue.Front() == nil {
		return nil
	}
	return queue.Remove(queue.Front())
}

func printQueue(queue *list.List) {
	for elem := queue.Front(); elem != nil; elem = elem.Next() {
		fmt.Print(elem.Value)
	}
}
