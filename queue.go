package gqueue

import (
	"bytes"
	"fmt"
)

type node[T any] struct {
	Value T
	Next  *node[T]
}

// GQueue is generic Queue implementation using linked list
type GQueue[T any] struct {
	len  int
	head *node[T]
	rear *node[T]
}

// New creates an instance of GQueue and pushes all given items to the created queue
func New[T any](items ...T) GQueue[T] {
	s := GQueue[T]{}
	for _, item := range items {
		s.Push(item)
	}
	return s
}

// Len returns the length of the stack
func (s *GQueue[T]) Len() int {
	return s.len
}

// Peek returns the first item without removing it or a zero object of given type
func (s *GQueue[T]) Peek() T {
	if s.head == nil {
		var zero T
		return zero
	}
	return s.head.Value
}

// Push puts the given item into the queue
func (s *GQueue[T]) Push(item T) {
	n := &node[T]{
		Value: item,
	}
	if s.rear == nil {
		s.head = n
	} else {
		s.rear.Next = n
	}
	s.rear = n
	s.len++
}

// Pop returns the first item and removes it from the queue or returns a zero object of given type
func (s *GQueue[T]) Pop() T {
	if s.head == nil {
		var zero T
		return zero
	}

	item := s.head.Value
	s.head = s.head.Next
	if s.head == nil {
		s.rear = nil
	}
	s.len--
	return item
}

// String returns the list of all items in text format
func (s *GQueue[T]) String() string {
	if s.len == 0 {
		return ""
	}
	var buf bytes.Buffer
	head := s.head
	for head != nil {
		buf.WriteString(fmt.Sprintf("%v", head.Value))
		if head.Next != nil {
			buf.WriteString(", ")
		}
		head = head.Next
	}
	return buf.String()
}

// Clear clears the queue
func (s *GQueue[T]) Clear() {
	s.head = nil
	s.rear = nil
	s.len = 0
}
