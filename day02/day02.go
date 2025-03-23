package day02

import (
	"advent-of-code-2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	totalSafe := part1()
	fmt.Println(totalSafe)
}

func part1() int {
	data := utils.ReadFile("inputs/02.txt")
	contents := formatInput(data)

	var totalSafe int

	for _, report := range contents {
		var direction int
		isSafe := true
		for i, value := range report {
			if i == len(report)-1 {
				break
			}

			val, _ := strconv.Atoi(value)
			nextVal, _ := strconv.Atoi(report[i+1])

			currentDirection := getDirection(val, nextVal)

			if direction == 0 {
				direction = currentDirection
			}

			if direction != currentDirection {
				isSafe = false
				break
			}

			diff := absInt(val, nextVal)

			if diff > 3 || diff < 1 {
				isSafe = false
				break
			}

		}

		if isSafe {
			totalSafe++
		}
	}
	return totalSafe
}

func formatInput(data []byte) [][]string {
	reports := strings.Split(string(data[:]), "\r\n")

	var contents [][]string
	for _, report := range reports {
		contents = append(contents, strings.Split(report, " "))
	}

	return contents
}

func getDirection(num1, num2 int) int {
	if num1 > num2 {
		return -1
	}

	if num1 < num2 {
		return 1
	}

	return 0
}

func absInt(x, y int) int {
	if x > y {
		return x - y
	}

	return y - x
}
