package list

import (
	"fmt"
	"testing"
)

func TestMakeBySlice(t *testing.T) {
	s := make([]int, 3)
	s[0] = 0
	s[1] = 1
	s[2] = 2

	l := MakeBySlice(s)

	one := l.first
	fmt.Println(one.Prev)
	assertNil(t, one.Prev)
	assertSame(t, 0, one.Value)

	two := one.Next
	assertSame(t, one, two.Prev)
	assertSame(t, 1, two.Value)

	three := two.Next
	assertSame(t, two, three.Prev)
	assertSame(t, 2, three.Value)
	assertNil(t, three.Next)
}



func TestList_GetSlice(t *testing.T) {
	one := Node{1, nil, nil}
	two := Node{2, nil, nil}
	three := Node{3, nil, nil}

	l := List{&one}
	one.Next = &two
	two.Prev = &one
	two.Next = &three
	three.Prev = &two

	s := make([]int, 3);
	s[0] = 1
	s[1] = 2
	s[2] = 3

	ls := l.GetSlice()

	assertSame(t, s[0], ls[0])
	assertSame(t, s[1], ls[1])
	assertSame(t, s[2], ls[2])
}