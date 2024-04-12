package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func parseSet(set string) bool {
	// a color is seperated by ','
	red, green, blue := 0, 0, 0
	words := strings.Split(set, ",")
	for _, word := range words {
		trimmedWord := strings.TrimSpace(word)
		values := strings.Split(trimmedWord, " ")
		value, _ := strconv.Atoi(values[0])
		color := values[1]
		switch color {
		case "red":
			red += value
			break
		case "green":
			green += value
			break
		case "blue":
			blue += value
			break
		}
	}
	return red <= 12 && green <= 13 && blue <= 14
}

func isPossibleGame(game string) bool {
	refactoredLine := game[8:]
	// a set is seperated by ';'
	sets := strings.Split(refactoredLine, ";")
	for _, set := range sets {
		if !parseSet(set) {
			return false
		}
	}
	return true
}

func main() {
	final_result := 0
	file, err := getInput("../input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	i := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := scanner.Text()
		valid := isPossibleGame(game)
		if valid {
			final_result += i
		}
		i++
	}
	fmt.Println("result:", final_result)
}
