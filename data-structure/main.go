package main

import "fmt"

func main() {
	// This is a data structure interview problem set focusing on Circular Queue
	// There are 5 challenges that test different aspects of queue operations:

	// 1. Implement NewCircularQueue() to create a queue with fixed capacity
	// 2. Implement Enqueue() to add elements while handling overflow
	// 3. Implement Dequeue() to remove elements while handling underflow
	// 4. Implement Peek() to view the front element without removing it
	// 5. Implement IsFull() and IsEmpty() to check queue status

	// Example usage:
	q := NewCircularQueue(5)

	// Enqueue some values
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		q.Enqueue(v)
	}

	// Run test cases with: go test
	// Implement the methods in circular_queue.go
	fmt.Println("Run 'go test' to verify your implementation")
}
