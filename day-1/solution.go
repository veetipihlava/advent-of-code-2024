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
	left, right, err := parseInput("./input")
	if err != nil {
		println(err.Error())
		return
	}

	println("part1: ")
	result, err := distance(left, right)
	if err != nil {
		println(err.Error())
		return
	}
	println(result)

	println("part2: ")
	result = similarity(left, right)
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

func distance(left []int, right []int) (int, error) {
	if len(left) != len(right) {
		return 0, errors.New("the length of the input should be the same")
	}

	slices.Sort(left)
	slices.Sort(right)

	distanceSum := 0
	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]
		distanceSum += abs(distance)
	}

	return distanceSum, nil
}

func similarity(left []int, right []int) int {
	occurrences := make(map[int]int)

	for i := 0; i < len(right); i++ {
		id := right[i]
		occurrences[id]++
	}

	similarityScore := 0
	for i := 0; i < len(left); i++ {
		id := left[i]
		score := id * occurrences[id]
		similarityScore += score
	}

	return similarityScore
}
