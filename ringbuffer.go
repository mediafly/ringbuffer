package ringbuffer

import (
	"fmt"
)

type StringBuffer interface {
	Append(string)
	Slice() []string
	Length() int
}

type stringBuffer struct {
	capacity int
	position int
	buffer   []string
}

func NewStringBuffer(capacity int) StringBuffer {
	if capacity < 1 {
		panic(fmt.Sprintf("capacity must be >= 1 but was %v", capacity))
	}

	return &stringBuffer{capacity: capacity}
}

func (b *stringBuffer) Append(value string) {
	if len(b.buffer) < b.capacity {
		b.buffer = append(b.buffer, value)
	} else {
		b.buffer[b.position] = value
	}

	b.position = (b.position + 1) % b.capacity
}

func (b *stringBuffer) Slice() []string {
	if len(b.buffer) == 0 {
		return []string{}
	}

	if len(b.buffer) < b.capacity {
		return b.buffer[0:len(b.buffer)]
	}

	buffer := make([]string, b.capacity, b.capacity)
	position := 0

	for i := b.position; i < b.capacity; i++ {
		buffer[position] = b.buffer[i]
		position++
	}

	for i := 0; i < b.position; i++ {
		buffer[position] = b.buffer[i]
		position++
	}

	return buffer
}

func (b *stringBuffer) Length() int {
	return len(b.buffer)
}
