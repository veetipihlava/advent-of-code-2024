package main

import (
	"bufio"
	"errors"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	first, second, err := parseInput("./input")
	if err != nil {
		println(err.Error())
		return
	}

	result, err := solve(first, second)
	if err != nil {
		println(err.Error())
		return
	}

	println(result)
}

func parseInput(path string) ([]int, []int, error) {
	var left []int
	var right []int

	file, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		pair := strings.Split(text, "   ")
		if len(pair) != 2 {
			return nil, nil, errors.New("invalid file")
		}
		integer, err := strconv.Atoi(pair[0])
		if err != nil {
			return nil, nil, err
		}
		left = append(left, integer)

		integer, err = strconv.Atoi(pair[1])
		if err != nil {
			return nil, nil, err
		}
		right = append(right, integer)
	}
	return left, right, nil
}

func abs(integer int) int {
	if integer < 0 {
		return -integer
	}

	return integer
}

func solve(first []int, second []int) (int, error) {
	if len(first) != len(second) {
		return 0, errors.New("the length of the input should be the same")
	}

	slices.Sort(first)
	slices.Sort(second)

	distanceSum := 0
	for i := 0; i < len(first); i++ {
		distance := first[i] - second[i]
		distanceSum += abs(distance)
	}

	return distanceSum, nil
}
