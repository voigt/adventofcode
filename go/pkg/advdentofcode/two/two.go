package two

import (
	"strconv"
	"strings"

	"github.com/voigt/adventofcode2021/pkg/utils"
)

type Position struct {
	horizontal int
	depth      int
	aim        int
}

func PartOne() int {
	inputs := utils.GetFileInputsAsStrSlice("pkg/advdentofcode/two/input.txt")
	p := Position{}

	for _, v := range inputs {
		f := strings.Split(v, " ")
		s, _ := strconv.Atoi(f[1])
		switch f[0] {
		case "down":
			p.depth += s
		case "up":
			p.depth -= s
		case "forward":
			p.horizontal += s
		}
	}

	return p.horizontal * p.depth
}

func PartTwo() int {
	inputs := utils.GetFileInputsAsStrSlice("pkg/advdentofcode/two/input.txt")
	p := Position{}

	for _, v := range inputs {
		f := strings.Split(v, " ")
		s, _ := strconv.Atoi(f[1])
		switch f[0] {
		case "down":
			// p.depth += s
			p.aim += s
		case "up":
			// p.depth -= s
			p.aim -= s
		case "forward":
			p.horizontal += s
			p.depth += p.aim * s
		}
	}

	return p.horizontal * p.depth
}
