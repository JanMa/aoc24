package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	sample, err := parseInput("sample.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sample2, err := parseInput("sample2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	if partOneTwo(sample, false) != 2 {
		panic("sample test failed")
	}

	input, err := parseInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Part one:")
	fmt.Println(partOneTwo(input, false))

	if partOneTwo(sample, true) != 4 {
		panic("sample test failed")
	}

	if partOneTwo(sample2, true) != 12 {
		panic("sample test failed")
	}

	fmt.Println("Part two:")
	fmt.Println(partOneTwo(input, true))

}

func parseInput(file string) (result [][]int, err error) {
	result = [][]int{}
	f, err := os.ReadFile(file)
	if err != nil {
		return
	}

	for _, line := range strings.Split(strings.TrimSpace(string(f)), "\n") {
		l := strings.Split(line, " ")
		row := []int{}
		for _, n := range l {
			r, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			row = append(row, r)
		}
		result = append(result, row)
	}
	return
}

func partOneTwo(input [][]int, dampen bool) (result int) {
	for _, row := range input {
		s := isSafe(row, dampen)
		if s {
			result++
		}
	}
	return result
}

func isSafe(row []int, dampen bool) (result bool) {
	result = true
	direction := row[1] - row[0]
	switch {
	case direction > 0: // ascending
		for i := 0; i < len(row)-1; i++ {
			if row[i] > row[i+1] || (row[i+1]-row[i] > 3 || row[i] == row[i+1]) {
				result = false
			}
		}
	case direction < 0: // descending
		for i := 0; i < len(row)-1; i++ {
			if row[i] < row[i+1] || (row[i]-row[i+1] > 3) || row[i] == row[i+1] {
				result = false
			}
		}
	default:
		result = false
	}

	if dampen {
		if !result {
			for i := 0; i < len(row); i++ {
				r := slices.Clone(row)
				tmp := isSafe(slices.Delete(r, i, i+1), false)
				if tmp {
					result = tmp
					return
				}
			}
		}
	}
	return
}
