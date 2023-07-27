package gqueue_test

import (
	"fmt"
	"github.com/sv-tools/gqueue"
	"testing"
)

func TestQueue(t *testing.T) {
	s := gqueue.New[int]()
	if l := s.Len(); l != 0 {
		t.Fatalf("expected empty queue, but got %d items", l)
	}
	if v := s.Pop(); v != 0 {
		t.Fatalf("expected zero item, but got %v", v)
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if l := s.Len(); l != 3 {
		t.Fatalf("expected three items, but got %d", l)
	}
	if v := s.Pop(); v != 1 {
		t.Fatalf("expected 3 as item, but got %v", v)
	}
	if l := s.Len(); l != 2 {
		t.Fatalf("expected two items, but got %d", l)
	}
	if v := s.Peek(); v != 2 {
		t.Fatalf("expected 2 as item, but got %v", v)
	}
	if v := s.String(); v != "2, 3" {
		t.Fatalf("expected '2, 3', but got %s", v)
	}
	s.Clear()
	if l := s.Len(); l != 0 {
		t.Fatalf("expected empty queue, but got %d items", l)
	}
}

func TestGSQueue_New(t *testing.T) {
	s := gqueue.New(1, 2, 3, 4)
	if l := s.Len(); l != 4 {
		t.Fatalf("expected four items, but got %d", l)
	}
	if v := s.Pop(); v != 1 {
		t.Fatalf("expected 1 as item, but got %v", v)
	}
	if v := s.Peek(); v != 2 {
		t.Fatalf("expected 2 as item, but got %v", v)
	}
}

func ExampleGQueue() {
	s := gqueue.New(1, 2, 3, 4)
	fmt.Println(s.Len())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(5)
	fmt.Println(s.Peek())
	// Output: 4
	// 1
	// 2
	// 3
	// 4
	// 5
}
