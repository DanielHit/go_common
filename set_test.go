package go_common

import (
	"github/testify/assert"
	"testing"
)

func TestSet_IsEmpty(t *testing.T) {
	s := NewSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	is := assert.New(t)
	is.True(s.Contains(1))
	is.False(s.Contains(4))
	s.Remove(1)
	is.False(s.Contains(1))
}
