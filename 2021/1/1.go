// https://adventofcode.com/2021/day/1

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

func part1(nums []int) int {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			count++
		}
	}
	return count
}

func part2(nums []int) int {
	count := 0
	prev := nums[0] + nums[1] + nums[2]
	for i := 3; i < len(nums); i++ {
		curr := prev + nums[i] - nums[i-3]
		if curr > prev {
			count++
		}
	}
	return count
}

func main() {
	bs, _ := ioutil.ReadFile(mainInput)
	s := string(bs)
	lines := strings.Split(s, "\n")

	nums := []int{}
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	fmt.Println(part1(nums))
	fmt.Println(part2(nums))
}
