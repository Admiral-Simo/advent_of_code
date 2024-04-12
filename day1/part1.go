package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func getInput(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func parseLine(line string) (*int, error) {
	// the first and second digit -> "1ksldjhfkshjdfkjshdfjkh89" first = 1 & second = 9
	var first rune
	var second rune
	// looping into each character `char` in the `line`
	for _, char := range line {
		// checking if the current `char` is a digit
		if unicode.IsDigit(char) {
			// if `first` digit is uninitialized then assign `first` with `char`
			if first == 0 {
				first = char
			}
			// each time we find a character `char` which is a digit just assign it to second
			// then the second becomes the last `digit` in the line
			second = char
		}
	}
	// converting the first character into an integer
	first_num := int(first - '0')
	second_num := int(second - '0')

	// doing the formula first_num
	result := first_num*10 + second_num
	return &result, nil
}

func main() {
	final_result := 0
	file, err := getInput("input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result, err := parseLine(line)
		if err != nil {
			continue
		}
		final_result += *result
	}
	fmt.Println("result:", final_result)
}
