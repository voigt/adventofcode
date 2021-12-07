package seven

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/voigt/adventofcode2021/pkg/utils"
)

func getIntArrayOfString(inputs []string) []int {
	var ints []int
	for _, s := range strings.Split(inputs[0], ",") {
		num, _ := strconv.Atoi(s)
		ints = append(ints, num)
	}
	return ints
}

func PartOne(inputs []string) int {
	crabs := getIntArrayOfString(inputs)
	sort.Ints(crabs)

	max, min := crabs[len(crabs)-1], crabs[0]
	minFuel := math.MaxInt

	for i := min; i <= max; i++ {
		currFuel := 0
		for _, crab := range crabs {
			v := utils.Abs(i - crab)
			currFuel += v
		}
		if currFuel < minFuel {
			minFuel = currFuel
		}
	}

	return minFuel
}

func PartTwo(inputs []string) int {
	crabs := getIntArrayOfString(inputs)
	sort.Ints(crabs)

	fmt.Printf("%v\n", crabs)

	max, min := crabs[len(crabs)-1], crabs[0]
	minFuel := math.MaxInt

	for i := min; i <= max; i++ {
		currFuel := 0
		for _, crab := range crabs {
			v := fuelPerUnit(utils.Abs(i - crab))
			currFuel += v
		}
		if currFuel < minFuel {
			minFuel = currFuel
		}
	}

	return minFuel
}

func fuelPerUnit(n int) int {
	return (n * (1 + n)) / 2
}
