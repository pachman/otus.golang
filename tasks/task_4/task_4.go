package main

import (
	"fmt"
)

type Item struct {
	value interface{}
	next  *Item
	prev  *Item
}

func (i *Item) Value() interface{} {
	return i.value
}

func (i *Item) Next() *Item {
	return i.next
}

func (i *Item) Prev() *Item {
	return i.prev
}

type List struct {
	first *Item
	last  *Item
	count int
}

func (l List) Len() int {
	return l.count
}

func (l List) First() interface{} {
	if l.first != nil {
		return l.first.Value()
	} else {
		return nil
	}
}

func (l List) Last() interface{} {
	if l.last != nil {
		return l.last.Value()
	} else {
		return nil
	}
}

func (l *List) PushFront(value interface{}) {
	if l.first == nil {
		l.first = &Item{value: value}
		l.last = l.first
	} else {
		i := &Item{value: value, next: l.first}

		l.first.prev = i
		l.first = i
	}

	l.count++
}

func (l *List) PushBack(value interface{}) {
	if l.last == nil {
		l.first = &Item{value: value}
		l.last = l.first
	} else {
		i := &Item{value: value, prev: l.last}

		l.last.next = i
		l.last = i
	}

	l.count++
}

func (l *List) Remove(value interface{}) {
	if l.first == nil {
		return
	}

	current := l.first
	for i := 0; i < l.count; i++ {
		if value == current.Value() {
			if l.First() == l.Last() {
				l.last = nil
				l.first = nil
			} else if l.first == current {
				l.first = l.first.Next()
			} else if l.last == current {
				prev := l.last.prev
				prev.next = nil
				l.last = prev
			} else {
				prev := current.prev
				next := current.next

				prev.next = next
				next.prev = prev
			}
			break
		}
		current = current.Next()
	}

	l.count--
}

func (l *List) Print() {
	current := l.first

	fmt.Println()
	for i := 0; i < l.Len(); i++ {
		fmt.Println(current.Value())
		current = current.Next()
	}
}

func (l *List) Get(index int) interface{} {
	current := l.first

	if index < 0 && index >= l.Len() {
		return nil
	}

	for i := 0; i < l.Len(); i++ {
		if index == i {
			return current.Value()
		}
		current = current.Next()
	}

	return nil
}

func main() {
	list := List{}

	list.PushFront(111)
	list.PushFront(222)
	list.PushBack(333)

	list.Print()

	list.Remove(222)

	list.Print()
}
