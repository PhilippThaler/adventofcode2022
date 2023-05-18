package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readFile, err := os.Open("./input")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	current := 0
  elves := []int{}
	for fileScanner.Scan() {
		row := fileScanner.Text()
		if len(row) != 0 {
			cur, err := strconv.Atoi(row)
			check(err)
			current += cur
		} else {
      elves = append(elves, current)
			current = 0
		}
	}
  sort.Ints(elves)

  // Part 1 Output
  fmt.Println("Part 1 Output: ", elves[len(elves)-1])

  topThreeElves := elves[len(elves)-3:]
  topCalories := 0
  for i := 0; i < len(topThreeElves); i++ {
    topCalories += topThreeElves[i]
  }
  fmt.Println("Part 2 Output: ", topCalories)
	readFile.Close()
}
