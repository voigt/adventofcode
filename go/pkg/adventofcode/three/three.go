package three

import (
	"fmt"
	"strings"

	"github.com/voigt/adventofcode2021/pkg/utils"
)

// type Position struct {
// 	horizontal int
// 	depth      int
// 	aim        int
// }

func PartOne(inputs []string) int {
	return powerconsumption(inputs)
}

func PartTwo(inputs []string) int {
	return oxygenGeneratorRating(inputs) * co2ScrubberRating(inputs)
}

func powerconsumption(input []string) int {
	ri := rotateInput(input)
	gamma := getGamma(ri)
	epsilon := getEpsilon(ri)

	return gamma * epsilon
}

func rotateInput(input []string) []string {
	res := []string{}

	for k := 0; k < len(input[0]); k++ {
		var b []string
		for i := 0; i < len(input); i++ {
			s := strings.Split(input[i], "")
			b = append(b, s[k])
		}
		res = append(res, strings.Join(b, ""))
	}
	return res
}

func getLeastCommonBitInString(input string) string {
	zero := strings.Count(input, "0")
	one := strings.Count(input, "1")
	if one < zero {
		return "1"
	}
	return "0"
}

func getMostCommonBitInString(input string) string {
	zero := strings.Count(input, "0")
	one := strings.Count(input, "1")
	if one > zero {
		return "1"
	}
	return "0"
}

func getEpsilon(input []string) int {
	s := []string{}

	for _, v := range input {
		s = append(s, getLeastCommonBitInString(v))
	}

	return utils.BinToDec(strings.Join(s, ""))
}

func getGamma(input []string) int {
	s := []string{}

	for _, v := range input {
		s = append(s, getMostCommonBitInString(v))
	}

	return utils.BinToDec(strings.Join(s, ""))
}

func oxygenGeneratorRating(input []string) int {
	resultSet := input
	for k := 0; k < len(input[0]); k++ {
		resultSet = getSetOfIndex1(resultSet, k)
		if len(resultSet) == 1 {
			return utils.BinToDec(resultSet[0])
		}
	}
	fmt.Println(resultSet)
	return utils.BinToDec(resultSet[0])
}

func co2ScrubberRating(input []string) int {
	resultSet := input
	for k := 0; k < len(input[0]); k++ {
		resultSet = getSetOfIndex0(resultSet, k)
		fmt.Println(resultSet)
		if len(resultSet) == 1 {
			return utils.BinToDec(resultSet[0])
		}

	}
	return utils.BinToDec(resultSet[0])
}

func getSetOfIndex1(input []string, i int) []string {
	resultSet1 := []string{}
	resultSet0 := []string{}

	for k := 0; k < len(input); k++ {
		l := strings.Split(input[k], "")
		fmt.Println(l)
		if l[i] == "1" {
			resultSet1 = append(resultSet1, input[k])
		} else {
			resultSet0 = append(resultSet0, input[k])
		}
	}

	if len(resultSet1) >= len(resultSet0) {
		return resultSet1
	} else {
		return resultSet0
	}
}

func getSetOfIndex0(input []string, i int) []string {
	resultSet1 := []string{}
	resultSet0 := []string{}

	for k := 0; k < len(input); k++ {
		l := strings.Split(input[k], "")
		if l[i] == "1" {
			resultSet1 = append(resultSet1, input[k])
		} else {
			resultSet0 = append(resultSet0, input[k])
		}
	}

	if len(resultSet1) < len(resultSet0) {
		return resultSet1
	} else {
		return resultSet0
	}
}
