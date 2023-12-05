package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var stringNumberMapping map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findFirstDigit(line string) int {
	for _, chr := range line {
		if unicode.IsDigit(chr) {
			value, _ := strconv.Atoi(string(chr))
			return value
			// return int(chr - '0')
		}
	}
	return 0 // cant be possible
}

func findLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if runeValue := rune(line[i]); unicode.IsDigit(runeValue) {
			value, _ := strconv.Atoi(string(runeValue))
			return value
			// return int(runeValue - '0')
		}
	}
	return 0 // cant be possible
}

func findFirstDigitOrStringEqv(line string) int {
	accum := ""
	for i := 0; i < len(line); i++ {
		accum += string(line[i])
		for key := range stringNumberMapping {
			if strings.HasSuffix(accum, key) {
				return stringNumberMapping[key]
			}
		}
		if runeValue := rune(line[i]); unicode.IsDigit(runeValue) {
			value, _ := strconv.Atoi(string(runeValue))
			return value
		}
	}
	return 0
}

func findLastDigitOrStringEqv(line string) int {
	accum := ""
	for i := len(line) - 1; i >= 0; i-- {
		accum = string(line[i]) + accum
		for key := range stringNumberMapping {
			if strings.HasPrefix(accum, key) {
				return stringNumberMapping[key]
			}
		}
		if runeValue := rune(line[i]); unicode.IsDigit(runeValue) {
			value, _ := strconv.Atoi(string(runeValue))
			return value
		}
	}
	return 0
}

func partOneSolution() {
	file, err := os.Open("./day01/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	result := 0

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)

		result += findFirstDigit(line)*10 + findLastDigit(line)
	}
	fmt.Println("part1", result)
}

func partTwoSolution() {
	file, err := os.Open("./day01/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	result := 0

	for {
		line, err := r.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}

		line = strings.TrimSpace(line)
		result += findFirstDigitOrStringEqv(line)*10 + findLastDigitOrStringEqv(line)
	}
	fmt.Println("part2", result)
}

func main() {
	partOneSolution()
	partTwoSolution()
}
