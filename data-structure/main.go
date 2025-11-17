package main

import "fmt"

func main() {
	fmt.Println("=== Circular Queue Demo ===")
	fmt.Println()

	// This is a data structure interview problem set focusing on Circular Queue
	// There are 5 challenges that test different aspects of queue operations:

	// 1. Implement NewCircularQueue() to create a queue with fixed capacity
	// 2. Implement Enqueue() to add elements while handling overflow
	// 3. Implement Dequeue() to remove elements while handling underflow
	// 4. Implement Peek() to view the front element without removing it
	// 5. Implement IsFull() and IsEmpty() to check queue status

	// Example usage:
	fmt.Println("1. Membuat Circular Queue dengan kapasitas 5")
	q := NewCircularQueue(5)
	fmt.Printf("Queue dibuat dengan kapasitas %d\n", len(q.items))
	fmt.Printf("Status: Empty=%t, Full=%t, Size=%d\n", q.IsEmpty(), q.IsFull(), q.size)
	fmt.Println()

	// Enqueue some values
	fmt.Println("2. Enqueue nilai 1, 2, 3, 4, 5")
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		success := q.Enqueue(v)
		fmt.Printf("Enqueue(%d): Success=%t, Size=%d, Empty=%t, Full=%t\n", v, success, q.size, q.IsEmpty(), q.IsFull())
	}
	fmt.Println()

	// Test peek
	fmt.Println("3. Peek elemen depan")
	if val, ok := q.Peek(); ok {
		fmt.Printf("Peek(): Value=%d, Size=%d (tidak berkurang)\n", val, q.size)
	} else {
		fmt.Println("Peek(): Queue kosong")
	}
	fmt.Println()

	// Dequeue some values
	fmt.Println("4. Dequeue 3 elemen")
	for i := 0; i < 3; i++ {
		if val, ok := q.Dequeue(); ok {
			fmt.Printf("Dequeue(): Value=%d, Size=%d, Empty=%t, Full=%t\n", val, q.size, q.IsEmpty(), q.IsFull())
		} else {
			fmt.Println("Dequeue(): Queue kosong")
		}
	}
	fmt.Println()

	// Enqueue more to test circular behavior
	fmt.Println("5. Enqueue lagi untuk test circular behavior")
	newValues := []int{6, 7}
	for _, v := range newValues {
		success := q.Enqueue(v)
		fmt.Printf("Enqueue(%d): Success=%t, Size=%d, Empty=%t, Full=%t\n", v, success, q.size, q.IsEmpty(), q.IsFull())
	}
	fmt.Println()

	// Dequeue remaining
	fmt.Println("6. Dequeue semua elemen tersisa")
	for !q.IsEmpty() {
		if val, ok := q.Dequeue(); ok {
			fmt.Printf("Dequeue(): Value=%d, Size=%d, Empty=%t, Full=%t\n", val, q.size, q.IsEmpty(), q.IsFull())
		}
	}
	fmt.Println()

	// Test edge cases
	fmt.Println("7. Test edge cases")
	fmt.Println("Coba dequeue dari queue kosong:")
	if _, ok := q.Dequeue(); !ok {
		fmt.Println("Dequeue(): Gagal - Queue kosong")
	}

	fmt.Println("Coba peek dari queue kosong:")
	if _, ok := q.Peek(); !ok {
		fmt.Println("Peek(): Gagal - Queue kosong")
	}

	fmt.Println()
	fmt.Println("=== Demo selesai ===")
	fmt.Println("Run 'go test' to verify your implementation")
}
