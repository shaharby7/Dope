package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Set(t *testing.T) {
	set := NewSet[string]()
	set.Add("one")
	assert.Equal(t, 1, set.Size())
	set.Add("one")
	assert.Equal(t, 1, set.Size())
	set.Add("two")
	assert.Equal(t, 2, set.Size())
	assert.True(t, set.Exists("one"))
}

func Test_FromSlice(t *testing.T) {
	slice := []int{1, 2, 2}
	set := NewSet(OptionFromSlice(slice))
	assert.True(t, set.Exists(1))
	assert.True(t, set.Exists(2))
	assert.Equal(t, 2, set.Size())
}
