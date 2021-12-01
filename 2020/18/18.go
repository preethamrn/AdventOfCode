// https://adventofcode.com/2020/day/14

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	testInput = "test.txt"
	mainInput = "main.txt"
)

var precedence map[rune]int{
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

func eval(line string) (int64, int) {
	result := 0

	stack := []string{}
	for i := 0; i < len(line); i++ {
		c := line[i]
		switch c {
		case '+': fallthrough
		case '-': fallthrough
		case '*': fallthrough
		case '/': fallthrough
		case '(':
			val, inc := eval(line[i+1:])
			i += inc + 1
			stack = append(stack, fmt.Sprintf("%d", val))
		case ')':
			return result, i
		default:
			
		}
	}

	return int64(result), len(line)
}

func part1(lines []string) int64 {
	sum := int64(0)
	for _, line := range lines {
		sum += eval(line)
	}
	return sum
}

func part2(lines []string) int64 {
}

func main() {
	bs, _ := ioutil.ReadFile(mainInput)
	s := string(bs)
	lines := strings.Split(s, "\n")

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}
