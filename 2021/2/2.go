// https://adventofcode.com/2021/day/2

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	testInput = "test.txt"
	mainInput = "main.txt"
)

func part1(nums [][]int) int {
	x, y := 0, 0
	for _, v := range nums {
		if v[0] == 0 {
			x += v[1]
		}
		y += v[0] * v[1]
	}
	return x * y
}

func part2(nums [][]int) int {
	x, y, aim := 0, 0, 0
	for _, v := range nums {
		if v[0] == 0 {
			x += v[1]
			y += v[1] * aim
		}
		aim += v[0] * v[1]
	}
	return x * y
}

func main() {
	bs, _ := ioutil.ReadFile(testInput)
	s := string(bs)
	lines := strings.Split(s, "\n")

	nums := [][]int{}
	for _, line := range lines {
		cmd := strings.Split(line, " ")
		c := 0
		switch cmd[0] {
		case "forward":
			c = 0
		case "down":
			c = 1
		case "up":
			c = -1
		}
		num, _ := strconv.Atoi(cmd[1])
		nums = append(nums, []int{c, num})
	}

	fmt.Println(part1(nums))
	fmt.Println(part2(nums))
}
