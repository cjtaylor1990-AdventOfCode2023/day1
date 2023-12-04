package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindsCorrectCodesForBasicExamples(t *testing.T) {
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expected := []int{12, 38, 15, 77}

	for i, v := range input {
		actual, err := FindCode(v)
		assert.NoError(t, err)
		assert.Equal(t, expected[i], actual)
	}
}

func TestFindsCorrectCodesForSpelledDigits(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen"}
	expected := []int{29, 83, 13, 24, 42, 14, 76}

	for i, v := range input {
		actual, err := FindCode(v)
		assert.NoError(t, err)
		assert.Equal(t, expected[i], actual)
	}
}
