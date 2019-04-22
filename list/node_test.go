package list

import (
	"fmt"
	"reflect"
	"testing"
)

// TODO separate
func assertSame(t *testing.T, want, got interface{})  {
	if want != got {
		t.Errorf("Expected '%s', got '%s'", want, got)
	}
}

func assertNil(t *testing.T, got interface{})  {
	if got == nil {
		return
	}

	if reflect.ValueOf(got).IsNil() {
		return
	}

	t.Errorf("Expected <nil>, got '%s'", got)
}

func TestNode_SetNext(t *testing.T) {
	node := Node{1, nil, nil}
	next := Node{2, nil, nil}
	node.SetNext(&next);

	assertSame(t, &next, node.Next)
}

func TestNode_SetPrev(t *testing.T) {
	node := Node{1, nil, nil}
	prev := Node{2, nil, nil}
	node.SetPrev(&prev);

	assertSame(t, &prev, node.Prev)
}

func TestNode_String(t *testing.T) {
	node := Node{1, nil, nil}

	assertSame(t, "1", fmt.Sprint(node))
}

func TestSwapValues(t *testing.T) {
	one := Node{1, nil, nil}
	two := Node{2, nil, nil}

	SwapValues(&one, &two)

	assertSame(t, 2, one.Value)
	assertSame(t, 1, two.Value)
}

func TestSwap(t *testing.T) {

	t.Run("SwapNearNM", func (t *testing.T) {
		one := Node{1, nil, nil}
		two := Node{2, nil, nil}

		one.Next = &two
		two.Prev = &one

		Swap(&one, &two)

		assertSame(t, one.Prev, &two)
		assertSame(t, two.Next, &one)
	})

	t.Run("SwapNearMN", func (t *testing.T) {
		one := Node{1, nil, nil}
		two := Node{2, nil, nil}

		one.Next = &two
		two.Prev = &one

		Swap(&two, &one)

		assertSame(t, one.Prev, &two)
		assertSame(t, two.Next, &one)
	})

	t.Run("SwapSeparate", func (t *testing.T) {
		one := Node{1, nil, nil}
		two := Node{2, nil, nil}
		three := Node{3, nil, nil}

		one.Next = &two
		two.Prev = &one
		two.Next = &three
		three.Prev = &two

		Swap(&one, &three)

		assertSame(t, one.Prev, &two)
		assertSame(t, three.Next, &two)

		assertSame(t, one.Prev.Prev, &three)
		assertSame(t, three.Next.Next, &one)
	})
}