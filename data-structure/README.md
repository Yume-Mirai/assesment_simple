# Circular Queue Implementation Challenges

This section contains a set of coding challenges focused on Circular Queue data structure implementation. The challenges are designed to test your understanding of FIFO operations, capacity management, and edge case handling.

## Overview

The implementation is in Go and focuses on a Circular Queue with fixed capacity. Each challenge builds upon the previous ones, testing different aspects of queue operations.

## Technical Requirements

- Go programming language
- Understanding of array operations
- Knowledge of queue properties (FIFO)
- Basic understanding of time complexity

## Challenges

### 1. Queue Initialization
Implement the `NewCircularQueue()` function to create a queue with fixed capacity.
- Initialize internal array with given capacity
- Set up front and rear pointers
- Expected time complexity: O(1)

### 2. Enqueue Operation
Implement the `Enqueue()` method to add elements to the queue.
- Handle queue overflow condition
- Update rear pointer correctly
- Maintain circular nature of the queue
- Expected time complexity: O(1)

### 3. Dequeue Operation
Implement the `Dequeue()` method to remove elements from the queue.
- Handle queue underflow condition
- Update front pointer correctly
- Return the dequeued element
- Expected time complexity: O(1)

### 4. Peek Operation
Implement the `Peek()` method to view the front element.
- Return front element without removing it
- Handle empty queue case
- Expected time complexity: O(1)

### 5. Status Checks
Implement the `IsFull()` and `IsEmpty()` methods.
- Correctly determine queue state
- Handle edge cases
- Expected time complexity: O(1)

## Testing

To verify your implementation:
```bash
go test
```

## Example Usage

```go
// Create a queue with capacity 5
q := NewCircularQueue(5)

// Add elements
q.Enqueue(1) // [1]
q.Enqueue(2) // [1,2]
q.Enqueue(3) // [1,2,3]

// Remove front element
val := q.Dequeue() // val = 1, queue = [2,3]

// Check front element
front := q.Peek() // front = 2
```

## Evaluation Criteria

- Correctness of implementation
- Proper handling of edge cases
- Code efficiency (constant time operations)
- Code readability and organization
- Proper use of Go idioms and best practices