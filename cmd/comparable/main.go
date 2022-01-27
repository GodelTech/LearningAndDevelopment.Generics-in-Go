// package main contains example of comparable generic type usage
package main

import (
	"fmt"
)

var (
	intArr = []int{1, 2, 3, 4, 5, 6}
	strArr = []string{"first", "second", "apple", "banana"}
)

// comparableSlice is an example of generic feature with new "comparable" type
type comparableSlice[T comparable] []T

// allEqual returns true when all slice elements are equal
// like [1, 1, 1] => true; [2, 1, 4] => false
func allEqual[T comparable](s comparableSlice[T]) bool {
	if len(s) == 0 {
		return true
	}

	last := s[0]
	for _, cur := range s[1:] {
		if cur != last {
			return false
		}
		last = cur
	}
	return true
}

// contains returns true if slice arr contains element v
func contains[T comparable](arr []T, v T) bool {
	for _, e := range arr {
		if e == v {
			return true
		}
	}
	return false
}

func main() {
	fmt.Printf("trying to find %v in given slice, the result is: %v\n", 5,  contains(intArr, 5))

	fmt.Printf("trying to find %v in given slice, the result is: %v\n", "cat",  contains(strArr, "cat"))

	fmt.Printf("trying to find %v in given slice, the result is: %v\n", "apple",  contains(strArr, "apple"))

	fmt.Printf("check all elemens of slice are equal: %v\n", allEqual([]int{2,3,4,5}))

	fmt.Printf("check all elemens of slice are equal: %v\n", allEqual([]string{"banana", "banana"}))

	// this one won't work 'cause it's impossible to compare integers and strings
	//fmt.Printf("trying to find %v in given slice, result is: %v", 1,  contains(strArr, 1))
}