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

var colors = []string{red, green, blue}

type game struct {
    id         uint
    sets     []map[string]uint
}

type set struct {
    redCount     uint
    greenCount    uint
    blueCount     uint
}

var rgx = regexp.MustCompile(`\D`)

func main() {
    lines, err := readFile()
    if err != nil {
        log.Fatal(err)
    }

    games, err := parseLines(lines)
    if err != nil {
        log.Fatal()
    }

    var count uint
    for _, g := range games {
        if allValid(g) {
            count += g.id
        }
    }

    fmt.Printf("sum: %v", count)
}

func allValid(g game) bool {
    for _, s := range g.sets {
        if s[red] > 12 || s[green] > 13 || s[blue] > 14 {
            return false
        }
    }
    return true
}

func parseLines(lines []string) ([]game, error) {
    var games []game
    for i, line := range lines {
        g := game{id: uint(i + 1)}
        sets := strings.Split(strings.Split(line,": ")[1], "; ")
        for j, s := range sets {
            g.sets = append(g.sets, make(map[string]uint))
            for _, color := range strings.Split(s, ", ") {
                for _, c := range colors {
                    if strings.Contains(color, c) {
                        g.sets[j][c] += getCount(color)
                    }
                }
            }
        }
        games = append(games, g)
    }
    return games, nil
}

func getCount(str string) uint {
    count, err := strconv.Atoi(rgx.ReplaceAllString(str, ""))
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