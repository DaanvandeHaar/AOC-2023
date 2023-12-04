package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	red string = "red"
	green        = "green"
	blue        = "blue"
)

type game struct {
	id 		uint
	sets 	[]map[string]uint
}

func main() {
	lines, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	games, err := parseLines(lines)
	if err != nil {
		log.Fatal()
	}

	var sum uint
	for _, g := range games {
		highest := getHighest(g.sets)
		sum += highest[red] * highest[green] * highest[blue]
	}
	fmt.Printf("sum: %v\n", sum)
}

func getHighest(sets []map[string]uint) map[string]uint {
	result := make(map[string]uint)

	for _, m := range sets {
		for key, value := range m {
			if current, ok := result[key]; !ok || value > current {
				result[key] = value
			}
		}
	}

	return result
}

func parseLines(lines []string) ([]game, error) {
	var games []game
	for i, line := range lines {
		g := game{id: uint(i + 1)}
		sets := strings.Split(strings.Split(line,": ")[1], "; ")
		for j, s := range sets {
			g.sets = append(g.sets, make(map[string]uint))
			for _, color := range strings.Split(s, ", ") {
				g.sets[j][strings.Split(color, " ")[1]] += mustAtoi(strings.Split(color, " ")[0])
			}
		}
		games = append(games, g)
	}
	return games, nil
}

func mustAtoi(str string) uint {
	count, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return uint(count)
}

func readFile() ([]string, error) {
	var lines []string

	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
