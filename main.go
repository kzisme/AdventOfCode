package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("C:\\Users\\kzism\\source\\repos\\AdventOfCode\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var s []int
	var numberIncreasesCount = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		inputInt, _ := strconv.Atoi(scanner.Text())
		s = append(s, inputInt)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 2 Answer
	MaxSubarraySum(s, 3)

	for i := range s {
		if i != 1999 {
			if s[i] < s[i+1] {
				numberIncreasesCount++
			}
		}
	}

	// Part 1 Answer
	fmt.Println(numberIncreasesCount)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Takes slice + window size - in this case 3
func MaxSubarraySum(slice []int, windowSize int) int {

	currentSum := 0
	tempSum := 0
	windowSumIncreasedCount := 0

	//Edge case checking if the slice length is less than the window size
	if windowSize > len(slice) {
		return 0
	}

	// First window Sum
	for i := 0; i < windowSize; i++ {
		currentSum += slice[i]
	}

	tempSum = currentSum

	// Subsequent windows
	for i := windowSize; i < len(slice); i++ {
		// First Iteration:
		// A Window: 581
		// B Window: 612 = 581 - 187 + 218
		tempSum = tempSum - slice[i-windowSize] + slice[i]
		//fmt.Println(currentSum)
		//fmt.Println(tempSum)

		if currentSum < tempSum {
			windowSumIncreasedCount++
		}

		currentSum = tempSum
	}
	return windowSumIncreasedCount
}
