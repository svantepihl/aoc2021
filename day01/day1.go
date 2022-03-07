package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
	"sync"
)

const INPUT_FILE_PATH = "./input.txt"

func parseInput(input string) (ints []int) {
	lines := strings.Split(input, "\n")

	// loop over lines
	for _, line := range lines {
		if i, err := strconv.Atoi(line); err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func worker(f func(), wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

func part1Solution(ints []int) (increaseCount int) {
	length := len(ints)
	for i := 1; i < length; i++ {
		if ints[i] > ints[i-1] {
			increaseCount++

		}
	}
	return increaseCount
}

func part2Solution(ints []int) (increaseCount int) {
	length := len(ints)
	for i := 1; i < length-2; i++ {
		if ints[i+2] > ints[i-1] {
			increaseCount++
		}
	}
	return increaseCount
}

func main() {
	inputBytes, err := os.ReadFile(INPUT_FILE_PATH)

	if err != nil {
		panic(err)
	}

	runPartOne := flag.Bool("runPartOne", true, "Whether to run part 1")
	runPartTwo := flag.Bool("runPartTwo", true, "Whether to run part 2")
	flag.Parse()

	input := parseInput(strings.TrimSpace(string(inputBytes)))

	waitGroup := new(sync.WaitGroup)

	if *runPartOne {
		worker(func() {
			println("Part 1:", part1Solution(input))
		}, waitGroup)
	}

	if *runPartTwo {
		worker(func() {
			println("Part 2:", part2Solution(input))
		}, waitGroup)
	}

	waitGroup.Wait()
}
