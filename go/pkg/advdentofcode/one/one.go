package one

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/voigt/adventofcode2021/pkg/utils"
)

func PartOne() int {
	INPUT := utils.GetFileInputsAsIntSlice("pkg/advdentofcode/one/input.txt")

	return GetIncreaseCount(INPUT)
}

func PartTwo() int {
	INPUT := utils.GetFileInputsAsIntSlice("pkg/advdentofcode/one/input.txt")
	return GetIncreaseCount(GetWindowSliceSums(INPUT))
}

func GetWindowSliceSums(w []int) []int {
	window := 0
	var res []int
	l := len(w) - 1
	for i, v := range w {
		if i <= l {
			window += v
		}
		if i+1 <= l {
			window += w[i+1]
		}
		if i+2 <= l {
			window += w[i+2]
		}
		res = append(res, window)
		window = 0
	}
	return res
}

func GetIncreaseCount(s []int) int {
	c := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] < s[i+1] {
			c++
		}
	}
	return c
}

func getFileInputsAsIntSlice(f string) []int {
	var res []int
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("unable to read input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		res = append(res, n)
	}
	return res
}
