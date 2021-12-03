package utils

import (
	"fmt"
	"testing"
)

func TestBinToDec(t *testing.T) {
	is := BinToDec("100110101101")
	should := 2477
	fmt.Printf("is:     %d\n", is)
	fmt.Printf("should: %d\n", should)
	if is != should {
		t.Errorf("tBinToDec")
		t.Fail()
	}
}
