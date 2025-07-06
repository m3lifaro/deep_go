package homework2

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	const queueSize = 3
	queue := NewCircularQueue[int16](queueSize)

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())

	assert.Equal(t, int16(-1), queue.Front())
	assert.Equal(t, int16(-1), queue.Back())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Push(1))
	assert.True(t, queue.Push(2))
	assert.True(t, queue.Push(3))
	assert.False(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int16{1, 2, 3}, queue.values))

	assert.False(t, queue.Empty())
	assert.True(t, queue.Full())

	assert.Equal(t, int16(1), queue.Front())
	assert.Equal(t, int16(3), queue.Back())

	assert.True(t, queue.Pop())
	assert.False(t, queue.Empty())
	assert.False(t, queue.Full())
	assert.True(t, queue.Push(4))

	assert.True(t, reflect.DeepEqual([]int16{4, 2, 3}, queue.values))

	assert.Equal(t, int16(2), queue.Front())
	assert.Equal(t, int16(4), queue.Back())

	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.True(t, queue.Pop())
	assert.False(t, queue.Pop())

	assert.True(t, queue.Empty())
	assert.False(t, queue.Full())
}
