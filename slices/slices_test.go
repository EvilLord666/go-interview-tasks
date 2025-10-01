package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// This is a set of interview mini tasks related 2 go slices
// Go sources - https://go.dev/src/runtime/slice.go

// TestSliceBasedOnSlicesWithoutReAlloc test
func TestSliceBasedOnSlicesWithoutReAlloc(t *testing.T) {
	// task to check what would be in a slice1 if slice2 based on slice1 item
	// value was changed
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice2 := slice1[2:5]
	slice2[1] = 6
	assert.Equal(t, slice1, []int{1, 2, 3, 6, 5, 6, 7, 8, 9, 0})
	assert.Equal(t, slice2, []int{3, 6, 5})
}

func TestSliceBasedOnSlicesWithReAlloc(t *testing.T) {
	// task to check what would be in a slice1 if slice2 based on slice1 item
	// slice1 has no capacity, we are appending a new value that is leading to
	// a new memory allocation for slice1, therefore slice1 and slice2 points to different
	// memory addresses
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice2 := slice1[2:5]
	slice1 = append(slice1, 10)
	slice2[0] = 9
	assert.Equal(t, slice1, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 10})
	assert.Equal(t, slice2, []int{9, 4, 5})
}

func TestSliceBasedOnSmallerSliceWithSmallerModification(t *testing.T) {
	//
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 0, 10)
	slice2 = append(slice2, slice1...)
	slice1[1] = 9
	assert.Equal(t, slice2, []int{1, 2, 3})
}
