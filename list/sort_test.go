package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_QSort(t *testing.T) {
	list := Create(3, 2, 1, 4, 6, 5)

	list = list.QSort()

	slice := []int{1, 2, 3, 4, 5, 6}

	assert.Equal(t, slice, list.GetSlice())
}

func TestList_BSort(t *testing.T) {
	list := Create(3, 2, 1, 4, 6, 5)

	list = list.BSort()

	slice := []int{1, 2, 3, 4, 5, 6}

	assert.Equal(t, slice, list.GetSlice())
}

func generateBigSlice() []int {
	var s []int

	for i := 10000; i > 0; i-- {
		s = append(s, i)
	}

	return s
}

func BenchmarkList_QSort(b *testing.B) {
	l := CreateBySlice(generateBigSlice())
	b.ResetTimer()

	l.QSort()
}

func BenchmarkList_BSort(b *testing.B) {
	l := CreateBySlice(generateBigSlice())
	b.ResetTimer()

	l.BSort()
}
