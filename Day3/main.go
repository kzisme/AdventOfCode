package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var test = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

type Node struct {
	count  int
	mcb    int
	lcb    int
	Zeroes int
	Ones   int
}

type Track struct {
	inputs      []int
	OnesCount   int
	ZeroesCount int
	count       int
	isActive    bool
}

var inputs []int

func activeCount(activeItems []Track) int {
	var activeItemsCount = 0
	for i := 0; i < len(activeItems); i++ {
		if activeItems[i].isActive {
			activeItemsCount++
		}
	}

	// If this returns 2 - we have to break the tie
	return activeItemsCount
}

func main() {

	file, err := os.Open("C:\\Users\\kzism\\source\\repos\\AdventOfCode\\Day3\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += scanner.Text()
		lines += "\n"
	}

	var val1, val2 string

	nums := strings.Split(lines, "\n")
	for i := 0; i < len(nums[0]); i++ {
		var zeroes, ones int
		for _, b := range nums {
			if len(b) != 0 {
				if b[i] == '0' {
					zeroes++
				} else {
					ones++
				}
			}
		}

		if zeroes > ones {
			val1 += "0"
			val2 += "1"
		} else {
			val1 += "1"
			val2 += "0"
		}
	}

	mostCommon, err := strconv.ParseInt(val1, 2, 64)
	if err != nil {
		panic(err)
	}

	leastCommon, err := strconv.ParseInt(val2, 2, 64)
	if err != nil {
		panic(err)
	}

	var part1 = mostCommon * leastCommon

	fmt.Println("Part 1 Answer: " + strconv.Itoa(int(part1)))

	//pt2

	//var trie *Node
	numbers := []Node{}
	setup := []Track{}
	//var binaryLoc = [5]int{}

	// 00100
	setupArrays := strings.Split(test, "\n")
	for i := 0; i < 50; i++ {
		var binaryTest = [5]int{}

		var onesCount = 0
		var zeroesCount = 0
		if i < 12 {
			for n, r := range setupArrays[i] {
				c := string(r)
				fmt.Println(c)
				// I don't quite understand how this works but it does from rune ---> int
				binaryTest[n] = int(r - '0')
				if binaryTest[n] == 1 {
					onesCount++
				} else {
					zeroesCount++
				}

			}
			s := Track{inputs: binaryTest[:], ZeroesCount: zeroesCount, OnesCount: onesCount, count: i, isActive: true}
			onesCount = 0
			zeroesCount = 0
			setup = append(setup, s)
		}
	}

	report := strings.Split(test, "\n")
	for i := 0; i < len(report[0]); i++ {
		var zeroes, ones int
		for _, b := range report {
			if len(b) != 0 {
				if b[i] == '0' {
					zeroes++
				} else {
					ones++
				}
			}
		}
		//insert here
		n := Node{Zeroes: zeroes, Ones: ones}
		numbers = append(numbers, n)
	}

	//ar mostCommonBit string
	for i := 0; i < len(numbers); i++ {
		if numbers[i].Ones == numbers[i].Zeroes {
			numbers[i].mcb = 1
		} else {
			if numbers[i].Ones > numbers[i].Zeroes {
				numbers[i].mcb = 1
			} else {
				numbers[i].mcb = 0
			}
		}
		numbers[i].count = i
	}

	for nodeIndex := 0; nodeIndex <= 4; nodeIndex++ {
		//fmt.Println(numbers[i])
		for x := 0; x <= 11; x++ {
			if setup[x].inputs[nodeIndex] != numbers[nodeIndex].mcb && setup[x].isActive == true {
				setup[x].isActive = false
			}
		}
	}
}
