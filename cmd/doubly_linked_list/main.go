package main

import "fmt"

type node[T comparable] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

type doublyLinkedList[T comparable] struct {
	Length int
	head   *node[T]
	tail   *node[T]
}

func NewDoublyLinkedList[T comparable]() *doublyLinkedList[T] {
	return &doublyLinkedList[T]{}
}

func (d *doublyLinkedList[T]) prepend(value T) {
	node := &node[T]{value: value}

	if d.head == nil {
		d.tail = node
		d.head = d.tail
		return
	}

	node.next = d.head
	d.head.prev = node
	d.head = node

	d.Length++
}

func (d *doublyLinkedList[T]) insertAt(value T, idx int) {
	if idx > d.Length {
		fmt.Println("An error occurred")
		return
	} else if idx == d.Length {
		d.append(value)
		return
	} else if idx == 0 {
		d.prepend(value)
		return
	}

	curr := d.head
	for i := 0; curr != nil && i < idx; i++ {
		curr = curr.next
	}

	node := &node[T]{value: value}

	node.next = curr
	node.prev = curr.prev
	curr.prev = node

	if curr.prev != nil {
		curr.prev.next = curr
	}

	d.Length++
}

func (d *doublyLinkedList[T]) append(value T) {
	node := &node[T]{value: value}

	if d.tail == nil {
		d.tail = node
		d.head = d.tail
		return
	}

	node.prev = d.tail
	d.tail.next = node
	d.tail = node

	d.Length++
}

func (d *doublyLinkedList[T]) remove(value T) T {
	curr := d.head
	for i := 0; curr != nil && i < d.Length; i++ {
		if curr.value == value {
			break
		}
		curr = curr.next
	}

	if curr == nil {
		var zero T
		return zero
	}

	d.Length--

	if d.Length == 0 {
		out := d.head.value
		d.tail = nil
		d.head = d.tail
		return out
	}

	if curr.prev != nil {
		curr.prev = curr.next
	}

	if curr.next != nil {
		curr.next = curr.prev
	}

	if curr == d.head {
		d.head = curr.next
	}

	if curr == d.tail {
		d.tail = curr.prev
	}

	curr.next = nil
	curr.prev = curr.next

	return curr.value
}

// TODO
func (d *doublyLinkedList[T]) get(idx int) T {

	var zero T
	return zero
}

// TODO
func (d *doublyLinkedList[T]) removeAt(idx int) T {

	var zero T
	return zero
}

func main() {
	fmt.Println("Linked list")
}
