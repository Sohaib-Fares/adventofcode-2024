package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func exportData(input *os.File) ([][]int, error) {

	scanner := bufio.NewScanner(input)
	out := make([][]int, 0)
	
	for scanner.Scan() {

		line := scanner.Text()
		stringFields := strings.Fields(line)

		var lineNumbers []int
		for _, s := range stringFields {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Printf("Error converting %s to int: %v\n", s, err)
				continue
			}
			lineNumbers = append(lineNumbers, num)
		}
		out = append(out, lineNumbers)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func gradIncreasingValidation(report []int) bool {
    safe := gradIncreasing(report)
    if safe == false {
        for i := range report {
            reportCopy := make([]int, len(report))
            copy(reportCopy, report)
            newReport := append(reportCopy[:i], reportCopy[i+1:]...)
            oneRemovedSafe := gradIncreasing(newReport)
            if oneRemovedSafe == true {
                return true
            }
        }
    }
    return safe
}

func gradDecreasingValidation(report []int) bool {
	safe := gradDecreasing(report)
	if safe == false {
		for i := range report {
			reportCopy := make([]int, len(report))
			copy(reportCopy, report)
			newReport := append(reportCopy[:i], reportCopy[i+1:]...)
			oneRemovedSafe := gradDecreasing(newReport)
			if oneRemovedSafe == true {
				return true
			}
		}
	}
	return safe
}

func gradIncreasing(input []int) bool {

	if len(input) == 0 {
		return true
	}

	for i := 0; i < len(input)-1; i++ {
		if input[i] >= input[i+1] {
			return false
		}
		if input[i+1]-input[i] > 3 || input[i+1]-input[i] < 1 {
			return false
		}
	}
	return true
}

func gradDecreasing(input []int) bool {

	if len(input) == 0 {
		return true
	}

	for i := 0; i < len(input)-1; i++ {
		if input[i] <= input[i+1] {
			return false
		}
		if input[i]-input[i+1] > 3 || input[i]-input[i+1] < 1 {
			return false
		}
	}
	return true
}

func safeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if gradIncreasing(report) || gradDecreasing(report) {
			count++
		}
	}
	return count
}

func safeReportsWithDampener(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if gradIncreasingValidation(report) || gradDecreasingValidation(report) {
			count++
		}
	}
	return count
}

func main() {

	f, err := os.Open("DAY-2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := exportData(f)

	if err != nil {
		log.Fatal(err)
	}

	safeReports := safeReports(data)
	fmt.Println(safeReports)

	safeReportsWithDampenerCount := safeReportsWithDampener(data)
	fmt.Println(safeReportsWithDampenerCount)

}
