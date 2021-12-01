// https://adventofcode.com/2020/day/7

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	testInput  = "7.2_test.txt"
	testInput2 = "../7.1/7.1_test.txt"
	mainInput  = "../7.1/7.1.txt"
)

var (
	parentRegex = regexp.MustCompile(`(?P<bag>.*) bags contain`)
	childRegex  = regexp.MustCompile(`(?P<cnt>\d+) (?P<bag>.*?) bags?[,.]`)
)

var (
	visited = map[int]int{}
	bagsMap = map[string]int{}
	graph   = map[int][]bagCnt{}
)

type bagCnt struct {
	child int
	cnt   int
}

func dfs(i int) int {
	if _, ok := visited[i]; ok {
		return visited[i]
	}

	visited[i] = 1
	for _, j := range graph[i] {
		visited[i] += j.cnt * dfs(j.child)
	}
	return visited[i]
}

func main() {
	lastBagIdx := 0

	bs, _ := ioutil.ReadFile(mainInput)
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
			cnt, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Printf("ERROR parsing: %s\n", match[1])
			}
			child := match[2]
			children = append(children, child)
			if _, ok := bagsMap[child]; !ok {
				bagsMap[child] = lastBagIdx
				lastBagIdx++
			}
			graph[bagsMap[parent]] = append(graph[bagsMap[parent]], bagCnt{bagsMap[child], cnt})
		}
	}

	fmt.Println(dfs(bagsMap["shiny gold"]) - 1)
	// NOTE: dfs returns the count including the source node. We subtract 1 to get the final answer.
}
