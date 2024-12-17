package day01

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Run() {
	fmt.Printf("Total difference: %d\n", part1())
	fmt.Printf("Total similarity: %d\n", part2())
}

func part1() int {	
	data := readFile()
	contents := formatInput(data)

	list1, list2 := createLists(contents)

	sort.Ints(list1)
	sort.Ints(list2)

	var diffs []int
	for index, x := range list1 {
		y := list2[index]

		if (x > y) {
			diffs = append(diffs, x - y)
		} else {
			diffs = append(diffs, y - x)
		}
	}

	total := 0
	for _, element := range diffs {
		total += element
	}

	return total
}

func part2() int {
	data := readFile()
	contents := formatInput(data)
	list1, list2 := createLists(contents)

	var occurenceScores []int
	
	for _, element := range list1 {
		total := 0
		for _, compElement := range list2 {
			if (element == compElement) {
				total ++
			}
		}

		occurenceScore := total * element
		occurenceScores = append(occurenceScores, occurenceScore)
	}

	similarity := 0
	for _, element := range occurenceScores {
		similarity += element
	}

	return similarity
}

func readFile() []byte {
	data, err := os.ReadFile("inputs/01.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}

	return data
}

func formatInput(data []byte) []string {
	removedSpaces := strings.ReplaceAll(string(data[:]), "   ", ",")
	removedBreaks := strings.ReplaceAll(removedSpaces, "\r\n", ",")

	contents := strings.Split(removedBreaks, ",")

	return contents
}

func createLists(contents []string) ([]int, []int) {
	var list1 []int
	var list2 []int

	for index, element := range contents {
		castElement, _ := strconv.Atoi(element)
		if index%2 == 0 {
			list1 = append(list1, castElement)
		} else {
			list2 = append(list2, castElement)
		}
	}

	return list1, list2
}