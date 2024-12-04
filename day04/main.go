package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	sample, err := parseInput("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println(strings.Join(sample, "\n"))
	// fmt.Println(strings.Join(transpose(sample), "\n"))
	// fmt.Println(strings.Join(diags(sample), "\n"))
	// fmt.Println(strings.Join(diagsRL(sample), "\n"))

	if partOne(sample) != 18 {
		fmt.Println(partOne(sample))
		panic("sample test failed")
	}

	input, err := parseInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Part one:")
	fmt.Println(partOne(input))
}

func parseInput(file string) (result []string, err error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return
	}
	result = strings.Split(strings.TrimSpace(string(f)), "\n")
	return
}

func transpose(input []string) []string {
	result := make([]string, len(input[0]))
	for i := range result {
		for j := range input {
			result[i] += string(input[j][i])
		}
	}
	return result
}

// diags returns all diagonals of the input matrix.
func diags(input []string) []string {
	result := []string{}
	for i := 0; i < len(input); i++ {
		result = append(result, diag(input, i, 0))
	}
	for i := 1; i < len(input[0]); i++ {
		result = append(result, diag(input, 0, i))
	}
	return result
}

// diagsRL returns all diagonals of the input matrix.
// going from right to left.
func diagsRL(input []string) []string {
	result := []string{}
	for i := 0; i < len(input); i++ {
		result = append(result, diagRL(input, i, 0))
	}
	for i := 1; i < len(input[0]); i++ {
		result = append(result, diagRL(input, len(input[0])-1, i))
	}
	return result
}

// diag returns the diagonal starting at (x, y) of the input matrix.
func diag(input []string, x, y int) string {
	result := ""
	for x < len(input[0]) && y < len(input) {
		result += string(input[y][x])
		x++
		y++
	}
	return result
}

// diagRL returns the diagonal starting at (x, y) of the input matrix.
// going from right to left.
func diagRL(input []string, x, y int) string {
	result := ""
	for x >= 0 && y < len(input) {
		result += string(input[y][x])
		x--
		y++
	}
	return result
}

func partOne(input []string) (result int) {
	re := regexp.MustCompile(`XMAS`)
	er := regexp.MustCompile(`SAMX`)
	for _, line := range input {
		result += len(re.FindAllString(line, -1))
		result += len(er.FindAllString(line, -1))
	}
	for _, line := range transpose(input) {
		result += len(re.FindAllString(line, -1))
		result += len(er.FindAllString(line, -1))
	}
	for _, line := range diags(input) {
		result += len(re.FindAllString(line, -1))
		result += len(er.FindAllString(line, -1))
	}
	for _, line := range diagsRL(input) {
		result += len(re.FindAllString(line, -1))
		result += len(er.FindAllString(line, -1))
	}

	return
}
