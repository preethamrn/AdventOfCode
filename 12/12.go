// https://adventofcode.com/2020/day/12

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

const (
	testInput = "test.txt"
	mainInput = "main.txt"
)

var (
	regex = regexp.MustCompile(`(?P<action>[NSEWLRF])(?P<cnt>\d*)`)
)

type action struct {
	move string
	cnt  int
}

func part1(actions []action) int {
	x, y, d := 0, 0, []float64{1, 0}
	for _, a := range actions {
		cnt := a.cnt
		switch a.move {
		case "N":
			y += cnt
		case "S":
			y -= cnt
		case "E":
			x += cnt
		case "W":
			x -= cnt
		case "F":
			x += int(d[0]) * cnt
			y += int(d[1]) * cnt
		case "R":
			cnt = -cnt
			fallthrough
		case "L":
			theta := float64(cnt) * math.Pi / 180
			tmp := []float64{0, 0}
			tmp[0] = d[0]*math.Cos(theta) - d[1]*math.Sin(theta)
			tmp[1] = d[0]*math.Sin(theta) + d[1]*math.Cos(theta)
			d = tmp
		}
	}

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func part2(actions []action) int {
	var x, y float64 = 10, 1
	var shipX, shipY float64
	for _, a := range actions {
		cnt := float64(a.cnt)
		switch a.move {
		case "N":
			y += cnt
		case "S":
			y -= cnt
		case "E":
			x += cnt
		case "W":
			x -= cnt
		case "F":
			shipX += math.Floor(x * cnt)
			shipY += math.Floor(y * cnt)
		case "R":
			cnt = -cnt
			fallthrough
		case "L":
			theta := cnt * math.Pi / 180
			tmp := []float64{0, 0}
			tmp[0] = x*math.Cos(theta) - y*math.Sin(theta)
			tmp[1] = x*math.Sin(theta) + y*math.Cos(theta)
			x = math.Round(tmp[0])
			y = math.Round(tmp[1])
		}
	}

	return int(math.Abs(shipX) + math.Abs(shipY))
}

func main() {
	bs, _ := ioutil.ReadFile(mainInput)
	s := string(bs)
	lines := strings.Split(s, "\n")

	actions := []action{}
	for _, line := range lines {
		matches := regex.FindStringSubmatch(line)

		cnt, _ := strconv.Atoi(matches[2])
		actions = append(actions, action{
			move: matches[1],
			cnt:  cnt,
		})
	}

	fmt.Println(part1(actions))
	fmt.Println(part2(actions))
}
