package two

import (
	"strconv"
	"strings"
)

type Position struct {
	horizontal int
	depth      int
	aim        int
}

func PartOne(inputs []string) int {
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

func PartTwo(inputs []string) int {
	p := Position{}

	for _, v := range inputs {
		f := strings.Split(v, " ")
		s, _ := strconv.Atoi(f[1])
		switch f[0] {
		case "down":
			p.aim += s
		case "up":
			p.aim -= s
		case "forward":
			p.horizontal += s
			p.depth += p.aim * s
		}
	}

	return p.horizontal * p.depth
}
