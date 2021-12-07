package main

import (
	"fmt"

	"github.com/voigt/adventofcode2021/pkg/adventofcode/four"
	"github.com/voigt/adventofcode2021/pkg/adventofcode/one"
	seven "github.com/voigt/adventofcode2021/pkg/adventofcode/seven"
	"github.com/voigt/adventofcode2021/pkg/adventofcode/three"
	"github.com/voigt/adventofcode2021/pkg/adventofcode/two"
	"github.com/voigt/adventofcode2021/pkg/utils"
)

func main() {
	inputDay1 := utils.GetFileInputsAsIntSlice("pkg/adventofcode/one/input.txt")
	fmt.Printf("Day 1 - Result PartOne: %d\n", one.PartOne(inputDay1))
	fmt.Printf("Day 1 - Result PartTwo: %d\n", one.PartTwo(inputDay1))
	fmt.Println("------------")
	inputDay2 := utils.GetFileInputsAsStrSlice("pkg/adventofcode/two/input.txt")
	fmt.Printf("Day 2 - Part One: %d\n", two.PartOne(inputDay2))
	fmt.Printf("Day 2 - Part Two: %d\n", two.PartTwo(inputDay2))
	fmt.Println("------------")
	inputDay3 := utils.GetFileInputsAsStrSlice("pkg/adventofcode/three/input.txt")
	fmt.Printf("Day 3 - Part One: %d\n", three.PartOne(inputDay3))
	fmt.Printf("Day 3 - Part Two: %d\n", three.PartTwo(inputDay3))
	fmt.Println("------------")
	inputDay4 := utils.GetFileInputsAsStrSlice("pkg/adventofcode/four/input.txt")
	fmt.Printf("Day 4 - Part One: %d\n", four.PartOne(inputDay4))
	fmt.Printf("Day 4 - Part Two: %d\n", four.PartTwo(inputDay4))
	fmt.Println("------------")
	inputDay7 := utils.GetFileInputsAsStrSlice("pkg/adventofcode/seven/input.txt")
	fmt.Printf("Day 7 - Part One: %d\n", seven.PartOne(inputDay7))
	fmt.Printf("Day 7 - Part Two: %d\n", seven.PartTwo(inputDay7))
}
