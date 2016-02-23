package ringbuffer

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestStringBuffer(t *testing.T) {
	buffer := NewStringBuffer(3)
	assert.Equal(t, []string{}, buffer.Slice())

	buffer.Append("a")
	assert.Equal(t, []string{"a"}, buffer.Slice())

	buffer.Append("b")
	assert.Equal(t, []string{"a", "b"}, buffer.Slice())

	buffer.Append("c")
	assert.Equal(t, []string{"a", "b", "c"}, buffer.Slice())

	buffer.Append("d")
	assert.Equal(t, []string{"b", "c", "d"}, buffer.Slice())

	buffer.Append("e")
	assert.Equal(t, []string{"c", "d", "e"}, buffer.Slice())

	buffer = NewStringBuffer(10)
	for i := 0; i < 25; i++ {
		buffer.Append(strconv.Itoa(i))
	}
	expected := []string{}
	for i := 15; i < 25; i++ {
		expected = append(expected, strconv.Itoa(i))
	}
	assert.Equal(t, expected, buffer.Slice())
}
