package one

import (
	"fmt"
	"testing"
)

func TestGetIncreaseCount(t *testing.T) {
	is := GetIncreaseCount([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	should := 9
	fmt.Printf("is:     %d\n", is)
	fmt.Printf("should: %d\n", should)
	if is != should {
		t.Errorf("GetIncreaseCount was incorrect")
		t.Fail()
	}
}

func TestGetWindowSliceSums(t *testing.T) {
	is := GetWindowSliceSums([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	should := []int{6, 9, 12, 15, 18, 21, 24, 27, 19, 10}
	fmt.Printf("is:     %v\n", is)
	fmt.Printf("should: %v\n", should)
	if EqualSlice(is, should) == false {
		// t.Errorf("getWindowSliceSums was incorrect, got: %v, want: %v.", is, should)
		t.Errorf("getWindowSliceSums was incorrect")
		t.Fail()
	}
}

func EqualSlice(a, b []int) bool {
	if len(a) != len(b) {
		// fmt.Printf("not equal length\n")
		return false
	}
	for i, v := range a {
		if v != b[i] {
			// fmt.Printf("not equal value\n")
			return false
		}
	}
	// fmt.Printf("everything equal\n")
	return true
}
