package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindsCorrectCodesForNonSpelledDigits(t *testing.T) {
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expected := []int{12, 38, 15, 77}
	tr := CreateNumberTranslator()
	rtr := CreateReverseNumberTranslator()

	for i, v := range input {
		actual, err := FindCode(v, tr, rtr)
		assert.NoError(t, err)
		assert.Equal(t, expected[i], actual)
	}
}

func TestFindsCorrectTotalsForNonSpelledDigits(t *testing.T) {
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	expected := 142

	actual, err := FindTotal(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
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
	tr := CreateNumberTranslator()
	rtr := CreateReverseNumberTranslator()

	for i, v := range input {
		actual, err := FindCode(v, tr, rtr)
		assert.NoError(t, err)
		assert.Equal(t, expected[i], actual)
	}
}

func TestFindsCorrectTotalsForSpelledDigits(t *testing.T) {
	input := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}
	expected := 281

	actual, err := FindTotal(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestFindsCorrectCodesForOverlappingSpelledNumbers(t *testing.T) {
	input := []string{
		"oneight",
	}
	expected := []int{18}
	tr := CreateNumberTranslator()
	rtr := CreateReverseNumberTranslator()

	for i, actual := range input {
		actual, err := FindCode(actual, tr, rtr)
		assert.NoError(t, err)
		assert.Equal(t, expected[i], actual)
	}
}
