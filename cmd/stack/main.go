package main

type node[T any] struct {
	value T
	prev  *node[T]
}

type stack[T any] struct {
	Length int
	head   *node[T]
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) push(value T) {
	node := &node[T]{value: value}

	if s.head == nil {
		s.head = node
	} else {
		node.prev = s.head
		s.head = node
	}

	s.Length++
}

func (s *stack[T]) pop() T {
	if s.head == nil {
		var zero T
		return zero
	}

	head := s.head
	s.head = head.prev

	// Optional Clean up
	head.prev = nil

	s.Length--

	return head.value
}

func (s *stack[T]) peek() T {
	if s.head == nil {
		var zero T
		return zero
	}

	return s.head.value
}
