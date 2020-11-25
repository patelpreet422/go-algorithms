package util

import "math/rand"

// Swap elements at index i and j
func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// FillRandomInts fills slice with random integers
func FillRandomInts(a []int) {
	for i := 0; i < len(a); i++ {
		a[i] = rand.Intn(100)
	}
}

// IsSorted checks if the slice is sorted or not
func IsSorted(a []int, comparator func(int, int) bool) bool {
	n := len(a)
	if n <= 1 {
		return true
	}

	for i := 1; i < n; i++ {
		if !comparator(a[i-1], a[i]) {
			return false
		}
	}

	return true
}

// LessEq is <= comparator for comparing integers
func LessEq(a, b int) bool {
	return a <= b
}

// Less is < comparator for comparing integers
func Less(a, b int) bool {
	return a < b
}

// GreaterEq is >= comparator for comparing integers
func GreaterEq(a, b int) bool {
	return a >= b
}

// Greater is > comparator for comparing integers
func Greater(a, b int) bool {
	return a > b
}

// Eq is == comparator for comparing integers
func Eq(a, b int) bool {
	return a == b
}

// NotEq is != comparator for comparing integers
func NotEq(a, b int) bool {
	return a != b
}
