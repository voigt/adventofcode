package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetFileInputsAsIntSlice(f string) []int {
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

func GetFileInputsAsStrSlice(f string) []string {
	var res []string
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("unable to read input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		n := scanner.Text()
		res = append(res, n)
	}
	return res
}

func BinToDec(binary string) int {
	var result int
	for i := 0; i < len(binary); i++ {
		if binary[i] == '1' {
			result += 1 << uint(len(binary)-i-1)
		}
	}
	return result
}
