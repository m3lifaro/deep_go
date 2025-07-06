package homework2

type CircularQueue[T int8 | int16 | int32 | int64] struct {
	values []T
	front  int
	rear   int
	size   int
}

func NewCircularQueue[T int8 | int16 | int32 | int64](size int) CircularQueue[T] {
	return CircularQueue[T]{
		values: make([]T, size),
		front:  0,
		rear:   -1,
	}
}

func (q *CircularQueue[T]) Push(value T) bool {
	if !q.Full() {
		q.rear = (q.rear + 1) % len(q.values)
		q.values[q.rear] = value
		q.size++
		return true
	}
	return false
}

func (q *CircularQueue[T]) Pop() bool {
	if !q.Empty() {
		q.front = (q.front + 1) % len(q.values)
		q.size--
		if q.size == 0 {
			q.front = 0
			q.rear = -1
		}
		return true
	}
	return false
}

func (q *CircularQueue[T]) Front() T {
	if !q.Empty() {
		return q.values[q.front]
	}
	return -1
}

func (q *CircularQueue[T]) Back() T {
	if !q.Empty() {
		return q.values[q.rear]
	}
	return -1
}

func (q *CircularQueue[T]) Empty() bool {
	return q.size == 0
}

func (q *CircularQueue[T]) Full() bool {
	if q.size == len(q.values) {
		return true
	}
	return false
}
