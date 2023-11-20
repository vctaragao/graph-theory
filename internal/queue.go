package internal

type Queue[T any] []T

func (q *Queue[T]) Pop() T {
	queue := *q
	first, queue := queue[0], queue[1:]
	*q = queue

	return first
}
