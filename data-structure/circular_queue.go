package main

// CircularQueue represents a fixed-size circular queue
type CircularQueue struct {
	items []int
	front int
	rear  int
	size  int
}

// NewCircularQueue creates a new circular queue with the given capacity
func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items: make([]int, capacity),
		front: 0,
		rear:  0,
		size:  0,
	}
}

// Enqueue adds an element to the rear of the queue
func (q *CircularQueue) Enqueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.items[q.rear] = value
	q.rear = (q.rear + 1) % len(q.items)
	q.size++
	return true
}

// Dequeue removes and returns the front element from the queue
func (q *CircularQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	value := q.items[q.front]
	q.front = (q.front + 1) % len(q.items)
	q.size--
	return value, true
}

// Peek returns the front element without removing it
func (q *CircularQueue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.items[q.front], true
}

// IsEmpty returns true if the queue is empty
func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
}

// IsFull returns true if the queue is full
func (q *CircularQueue) IsFull() bool {
	return q.size == len(q.items)
}
