// https://adventofcode.com/2020/day/8

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	testInput = "9.1_test.txt"
	mainInput = "9.1.txt"
)

func main() {
	bs, _ := ioutil.ReadFile(mainInput)
	s := string(bs)
	lines := strings.Split(s, "\n")

	nums := []int{}
	preambleLen, _ := strconv.Atoi(lines[0])
	for _, line := range lines[1:len(lines)] {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	// fmt.Println(nums)

	lastNums := map[int]int{}
	curr := 0
	for curr < preambleLen {
		lastNums[nums[curr]]++
		curr++
	}
	// fmt.Println(lastNums)

	for i, n := range nums[preambleLen:len(nums)] {
		valid := false
		for j := i; j < i+preambleLen; j++ {
			if lastNums[n-nums[j]] > 0 || (lastNums[nums[j]] >= 2 && n == nums[j]*2) {
				valid = true
				break
			}
		}
		lastNums[n]++
		lastNums[nums[i]]--

		if !valid {
			fmt.Println(n)
			break
		}
	}
}
