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
