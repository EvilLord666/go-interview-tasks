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
