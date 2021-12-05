package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	Direction string
	Movement  int
}

func main() {
	file, err := os.Open("C:\\Users\\kzism\\source\\repos\\AdventOfCode\\Day2\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var horizontalPos = 0
	var depthPos = 0
	var aim = 0

	s := []Submarine{}

	// s = append(s, Submarine{Movement: 13, Direction: "Forward"})
	// fmt.Println(s)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		input := scanner.Text()

		inputLine := strings.Fields(input)
		// Can I do this inline?
		i, _ := strconv.Atoi(inputLine[1])

		s = append(s, Submarine{Direction: inputLine[0], Movement: i})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	for i := range s {
		// Not sure if I need this if check or not...maybe index out of range?
		if i != 1000 {
			if s[i].Direction == "forward" {
				horizontalPos += s[i].Movement
			} else {
				// If we made it here it's either Up or Down
				if s[i].Direction == "down" {
					depthPos += s[i].Movement
				} else {
					// We are going up so ~decreasing~ depth
					depthPos -= s[i].Movement
				}
			}
		}
	}

	fmt.Println("Depth Position: " + strconv.Itoa(depthPos) + " Horizontal Position: " + strconv.Itoa(horizontalPos))
	var answer1 = horizontalPos * depthPos
	fmt.Println(strconv.Itoa(answer1))

	// Reset for part 2
	depthPos = 0
	horizontalPos = 0

	// Part 2
	for i := range s {
		// Not sure if I need this if check or not...maybe index out of range?
		if i != 1000 {
			if s[i].Direction == "forward" {
				horizontalPos += s[i].Movement
				depthPos += aim * s[i].Movement
			} else {
				// If we made it here it's either Up or Down
				if s[i].Direction == "down" {
					aim += s[i].Movement
				} else {
					// We are going up so ~decreasing~ depth
					aim -= s[i].Movement
				}
			}
		}
	}

	fmt.Println("Depth Position Day 2 : " + strconv.Itoa(depthPos) + " Horizontal Position Day 2: " + strconv.Itoa(horizontalPos) + " Aim: " + strconv.Itoa(aim))
	var answer2 = horizontalPos * depthPos
	fmt.Println(strconv.Itoa(answer2))

}
