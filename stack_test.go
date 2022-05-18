package go_common

import (
	"github/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push("hello")
	is := assert.New(t)
	is.Equal(s.Peak(), "hello")
	is.Equal(s.Pop(), "hello")
	is.Equal(s.Pop(), 1)
	is.Equal(s.Pop(), nil)
	is.True(s.IsEmpty())
}
