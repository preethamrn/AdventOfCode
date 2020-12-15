// https://adventofcode.com/2020/day/15

package main

import (
	"fmt"
)

var (
	testInput1 = []int{0, 3, 6}
	testInput2 = []int{1, 3, 2}
	testInput3 = []int{2, 1, 3}
	testInput4 = []int{1, 2, 3}
	testInput5 = []int{2, 3, 1}
	testInput6 = []int{3, 2, 1}
	testInput7 = []int{3, 1, 2}
	mainInput  = []int{8, 11, 0, 19, 1, 2}
)

func part(nums []int, final int) int {
	pos := map[int]int{}
	nextNum := 0

	for i, n := range nums {
		if idx, ok := pos[n]; !ok {
			nextNum = 0
		} else {
			nextNum = i - idx
		}
		pos[n] = i
	}
	for i := len(nums); i < final-1; i++ {
		if idx, ok := pos[nextNum]; !ok {
			pos[nextNum] = i
			nextNum = 0
		} else {
			pos[nextNum] = i
			nextNum = i - idx
		}
	}
	return nextNum
}

func main() {
	nums := mainInput

	fmt.Println(part(nums, 2020))
	fmt.Println(part(nums, 30000000))
}
