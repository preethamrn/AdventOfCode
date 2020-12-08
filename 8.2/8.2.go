// https://adventofcode.com/2020/day/8

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	testInput = "../8.1/8.1_test.txt"
	mainInput = "../8.1/8.1.txt"
)

var (
	regex = regexp.MustCompile(`(?P<cmd>nop|acc|jmp) (?P<cnt>[+-]\d+)`)
)

type Operation int

const (
	NOP Operation = iota
	ACC Operation = iota
	JMP Operation = iota
)

type command struct {
	op  Operation
	num int
}

func compute(cmds []command) (int, bool) {
	acc := 0
	visited := map[int]struct{}{}
	curr := 0
	for true {
		visited[curr] = struct{}{}
		// fmt.Println(curr, cmds[curr], acc)
		switch cmds[curr].op {
		case NOP:
			curr++
		case ACC:
			acc += cmds[curr].num
			curr++
		case JMP:
			curr += cmds[curr].num
		}
		_, ok := visited[curr]
		if ok || curr < 0 || curr > len(cmds) {
			return 0, false
		}
		if curr == len(cmds) {
			return acc, true
		}
	}

	return acc, true
}

func main() {
	cmds := []command{}

	bs, _ := ioutil.ReadFile(mainInput)
	s := string(bs)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)
		cmd := command{}
		switch matches[1] {
		case "nop":
			cmd.op = NOP
		case "acc":
			cmd.op = ACC
		case "jmp":
			cmd.op = JMP
		}
		cmd.num, _ = strconv.Atoi(matches[2])
		cmds = append(cmds, cmd)
	}
	// fmt.Println(cmds)

	for i := range cmds {
		originalCmd := cmds[i]
		if cmds[i].op == NOP {
			cmds[i].op = JMP
		} else if cmds[i].op == JMP {
			cmds[i].op = NOP
		}
		acc, success := compute(cmds)
		if success {
			fmt.Println(acc)
			break
		}
		cmds[i] = originalCmd
	}
}
