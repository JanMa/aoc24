package main

import (
	"fmt"
	"math"
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
	input, err := parseInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	if partOne(sample) != 11 {
		panic("sample test failed")
	}

	fmt.Println("Part one:")
	fmt.Println(partOne(input))

	if partTwo(sample) != 31 {
		fmt.Println(partTwo(sample))
		panic("sample test failed")
	}

	fmt.Println("Part two:")
	fmt.Println(partTwo(input))
}

func parseInput(file string) (result [][]int, err error) {
	result = make([][]int, 2)
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	for _, line := range lines {
		l := strings.Split(line, "   ")
		r0, err := strconv.Atoi(l[0])
		if err != nil {
			return nil, err
		}
		r1, err := strconv.Atoi(l[1])
		if err != nil {
			return nil, err
		}
		result[0] = append(result[0], r0)
		result[1] = append(result[1], r1)
	}

	return
}

func partOne(input [][]int) (result int) {
	l := input[0]
	r := input[1]
	slices.Sort(l)
	slices.Sort(r)

	for i := 0; i < len(l); i++ {
		result += int(math.Abs(float64(l[i] - r[i])))
	}
	return
}

func partTwo(input [][]int) (result int) {
	l := input[0]
	r := input[1]

	for _, i := range l {
		if slices.Contains(r, i) {
			var c int
			for _, j := range r {
				if i == j {
					c++
				}
			}
			result += i * c
		}

	}
	return
}
