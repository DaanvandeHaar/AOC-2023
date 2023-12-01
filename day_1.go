package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	sum := 0
	lines, err := readFile()
	if err != nil {
		log.Fatal(err)
	}

	rgx := regexp.MustCompile(`\D`)
	for _, line := range lines {
		for i, digit := range digits {
			line = strings.Replace(line, digit, fmt.Sprintf("%v%v%v", digit[0:1], i + 1, digit[len(digit)-1:]), -1)
		}

		numbers := rgx.ReplaceAllString(line, "")

		number := fmt.Sprintf("%v%v", numbers[0:1], numbers[len(numbers)-1:])
		i, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}

		sum += i
	}

	fmt.Printf("sum: %v\n", sum)
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
