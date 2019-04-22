package list

import (
	"testing"
)

func TestSortEquals(t *testing.T) {
	s := make([]int, 4);
	s[0] = 3
	s[1] = 4
	s[2] = 1
	s[3] = 2

	l := MakeBySlice(s)

	lbs := l.BSort().GetSlice()
	lqs := l.QSort().GetSlice()

	assertSame(t, lbs[0], lqs[0])
	assertSame(t, lbs[1], lqs[1])
	assertSame(t, lbs[2], lqs[2])
	assertSame(t, lbs[3], lqs[3])
}

func generateBigSlice() []int {
	var s []int

	for i := 1000; i > 0; i++ {
		s = append(s, i)
	}

	return s
}

func BenchmarkList_BSort(b *testing.B) {
	l := MakeBySlice(generateBigSlice())

}
