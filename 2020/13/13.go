// https://adventofcode.com/2020/day/13

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"strconv"
	"strings"
)

const (
	testInput  = "test.txt"
	testInput2 = "test2.txt"
	mainInput  = "main.txt"
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

// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
var one = big.NewInt(1)

func crt(a, n []*big.Int) (*big.Int, *big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return p, x.Mod(&x, p), nil
}

func part2(buses, offsets []int) int64 {
	n := []*big.Int{}
	a := []*big.Int{}

	for i := range buses {
		n = append(n, big.NewInt(int64(buses[i])))
		a = append(a, big.NewInt(int64(offsets[i])))
	}

	ret, res, _ := crt(a, n)
	ret.Sub(ret, res)
	return ret.Int64()
}

func main() {
	bs, _ := ioutil.ReadFile(mainInput)
	s := string(bs)
	lines := strings.Split(s, "\n")

	t0, _ := strconv.Atoi(lines[0])
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

		fmt.Println(part1(t0, buses))
		fmt.Println(part2(buses, busOffsets))
	}
}
