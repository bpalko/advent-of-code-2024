package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	result := 0

	for scanner.Scan() {
		line := scanner.Text()
		lines := strings.Fields(line)
		var isIncreasing bool
		isValid := true
		for i := 0; i < len(lines)-1; i++ {
			curr, _ := strconv.Atoi(lines[i])
			currNext, _ := strconv.Atoi(lines[i+1])

			if i == 0 {
				// Set the direction based on first pair
				if curr < currNext {
					isIncreasing = true
				} else {
					isIncreasing = false
				}
			}

			if !checkIncDec(isIncreasing, curr, currNext) || !checkDifference(curr, currNext) {
				isValid = false
				break // Exit the loop as soon as we find a violation
			}

		}
		if isValid {
			result += 1
		}
	}
	fmt.Println(result)
}

func checkIncDec(pres bool, a, b int) bool {
	if (!pres && a > b) || (pres && a < b) {
		return true
	} else {
		return false
	}
}

func checkDifference(a, b int) bool {
	absInt := abs(a - b)
	if 1 <= absInt && absInt <= 3 {
		return true
	} else {
		return false
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
