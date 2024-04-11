package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getInput(fileName string) string {
	// opening the file
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// closing the file after the termination of the program
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileSize := fileInfo.Size()
	// make buffer based on the size of the file
	buffer := make([]byte, fileSize)

	// put the file into the buffer
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(buffer)
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
	first_num, _ := strconv.Atoi(string(first))
	second_num, _ := strconv.Atoi(string(second))

	// doing the formula first_num
	result := first_num*10 + second_num
	return &result, nil
}

func main() {
	final_result := 0
	for _, line := range strings.Split(getInput("input.txt"), "\n") {
		result, err := parseLine(line)
		if err != nil {
			continue
		}
		final_result += *result
	}
	fmt.Println("result:", final_result)
}
