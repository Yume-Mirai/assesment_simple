package main

import "testing"

func TestNewCircularQueue(t *testing.T) {
	q := NewCircularQueue(5)
	if q == nil {
		t.Error("NewCircularQueue(5) returned nil")
	}
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	if q.IsFull() {
		t.Error("New queue should not be full")
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := NewCircularQueue(3)

	// Test enqueue
	if !q.Enqueue(1) {
		t.Error("Failed to enqueue 1")
	}
	if !q.Enqueue(2) {
		t.Error("Failed to enqueue 2")
	}
	if !q.Enqueue(3) {
		t.Error("Failed to enqueue 3")
	}
	// Queue should be full now
	if !q.IsFull() {
		t.Error("Queue should be full after enqueueing 3 elements")
	}
	// Should not be able to enqueue when full
	if q.Enqueue(4) {
		t.Error("Should not be able to enqueue when full")
	}

	// Test dequeue
	val, ok := q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("Dequeue() = %v, %v; want 1, true", val, ok)
	}
	val, ok = q.Dequeue()
	if !ok || val != 2 {
		t.Errorf("Dequeue() = %v, %v; want 2, true", val, ok)
	}
	val, ok = q.Dequeue()
	if !ok || val != 3 {
		t.Errorf("Dequeue() = %v, %v; want 3, true", val, ok)
	}
	// Queue should be empty now
	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
	// Should not be able to dequeue when empty
	_, ok = q.Dequeue()
	if ok {
		t.Error("Should not be able to dequeue when empty")
	}
}

func TestPeek(t *testing.T) {
	q := NewCircularQueue(3)

	// Peek on empty queue
	_, ok := q.Peek()
	if ok {
		t.Error("Peek on empty queue should return false")
	}

	// Add some elements
	q.Enqueue(1)
	q.Enqueue(2)

	// Test peek
	val, ok := q.Peek()
	if !ok || val != 1 {
		t.Errorf("Peek() = %v, %v; want 1, true", val, ok)
	}

	// Peek shouldn't remove the element
	val, ok = q.Dequeue()
	if !ok || val != 1 {
		t.Errorf("After Peek, Dequeue() = %v, %v; want 1, true", val, ok)
	}
}

func TestCircularBehavior(t *testing.T) {
	q := NewCircularQueue(3)

	// Fill the queue
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	// Remove two elements
	q.Dequeue()
	q.Dequeue()

	// Add elements to wrap around
	if !q.Enqueue(4) {
		t.Error("Failed to enqueue 4")
	}
	if !q.Enqueue(5) {
		t.Error("Failed to enqueue 5")
	}

	// Verify circular behavior
	val, ok := q.Dequeue()
	if !ok || val != 3 {
		t.Errorf("Dequeue() = %v, %v; want 3, true", val, ok)
	}
	val, ok = q.Dequeue()
	if !ok || val != 4 {
		t.Errorf("Dequeue() = %v, %v; want 4, true", val, ok)
	}
	val, ok = q.Dequeue()
	if !ok || val != 5 {
		t.Errorf("Dequeue() = %v, %v; want 5, true", val, ok)
	}
}
