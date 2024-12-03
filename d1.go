package main

import (
	"slices"
	"strconv"
	"strings"
)

func d1p1(in []byte) (int, error) {
	left := []int{}
	right := []int{}
	for _, line := range strings.Split(string(in), "\n") {
		parts := strings.Split(line, " ")
		l, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		r, err := strconv.Atoi(parts[3])
		if err != nil {
			return 0, err
		}
		left = append(left, l)
		right = append(right, r)
	}
	slices.Sort(left)
	slices.Sort(right)
	total := 0
	for i := 0; i < len(left); i++ {
		if left[i] > right[i] {
			total += left[i] - right[i]
		} else {
			total += right[i] - left[i]
		}
	}
	return total, nil
}

func d1p2(in []byte) (int, error) {
	existsInLeft := make(map[int]bool)
	timesInRight := make(map[int]int)
	for _, line := range strings.Split(string(in), "\n") {
		parts := strings.Split(line, " ")
		l, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		existsInLeft[l] = true
		r, err := strconv.Atoi(parts[3])
		if err != nil {
			return 0, err
		}
		timesInRight[r] = timesInRight[r] + 1
	}
	total := 0
	for i, n := range timesInRight {
		if existsInLeft[i] {
			total += (i * n)
		}
	}
	return total, nil
}
