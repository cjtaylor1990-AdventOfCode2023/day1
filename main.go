package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	regexString        = "[0-9]|one|two|three|four|five|six|seven|eight|nine"
	reverseRegexString = "[0-9]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin"
)

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

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

type NumberTranslator struct {
	Translation map[string]string
}

func (t NumberTranslator) Translate(input string) string {
	val, ok := t.Translation[input]
	if ok {
		return val
	}
	return input
}

func FindCode(entry string, translator NumberTranslator, reverseTranslator NumberTranslator) (int, error) {
	re := regexp.MustCompile(regexString)
	rr := regexp.MustCompile(reverseRegexString)
	first := translator.Translate(re.FindString(entry))
	if len(first) == 0 {
		return 0, fmt.Errorf("no matches found")
	}
	last := reverseTranslator.Translate(rr.FindString(reverse(entry)))
	code, err := strconv.Atoi(first + last)
	if err != nil {
		return 0, fmt.Errorf("error converting code: %v", err)
	}
	return code, nil
}

func CreateNumberTranslator() NumberTranslator {
	translation := map[string]string{
		"one": "1", "two": "2", "three": "3",
		"four": "4", "five": "5", "six": "6",
		"seven": "7", "eight": "8", "nine": "9",
	}
	return NumberTranslator{Translation: translation}
}

func CreateReverseNumberTranslator() NumberTranslator {
	translation := map[string]string{
		"eno": "1", "owt": "2", "eerht": "3",
		"ruof": "4", "evif": "5", "xis": "6",
		"neves": "7", "thgie": "8", "enin": "9",
	}
	return NumberTranslator{Translation: translation}
}

func FindTotal(input []string) (int, error) {
	translator := CreateNumberTranslator()
	reverseTranslator := CreateReverseNumberTranslator()
	total := 0
	for _, v := range input {
		code, err := FindCode(v, translator, reverseTranslator)
		if err != nil {
			return 0, fmt.Errorf("error finding total: %v", err)
		}
		total += code
	}
	return total, nil
}

func main() {
	input, err := ReadInput("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	total, err := FindTotal(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total: ", total)
}
