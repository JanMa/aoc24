package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sample, err := parseInput("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	if partOne(sample) != 161 {
		panic("sample test failed")
	}

	input, err := parseInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Part one:")
	fmt.Println(partOne(input))

	sample2, err := parseInput("sample2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	if partTwo(sample2) != 48 {
		panic("sample test failed")
	}

	fmt.Println("Part two:")
	fmt.Println(partTwo(input))

}

func parseInput(file string) (result string, err error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return
	}

	result = strings.TrimSpace(string(f))
	return
}

func sanitize(input string) []string {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	return re.FindAllString(input, -1)
}

func calculate(input []string) (result int) {
	for _, i := range input {
		s := strings.Split(strings.TrimSuffix(strings.TrimPrefix(i, "mul("), ")"), ",")
		r, _ := strconv.Atoi(s[0])
		for _, v := range s[1:] {
			o, _ := strconv.Atoi(v)
			r *= o
		}
		result += r
	}
	return
}

func partOne(input string) (result int) {
	return calculate(sanitize(input))
}

func conditional(input string) string {
	re := regexp.MustCompile(`(?sU)don't\(\).*do\(\)`)

	return re.ReplaceAllString(input, "")
}

func partTwo(input string) (result int) {
	return calculate(sanitize(conditional(input)))
}
