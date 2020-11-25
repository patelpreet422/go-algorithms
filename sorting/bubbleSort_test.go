package sorting

import (
	"testing"

	"github.com/patelpreet422/go-algorithms/util"
)

func TestBubbleSort(t *testing.T) {
	t.Parallel()

	a := make([]int, 10)
	util.FillRandomInts(a)

	BubbleSort(a, util.LessEq)
	isSorted := util.IsSorted(a, util.LessEq)

	if !isSorted {
		t.Errorf("Array is not sorted")
	}
}
