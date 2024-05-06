package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append([]int{item}, q.items...)
}

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1 // Retorna -1 se a fila estiver vazia
	}
	item := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	return item
}

func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}
