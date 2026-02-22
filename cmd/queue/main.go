package main

type node[T any] struct {
	data T
	next *node[T]
}

type queue[T any] struct {
	Length int
	head   *node[T]
	tail   *node[T]
}

// Constructor
func NewQueue[T any]() *queue[T] {
	return &queue[T]{}
}

// Add data to the queue
func (q *queue[T]) enqueue(data T) {
	node := &node[T]{data: data}

	if q.tail == nil {
		// Add first node
		q.head = node
		q.tail = q.head
	} else {
		q.tail.next = node
		q.tail = node
	}

	q.Length++
}

// Removing data from the queue
func (q *queue[T]) deque() T {
	if q.head == nil {
		var zero T
		return zero
	}

	head := q.head
	q.head = q.head.next

	// Optional Clean up
	head.next = nil

	q.Length--

	return head.data
}

func (q *queue[T]) peek() T {
	if q.head == nil {
		var zero T
		return zero
	}

	return q.head.data
}
