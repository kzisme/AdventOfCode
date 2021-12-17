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

	numbers := []Node{}
	setup := []Track{}
	var oxygenRating []int
	var co2ScrubberRating []int
	var lifeSupportRating int64

	// 00100
	setupArrays := strings.Split(lines, "\n")
	for i := 0; i < len(lines); i++ {
		var binaryTest = [12]int{}

		var onesCount = 0
		var zeroesCount = 0
		if i < 1000 {
			for n, r := range setupArrays[i] {
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

	report := strings.Split(lines, "\n")
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

	// Oxygen Scrubber Rating - Most Common ---> 1
	var nodeMcb = 0
	for nodeIndex := 0; nodeIndex <= 11; nodeIndex++ {
		nodeMcb = getMostCommonBit(setup, nodeIndex)
		for x := 0; x <= 999; x++ {
			if setup[x].inputs[nodeIndex] != nodeMcb && setup[x].isActive == true {
				setup[x].isActive = false
			}
		}
		nodeMcb = getMostCommonBit(setup, nodeIndex)

		// We found our oxygen rating
		if remainingMcb(setup) {
			oxygenRating = outputOxygenRating(setup)
			fmt.Printf("%v", oxygenRating)
		}
	}

	// Reset all items to active
	resetActive(setup)

	var nodeLcb = 0
	for nodeIndex := 0; nodeIndex <= 11; nodeIndex++ {
		nodeLcb = getLeastCommonBit(setup, nodeIndex)
		for x := 0; x <= 999; x++ {
			if setup[x].inputs[nodeIndex] != nodeLcb && setup[x].isActive == true {
				setup[x].isActive = false
			}
		}
		nodeLcb = getLeastCommonBit(setup, nodeIndex)

		// We found our oxygen rating
		if remainingLcb(setup) {
			co2ScrubberRating = outputCo2ScrubberRating(setup)
			fmt.Printf("%v", co2ScrubberRating)
		}
	}

	lifeSupportRating = join(co2ScrubberRating) * join(oxygenRating)

	fmt.Println("Part 2 Answer: " + strconv.Itoa(int(lifeSupportRating)))
}

func getMostCommonBit(items []Track, bitPos int) int {

	var zeroesCount = 0
	var onesCount = 0
	for i := 0; i < len(items); i++ {
		if items[i].inputs[bitPos] == 1 && items[i].isActive {
			onesCount++
		} else if items[i].inputs[bitPos] == 0 && items[i].isActive {
			zeroesCount++
		}
	}
	if zeroesCount > onesCount {
		return 0
	} else if zeroesCount < onesCount {
		return 1
	} else {
		// Tie in number of bits - choose 1
		return 1
	}
}

func getLeastCommonBit(items []Track, bitPos int) int {
	var zeroesCount = 0
	var onesCount = 0

	for i := 0; i < len(items); i++ {
		if items[i].inputs[bitPos] == 1 && items[i].isActive {
			onesCount++
		} else if items[i].inputs[bitPos] == 0 && items[i].isActive {
			zeroesCount++
		}
	}
	if zeroesCount < onesCount {
		return 0
	} else if zeroesCount > onesCount {
		return 1
	} else {
		// Tie in number of bits - choose 0
		return 0
	}
}

func remainingMcb(items []Track) bool {
	var count = len(items)
	for i := 0; i < len(items); i++ {
		if !items[i].isActive {
			count--
		}
	}

	if count == 1 {
		return true
	} else {
		return false
	}
}

func remainingLcb(items []Track) bool {
	var count = len(items)
	for i := 0; i < len(items); i++ {
		if !items[i].isActive {
			count--
		}
	}

	if count == 1 {
		return true
	} else {
		return false
	}
}

func outputOxygenRating(items []Track) []int {
	var retVal []int
	for i := 0; i < len(items); i++ {
		if items[i].isActive {
			retVal = items[i].inputs
		}
	}
	return retVal
}

func outputCo2ScrubberRating(items []Track) []int {
	var retVal []int
	for i := 0; i < len(items); i++ {
		if items[i].isActive {
			retVal = items[i].inputs
		}
	}
	return retVal
}

func resetActive(items []Track) {
	for i := 0; i < len(items); i++ {
		if !items[i].isActive {
			items[i].isActive = true
		}
	}
}

func join(nums []int) int64 {
	var str string
	for i := range nums {
		str += strconv.Itoa(nums[i])
	}
	val, _ := strconv.ParseInt(str, 2, 64)

	return val
}
