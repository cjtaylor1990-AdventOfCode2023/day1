package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func ReadInput(fileName string) ([]string, error) {

	// Opening file input.txt
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("there was an error: %v", err)
	}

	// Creating a scanner for the opened file
	scanner := bufio.NewScanner(file)

	// Tells scanner how it wants to split (by lines)
	// bufio.ScanLines is default, but we can be explicit here
	scanner.Split(bufio.ScanLines)

	// Scanning through file and construct output
	output := []string{}
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return output, nil
}

func IsNumber(b byte) (bool, error) {
	return regexp.Match("[0-9]+", []byte{b})
}

func FindCode(entry string) (int, error) {
	foundFirst := false
	foundLast := false
	var first, last string
	for i := 0; i < len(entry); i++ {
		if foundFirst && foundLast {
			break
		}
		if !foundFirst {
			c := entry[i]
			isNum, _ := IsNumber(c)
			if isNum {
				first = string(c)
				foundFirst = true
			}
		}
		if !foundLast {
			c := entry[len(entry)-i-1]
			isNum, _ := IsNumber(c)
			if isNum {
				last = string(c)
				foundLast = true
			}
		}
	}
	return strconv.Atoi(first + last)
}

func main() {
	input, err := ReadInput("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	total := 0
	for _, v := range input {
		code, err := FindCode(v)
		if err != nil {
			fmt.Println("error finding answer:", err)
			return
		}
		total += code
	}
	fmt.Println(total)
}
