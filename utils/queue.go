package utils

import "slices"

type Queue[T comparable] struct {
	Items []T
}

func CreateQueue[T comparable]() *Queue[T] {
	queue := make([]T, 0)
	return &Queue[T]{Items: queue}
}

func (queue *Queue[T]) Pop() T {
	item := queue.Items[0]
	queue.Items = queue.Items[1:]
	return item
}

func (queue *Queue[T]) Push(items ...T) {
	queue.Items = append(queue.Items, items...)
}

func (queue *Queue[T]) PushNoDuplicate(items ...T) {
	for _, item := range items {
		if slices.Contains((*queue).Items, item) {
			continue
		}

		queue.Push(item)
	}
}

func (queue *Queue[T]) Len() int {
	return len(queue.Items)
}
