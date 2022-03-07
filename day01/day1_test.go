package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func getTestData() string {
	testData := `199
200
208
210
200
207
240
269
260
263`
	return testData
}

func TestDay1_TestParseInput(t *testing.T) {
	input := getTestData()

	expected := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	actual := parseInput(input)

	if cmp.Equal(expected, actual) == false {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDay1_TestPart1Solution(t *testing.T) {
	expected := 7
	actual := part1Solution([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestDay1_TestPart2Solution(t *testing.T) {
	expected := 5
	actual := part2Solution([]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}

}
