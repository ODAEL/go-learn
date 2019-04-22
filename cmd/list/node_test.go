package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_SetNext(t *testing.T) {

	t.Run("Node", func(t *testing.T) {
		node := Node{1, nil, nil}
		next := Node{2, nil, nil}
		node.SetNext(&next);

		assert.Equal(t, &next, node.Next)
		assert.Equal(t, &node, node.Next.Prev)
	})

	t.Run("Nil", func(t *testing.T) {
		node := Node{1, nil, nil}
		node.SetNext(nil);

		assert.Nil(t, node.Next)
	})
}

func TestNode_SetPrev(t *testing.T) {

	t.Run("Node", func(t *testing.T) {
		node := Node{1, nil, nil}
		prev := Node{2, nil, nil}
		node.SetPrev(&prev);

		assert.Equal(t, &prev, node.Prev)
		assert.Equal(t, &node, node.Prev.Next)
	})

	t.Run("Nil", func(t *testing.T) {
		node := Node{1, nil, nil}
		node.SetPrev(nil);

		assert.Nil(t, node.Prev)
	})

}

func TestSwapValues(t *testing.T) {
	one := Node{1, nil, nil}
	two := Node{2, nil, nil}

	SwapValues(&one, &two)

	assert.Equal(t, 2, one.Value)
	assert.Equal(t, 1, two.Value)
}

func TestSwap(t *testing.T) {

	t.Run("SwapNearLeftRight", func (t *testing.T) {
		one := Node{1, nil, nil}
		two := Node{2, nil, nil}

		one.SetNext(&two)

		Swap(&one, &two)

		assert.Equal(t, one.Prev, &two)
		assert.Equal(t, two.Next, &one)

		assert.Nil(t, two.Prev)
		assert.Nil(t, one.Next)
	})

	t.Run("SwapNearRightLeft", func (t *testing.T) {
		one := Node{1, nil, nil}
		two := Node{2, nil, nil}

		one.SetNext(&two)

		Swap(&two, &one)

		assert.Equal(t, one.Prev, &two)
		assert.Equal(t, two.Next, &one)

		assert.Nil(t, two.Prev)
		assert.Nil(t, one.Next)
	})

	t.Run("SwapSeparateEnds", func (t *testing.T) {
		one := Node{1, nil, nil}
		two := Node{2, nil, nil}
		three := Node{3, nil, nil}

		one.SetNext(&two)
		two.SetNext(&three)

		Swap(&one, &three)

		assert.Equal(t, one.Prev, &two)
		assert.Equal(t, three.Next, &two)

		assert.Nil(t, one.Next)
		assert.Nil(t, three.Prev)
	})

	t.Run("SwapSeparateInside", func (t *testing.T) {
		one := Node{1, nil, nil}
		two := Node{2, nil, nil}
		three := Node{3, nil, nil}
		four := Node{4, nil, nil}
		five := Node{5, nil, nil}

		one.SetNext(&two)
		two.SetNext(&three)
		three.SetNext(&four)
		four.SetNext(&five)

		Swap(&two, &four)

		assert.Equal(t, two.Prev, &three)
		assert.Equal(t, four.Next, &three)

		assert.Equal(t, two.Next, &five)
		assert.Equal(t, four.Prev, &one)
	})
}

func TestNode_String(t *testing.T) {
	node := Node{1, nil, nil}

	assert.Equal(t, "1", fmt.Sprint(node))
}
