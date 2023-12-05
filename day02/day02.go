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
)

type GameSet map[string]int

var maxValuesForColors GameSet = GameSet{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func partOneSolution() {
	file, err := os.Open("./day02/game.txt")
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
		tmp := strings.Split(line, ": ")
		stringGameID, gameSets := strings.TrimPrefix(tmp[0], "Game "), tmp[1]
		gameID, _ := strconv.Atoi(stringGameID) // err is not possible
		isGameValid := true

	PAIR_LOOP:
		// iterate over game sets of certain game
		for _, set := range strings.Split(gameSets, ";") {
			set = strings.TrimSpace(set)
			// iterate over set pairs "amount color"
			for _, pair := range strings.Split(set, ", ") {
				pairSplitted := strings.Split(pair, " ")
				amount, _ := strconv.Atoi(pairSplitted[0])
				if amount > maxValuesForColors[pairSplitted[1]] {
					isGameValid = !isGameValid
					break PAIR_LOOP
				}
			}
		}
		if isGameValid {
			result += gameID
		}
	}
	fmt.Println("part1", result)
}

func partTwoSolution() {
	file, err := os.Open("./day02/game.txt")
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
		tmp := strings.Split(line, ": ")
		gameSets := tmp[1]
		maxColorsPerGame := GameSet{}

		// iterate over game sets of certain game
		for _, set := range strings.Split(gameSets, ";") {
			set = strings.TrimSpace(set)
			// iterate over set pairs "amount color"
			for _, pair := range strings.Split(set, ", ") {
				pairSplitted := strings.Split(pair, " ")
				color := pairSplitted[1]
				amount, _ := strconv.Atoi(pairSplitted[0])
				if maxColorsPerGame[color] < amount {
					maxColorsPerGame[color] = amount
				}
			}
		}
		powerOfGame := 1
		for _, v := range maxColorsPerGame {
			powerOfGame *= v
		}
		result += powerOfGame
	}
	fmt.Println("part2", result)
}

func main() {
	partOneSolution()
	partTwoSolution()
}
