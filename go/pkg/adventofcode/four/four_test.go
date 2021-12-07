package four

import (
	"fmt"
	"testing"
)

func TestHasRowWin(t *testing.T) {
	b := Board{
		State: [][]int{
			{-1, -1, -1, -1, -1},
			// {22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		Winner: false,
	}
	is := b.HasRowWin()
	should := true
	fmt.Printf("is:     %v\n", is)
	fmt.Printf("should: %v\n", should)
	if is != should {
		t.Errorf("IsRowWin")
		t.Fail()
	}
}

func TestHasColWin(t *testing.T) {
	b := Board{
		State: [][]int{
			{22, -1, 17, 11, 0},
			{8, -1, 23, 4, 24},
			{21, -1, 14, 16, 7},
			{6, -1, 3, 18, 5},
			{1, -1, 20, 15, 19},
		},
		Winner: false,
	}
	is := b.HasColWin()
	should := true
	fmt.Printf("is:     %v\n", is)
	fmt.Printf("should: %v\n", should)
	if is != should {
		t.Errorf("IsRowWin")
		t.Fail()
	}
}

// func IsColWin(t *testing.T) {

// }
