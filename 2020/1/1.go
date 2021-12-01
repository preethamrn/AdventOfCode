// https://adventofcode.com/2020/day/10

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const (
	testInput = "test.txt"
	mainInput = "main.txt"
)

func part1(nums []int, target int) int {
	found := map[int]struct{}{}
	for _, n := range nums {
		if _, ok := found[target-n]; ok {
			return n * (target - n)
		}
		found[n] = struct{}{}
	}
	return -1
}

func part2(nums []int, target int) int {
	for _, n := range nums {
		ret := part1(nums, target-n)
		if ret != -1 {
			return n * ret
		}
	}
	return -1
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

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(part1(nums, 2020))
	fmt.Println(part2(nums, 2020))
}
