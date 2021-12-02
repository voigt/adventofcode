package main

import (
	"fmt"

	"github.com/voigt/adventofcode2021/pkg/advdentofcode/one"
	"github.com/voigt/adventofcode2021/pkg/advdentofcode/two"
	"github.com/voigt/adventofcode2021/pkg/utils"
)

func main() {
	inputDay1 := utils.GetFileInputsAsIntSlice("pkg/advdentofcode/one/input.txt")
	fmt.Printf("Day 1 - Result PartOne: %d\n", one.PartOne(inputDay1))
	fmt.Printf("Day 1 - Result PartTwo: %d\n", one.PartTwo(inputDay1))
	fmt.Println("------------")
	inputDay2 := utils.GetFileInputsAsStrSlice("pkg/advdentofcode/two/input.txt")
	fmt.Printf("Day 2 - Part One: %d\n", two.PartTwo(inputDay2))
	fmt.Printf("Day 2 - Part Two: %d\n", two.PartTwo(inputDay2))
}
