package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	total := 0
	a, b := readIn("test.txt")
	
	// Sort the slices
	sort.Ints(a)
	sort.Ints(b)

	// Zip the slices into pairs
	zipped := zip(a, b)

	// Add the pairs together
	sums := make([]int, len(zipped))
	for i, pair := range zipped {
		sums[i] = absInt(pair[0] - pair[1])
		total = total + sums[i]
	}
	fmt.Println(total)
}

func zip(a, b []int) [][2]int {
	minLength := len(a)
	if len(b) < minLength {
		minLength = len(b)
	}

	result := make([][2]int, minLength)
	for i := 0; i < minLength; i++ {
		result[i] = [2]int{a[i], b[i]}
	}
	return result
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readIn(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	defer file.Close()

	var col1, col2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		columns := strings.Fields(line)

		if len(columns) >= 2 {
			val1, err1 := strconv.Atoi(columns[0])
			if err1 == nil {
				col1 = append(col1, val1)
			}

			val2, err2 := strconv.Atoi(columns[1])
			if err2 == nil {
				col2 = append(col2, val2)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return col1, col2
}
