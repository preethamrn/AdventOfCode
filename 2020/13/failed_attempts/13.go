// https://adventofcode.com/2020/day/13

// I believe this failed due to integer overflow in one of the intemediary steps.
// This attempts (as well as the python port) used a handrolled version of the extended Euclidean algorithm
// I later realized that Go has an inbuilt version of this in the math/big library so I rewrote part 2 code using that.

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	testInput  = "../test.txt"
	testInput2 = "../test2.txt"
	mainInput  = "../main.txt"
)

func part1(t0 int, buses []int) int {
	min := math.MaxInt32
	ret := 0
	for _, bus := range buses {
		diff := bus*int(math.Ceil(float64(t0)/float64(bus))) - t0
		if diff < min {
			min = diff
			ret = diff * bus
		}
	}
	return ret
}

func divmod(a, b int64) (quo int64, rem int64) {
	return a / b, a % b
}

func egcd(a, b int64) (int64, int64, int64) {
	or, r := a, b
	var os, s int64 = 1, 0
	var ot, t int64 = 0, 1

	for r != 0 {
		quo, rem := divmod(or, r)
		or, r = r, rem
		os, s = s, os-quo*s
		ot, t = t, ot-quo*t
	}
	return or, os, ot
}

func alignment(a, b, pa, pb int64) (int64, int64) {
	gcd, s, _ := egcd(a, b)
	diff := pa - pb
	pdMult, pdRem := divmod(diff, gcd)
	if pdRem != 0 {
		panic("phase mismatch")
	}
	period := a / gcd * b
	phase := pa - s*pdMult*(a%period)
	return period, phase
}

func check(period, offset int64, buses, offsets []int, i int) bool {
	for j := range buses[0 : i+1] {
		if period%int64(buses[j]) != 0 {
			fmt.Println(period, offset, buses[j])
			return false
		}
		if (offset-int64(offsets[j]))%int64(buses[j]) != 0 {
			fmt.Println(period, offset, buses[j], offsets[j], j, i)
			return false
		}
	}
	return true
}

func part2(buses, offsets []int) int64 {
	period := int64(buses[0])
	offset := int64(0)
	for i := 1; i < len(buses); i++ {
		var phase int64
		period, phase = alignment(period, int64(buses[i]), offset, int64(offsets[i]))
		offset = (period + phase) % period
		check(period, offset, buses, offsets, i)
	}
	return period - offset
}

func main() {
	bs, _ := ioutil.ReadFile(mainInput)
	s := string(bs)
	lines := strings.Split(s, "\n")

	// t0, _ := strconv.Atoi(lines[0])
	for k := 1; k < len(lines); k++ {
		busesStr := strings.Split(lines[k], ",")
		buses := []int{}
		busOffsets := []int{}
		for i, bus := range busesStr {
			if bus == "x" {
				continue
			}
			t, _ := strconv.Atoi(bus)
			buses = append(buses, t)
			busOffsets = append(busOffsets, i)
		}

		// fmt.Println(part1(t0, buses))
		fmt.Println(part2(buses, busOffsets))
	}
}
