// https://adventofcode.com/2020/day/14

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	testInput  = "test.txt"
	testInput2 = "test2.txt"
	mainInput  = "main.txt"
)

var (
	maskRegex = regexp.MustCompile(`mask = (?P<mask>.*)`)
	memRegex  = regexp.MustCompile(`mem\[(?P<addr>\d+)\] = (?P<val>\d+)`)
)

type operation struct {
	mask string
	addr int
	val  int64
}

func maskToIntMasks(mask string) (int64, int64) {
	orMask := int64(0)
	andMask := int64(0xfffffffff)
	for i := len(mask) - 1; i >= 0; i-- {
		shift := len(mask) - i
		if mask[i] == '0' {
			andMask &= (0xfffffffff << (shift)) | (0xfffffffff >> (i + 1))
		} else if mask[i] == '1' {
			orMask |= 1 << (shift - 1)
		}
	}
	return orMask, andMask
}

func part1(ops []operation) int64 {
	memory := map[int]int64{}
	orMask, andMask := maskToIntMasks("")
	for _, op := range ops {
		if op.mask != "" {
			orMask, andMask = maskToIntMasks(op.mask)
			// fmt.Println(strconv.FormatInt(orMask, 2), strconv.FormatInt(andMask, 2))
		} else {
			memory[op.addr] = op.val&andMask | orMask
		}
	}

	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	return sum
}

func maskToMemories(mask string) (int64, []int64) {
	memoryIdx := int64(0)
	notcares := []int{}
	for i, x := range mask {
		switch x {
		case '1':
			memoryIdx |= 1 << (len(mask) - i - 1)
		case 'X':
			notcares = append(notcares, (len(mask) - i - 1))
		}
	}

	memories := []int64{memoryIdx}
	iMask := int64(0xfffffffff)
	for _, idx := range notcares {
		oldLen := len(memories)
		rightShift := (len(mask) - idx)
		for i := 0; i < oldLen; i++ {
			iMask &= (0xfffffffff << (idx + 1)) | (0xfffffffff >> rightShift)
			memories = append(memories, memories[i]|(1<<idx))
		}
	}
	return iMask, memories
}

func part2(ops []operation) int64 {
	memory := map[int64]int64{}
	mask := int64(0xfffffffff)
	memMasks := []int64{}
	for _, op := range ops {
		if op.mask != "" {
			mask, memMasks = maskToMemories(op.mask)
		} else {
			maskedMem := int64(op.addr) & mask
			for _, memMask := range memMasks {
				memory[maskedMem|memMask] = op.val
			}
		}
	}

	sum := int64(0)
	for _, v := range memory {
		sum += v
	}
	return sum
}

func main() {
	bs, _ := ioutil.ReadFile(mainInput)
	s := string(bs)
	lines := strings.Split(s, "\n")

	ops := []operation{}
	for _, line := range lines {
		matches := maskRegex.FindStringSubmatch(line)
		if matches != nil {
			ops = append(ops, operation{mask: matches[1]})
			continue
		}
		matches = memRegex.FindStringSubmatch(line)
		if matches != nil {
			addr, _ := strconv.Atoi(matches[1])
			val, _ := strconv.Atoi(matches[2])
			ops = append(ops, operation{addr: addr, val: int64(val)})
			continue
		} else {
			panic("operation " + line + " doesn't match anything.")
		}
	}

	fmt.Println(part1(ops))
	fmt.Println(part2(ops))
}
