package main

import (
	"regexp"
	"strconv"
	"strings"
)

func d3p1(in []byte) (int, error) {
	total := 0
	for _, cmd := range regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`).FindAll(in, -1) {
		cmdStr := string(cmd)
		cmdStr = strings.Replace(cmdStr, "mul(", "", 1)
		cmdStr = strings.Replace(cmdStr, ")", "", 1)
		xy := strings.Split(cmdStr, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		total += (x * y)
	}
	return total, nil
}

func d3p2(in []byte) (int, error) {
	enabled := true
	total := 0
	for _, cmd := range regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)|do\(\)|don't\(\)`).FindAll(in, -1) {
		cmdStr := string(cmd)
		if cmdStr == "do()" {
			enabled = true
		} else if cmdStr == "don't()" {
			enabled = false
		} else if enabled {
			cmdStr = strings.Replace(cmdStr, "mul(", "", 1)
			cmdStr = strings.Replace(cmdStr, ")", "", 1)
			xy := strings.Split(cmdStr, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			total += (x * y)
		}
	}
	return total, nil
}
