// https://adventofcode.com/2020/day/9

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const (
	testInput = "../9.1/9.1_test.txt"
	mainInput = "../9.1/9.1.txt"
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

	invalidNum := 0
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
			invalidNum = n
			break
		}
	}

	i, j := 0, 1
	sum := nums[i] + nums[j]
	for i < j && j < len(nums) {
		for sum > invalidNum {
			sum -= nums[i]
			i++
		}
		for sum < invalidNum {
			j++
			sum += nums[j]
		}
		if sum == invalidNum {
			break
		}
	}
	// fmt.Println(i, j)
	rangeNums := nums[i : j+1]

	sort.Slice(rangeNums, func(i, j int) bool {
		return rangeNums[i] < rangeNums[j]
	})
	fmt.Println(rangeNums[0] + rangeNums[len(rangeNums)-1])
}
