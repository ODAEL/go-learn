package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	list := Create(1, 2, 3)

	node := list.first
	assert.Nil(t, node.Prev)
	assert.Equal(t, 1, node.Value)

	node = node.Next
	assert.Equal(t, 2, node.Value)

	node = node.Next
	assert.Equal(t, 3, node.Value)
	assert.Nil(t, node.Next)
}

func TestCreateBySlice(t *testing.T) {
	slice := []int{1, 2, 3}

	list := CreateBySlice(slice)

	one := list.first
	assert.Nil(t, one.Prev)
	assert.Equal(t, 1, one.Value)

	two := one.Next
	assert.Equal(t, one, two.Prev)
	assert.Equal(t, 2, two.Value)

	three := two.Next
	assert.Equal(t, two, three.Prev)
	assert.Equal(t, 3, three.Value)
	assert.Nil(t, three.Next)
}

func TestList_GetSlice(t *testing.T) {
	list := Create(1, 2, 3)

	slice := []int{1, 2, 3};

	ls := list.GetSlice()

	assert.Equal(t, slice[0], ls[0])
	assert.Equal(t, slice[1], ls[1])
	assert.Equal(t, slice[2], ls[2])
}
