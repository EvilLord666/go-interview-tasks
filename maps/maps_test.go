package maps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapIntersection(t *testing.T) {
	// Test task to find intersection between slices

	s1 := []int{10, 20, 30, 40}
	s2 := []int{20, 30}

	intersection := findIntersection(s1, s2)
	assert.Equal(t, intersection, []int{20, 30})

	s1 = []int{2, 3, 2, 2, 4}
	s2 = []int{3, 2, 2}

	intersection = findIntersection(s1, s2)
	assert.Equal(t, intersection, []int{3, 2, 2})

	s1 = []int{2, 3, 2, 4}
	s2 = []int{2, 2, 2, 2, 4}

	intersection = findIntersection(s1, s2)
	assert.Equal(t, intersection, []int{2, 2, 4})
}

func findIntersection(s1 []int, s2 []int) []int {
	var result []int
	// int - key of slice item, map value is a count of intersected values
	searchMap := make(map[int]int)
	// 1. range s1, fill map
	for _, v := range s1 {
		c, ok := searchMap[v]
		if ok {
			searchMap[v] = c + 1
		} else {
			searchMap[v] = 1
		}
	}
	// 2. range s2, decrease values if found and while if value is non-zero
	// save to the result
	for _, v := range s2 {
		c, ok := searchMap[v]
		if ok {
			if c > 0 {
				result = append(result, v)
				searchMap[v] = c - 1
			}
		}
	}

	return result
}
