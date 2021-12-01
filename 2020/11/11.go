// https://adventofcode.com/2020/day/11

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	testInput  = "test.txt"
	testInput2 = "test2.txt"
	mainInput  = "main.txt"
)

type seat int

const (
	floor seat = iota
	empty
	occupied
)

func eq(old [][]seat, new [][]seat) bool {
	if len(old) != len(new) && len(old[0]) != len(new[0]) {
		return false
	}
	for i := range old {
		for j := range old[i] {
			if old[i][j] != new[i][j] {
				return false
			}
		}
	}
	return true
}

func copy(old [][]seat) (new [][]seat) {
	for i := range old {
		newline := []seat{}
		for j := range old[i] {
			newline = append(newline, old[i][j])
		}
		new = append(new, newline)
	}
	return
}

func countAdj1(seats [][]seat, x, y int) int {
	cnt := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i < 0 || j < 0 || i >= len(seats) || j >= len(seats[i]) || (i == x && j == y) {
				continue
			}
			if seats[i][j] == occupied {
				cnt++
			}
		}
	}
	return cnt
}

func countAdj2(seats [][]seat, x, y int) int {
	cnt := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			k := x + i
			l := y + j
			for true {
				if k < 0 || l < 0 || k >= len(seats) || l >= len(seats[k]) {
					break
				}
				if seats[k][l] == empty {
					break
				}
				if seats[k][l] == occupied {
					cnt++
					break
				}
				k += i
				l += j
			}
		}
	}
	return cnt
}

func countOccupied(seats [][]seat) (cnt int) {
	for i := range seats {
		for j := range seats[i] {
			if seats[i][j] == occupied {
				cnt++
			}
		}
	}
	return cnt
}

func part1(seats [][]seat) int {
	return solver(seats, 4, countAdj1)
}

func part2(seats [][]seat) int {
	return solver(seats, 5, countAdj2)
}

func solver(seats [][]seat, occLimit int, countAdj func([][]seat, int, int) int) int {
	old := copy(seats)
	for true {
		for i, line := range old {
			for j, c := range line {
				switch c {
				case empty:
					if countAdj(old, i, j) == 0 {
						seats[i][j] = occupied
					}
				case occupied:
					if countAdj(old, i, j) >= occLimit {
						seats[i][j] = empty
					}
				}
			}
		}
		if eq(seats, old) {
			return countOccupied(seats)
		}
		old = copy(seats)
	}
	return -1
}

func main() {
	bs, _ := ioutil.ReadFile(mainInput)
	s := string(bs)
	lines := strings.Split(s, "\n")

	seats := [][]seat{}
	for _, line := range lines {
		squares := []seat{}
		for _, c := range line {
			switch c {
			case 'L':
				squares = append(squares, empty)
			case '.':
				squares = append(squares, floor)
			case '#':
				squares = append(squares, occupied)
			}
		}
		seats = append(seats, squares)
	}
	seats1 := seats
	seats2 := copy(seats)

	fmt.Println(part1(seats1))
	fmt.Println(part2(seats2))
}
