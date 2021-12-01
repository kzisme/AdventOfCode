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
		fmt.Println(scanner.Text())
	}

	for i := range s {
		if i != 1999 {
			if s[i] < s[i+1] {
				numberIncreasesCount++
			}
		}
	}

	//printSlice(s)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
