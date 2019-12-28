package main

import "testing"

func TestLen(t *testing.T) {
	list := List{}

	list.PushFront(111)
	list.PushFront(222)
	list.PushFront(333)
	list.Remove(222)
	expected := 2
	actual := list.Len()

	if actual != expected {
		t.Errorf("actual=%v, expected=%v", actual, expected)
	}
}

func TestFirst(t *testing.T) {
	list := List{}

	list.PushFront(111)
	list.PushFront(222)

	expected := 222
	actual := list.First()

	if actual != expected {
		t.Errorf("actual=%v, expected=%v", actual, expected)
	}
}

func TestLast(t *testing.T) {
	list := List{}

	list.PushFront(111)
	list.PushFront(222)
	list.PushBack(333)

	expected := 333
	actual := list.Last()

	if actual != expected {
		t.Errorf("actual=%v, expected=%v", actual, expected)
	}
}

func TestPushFront(t *testing.T) {
	list := List{}

	list.PushFront(111)
	list.PushFront(222)
	list.PushFront(333)

	expected := 333
	actual := list.First()

	if actual != expected {
		t.Errorf("actual=%v, expected=%v", actual, expected)
	}
}

func TestPushBack(t *testing.T) {
	list := List{}

	list.PushFront(111)
	list.PushFront(333)
	list.PushBack(222)

	expected := 222
	actual := list.Last()

	if actual != expected {
		t.Errorf("actual=%v, expected=%v", actual, expected)
	}
}

func TestRemove(t *testing.T) {
	list := List{}

	list.PushFront(111)
	list.PushFront(222)
	list.PushBack(333)

	list.Remove(333)

	expected := 111
	actual := list.Last()

	if actual != expected {
		t.Errorf("actual=%v, expected=%v", actual, expected)
	}
}
