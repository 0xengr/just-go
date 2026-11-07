package main

import "fmt"

// NOTE: Exercise: a generic queue.
// * A queue is FIFO: the value added first comes out first.
// * Queue[T] stores a slice of T and repeats [T] on every method receiver.
// * Pop returns (T, bool) so the caller can tell an empty queue from a stored
// * zero value.

type Queue[T any] struct {
	values []T
}

func (q *Queue[T]) Push(value T) {
	q.values = append(q.values, value)
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.values) == 0 {
		var zero T
		return zero, false
	}
	value := q.values[0]
	q.values = q.values[1:]
	return value, true
}

func (q *Queue[T]) Len() int {
	return len(q.values)
}

func main() {
	var numbers Queue[int]
	numbers.Push(1)
	numbers.Push(2)
	numbers.Push(3)

	for numbers.Len() > 0 {
		value, ok := numbers.Pop()
		fmt.Println(value, ok)
	}

	_, ok := numbers.Pop()
	fmt.Println("empty queue:", ok)

	var words Queue[string]
	words.Push("first")
	words.Push("second")
	first, _ := words.Pop()
	second, _ := words.Pop()
	fmt.Println(first, second)
}
