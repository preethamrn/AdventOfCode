// https://adventofcode.com/2020/day/7

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var (
	parentRegex = regexp.MustCompile(`(?P<bag>.*) bags contain`)
	childRegex  = regexp.MustCompile(`\d+ (?P<bag>.*?) bags?[,.]`)
)

var (
	visited = map[int]struct{}{}
	bagsMap = map[string]int{}
	graph   = map[int][]int{}
)

func dfs(i int) int {
	if _, ok := visited[i]; ok {
		return 0
	}
	visited[i] = struct{}{}

	cnt := 1
	for _, j := range graph[i] {
		cnt += dfs(j)
	}
	return cnt
}

func main() {
	lastBagIdx := 0

	bs, _ := ioutil.ReadFile("7.1.txt")
	s := string(bs)
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		matches := parentRegex.FindAllStringSubmatch(line, -1)
		parent := matches[0][1]
		if _, ok := bagsMap[parent]; !ok {
			bagsMap[parent] = lastBagIdx
			lastBagIdx++
		}

		matches = childRegex.FindAllStringSubmatch(line, -1)
		children := []string{}
		for _, match := range matches {
			child := match[1]
			children = append(children, child)
			if _, ok := bagsMap[child]; !ok {
				bagsMap[child] = lastBagIdx
				lastBagIdx++
			}
			graph[bagsMap[child]] = append(graph[bagsMap[child]], bagsMap[parent])
		}
	}

	fmt.Println(dfs(bagsMap["shiny gold"]) - 1)
	// NOTE: dfs returns the count including the source node. We subtract 1 to get the final answer.
}
