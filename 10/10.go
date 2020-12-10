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
	testInput  = "test.txt"
	testInput2 = "test2.txt"
	mainInput  = "main.txt"
)

func part1(nums []int) {
	diffs := []int{0, 0, 0}
	prev := 0
	for _, n := range nums {
		diffs[n-prev-1]++
		prev = n
	}
	diffs[2]++ // account for the last adapter being able to go up another 3 jolts.
	fmt.Println(diffs)
	fmt.Println(diffs[0] * diffs[2])
}

func part2(nums []int) {
	ways := map[int]int{}
	for _, n := range nums {
		ways[n] = 0
	}
	ways[nums[len(nums)-1]] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		n := nums[i]
		ways[n] = ways[n+1] + ways[n+2] + ways[n+3]
	}
	fmt.Println(ways[1] + ways[2] + ways[3])
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
	fmt.Println(nums)

	part1(nums)
	part2(nums)
}
