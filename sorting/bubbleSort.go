package sorting

import (
	"github.com/patelpreet422/go-algorithms/util"
)

// BubbleSort sorts the slice using bubble sort algorithm
func BubbleSort(a []int, comparator func(int, int) bool) {
	n := len(a)
	sorted := true

	for i := 0; i < n-1; i++ {
		sorted = true

		for j := 0; j < n-1-i; j++ {
			if !comparator(a[j], a[j+1]) {
				util.Swap(a, j, j+1)
				sorted = false
			}
		}

		if sorted {
			return
		}
	}
}
