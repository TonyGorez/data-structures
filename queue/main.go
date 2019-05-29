package main

import "fmt"

//Queue type
type Queue struct {
	items  []string
	length int
}

// Enqueue func
func (queue *Queue) Enqueue(item string) {

	queue.items = append(queue.items, item)
	queue.length++

}

// Dequeue func
func (queue *Queue) Dequeue() string {

	queue.items = queue.items[1:len(queue.items)]
	queue.length--
	return queue.items[0]

}

func main() {

	q := Queue{}
	q.Enqueue("Tony")
	q.Enqueue("Jamila")
	q.Enqueue("Zac")
	fmt.Println(q.items, q.length)
	q.Dequeue()
	fmt.Println(q.items, q.length)

}
