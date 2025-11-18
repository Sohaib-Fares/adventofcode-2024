package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func exportData(input *os.File) ([][]string, error) {

	scanner := bufio.NewScanner(input)
	out := make([][]string, 0)
	var validID = regexp.MustCompile(`mul\(\d+,\d+\)`)
	for scanner.Scan() {
		line := scanner.Text()
		strings := validID.FindAllString(line, -1)
		out = append(out, strings)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func flattenStrings(input [][]string) []string {
	result := make([]string, 0, len(input))
	for _, innerSlice := range input {
		result = append(result, innerSlice...)
	}
	return result
}

func parseMul(input string) (int, int, error) {

	// fmt.Println("Input:", input)

	trimmed := strings.TrimPrefix(input, "mul(")
	// fmt.Println("After TrimPrefix:", trimmed)

	trimmed = strings.TrimSuffix(trimmed, ")")
	// fmt.Println("After TrimSuffix:", trimmed)

	parts := strings.Split(trimmed, ",")
	// fmt.Printf("Parts: [%q, %q]\n", parts[0], parts[1])

	if len(parts) != 2 {
		return 0, 0, errors.New("invalid format")
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

func execMul(str string) (int, error) {
	x, y, err := parseMul(str)
	if err != nil {
		return 0, err
	}
	return x * y, nil
}

func processData(data []string) int {
	sum := 0
	for _, op := range data {
		result, err := execMul(op)
		if err != nil {
			fmt.Printf("Error processing %s: %v\n", op, err)
		}
		sum += result
	}
	return sum
}

func getInput(input string) [][]string {
	f, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := exportData(f)

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	data := getInput("DAY-3/input.txt")
	flattenedData := flattenStrings(data)
	result := processData(flattenedData)
	fmt.Printf("The result is: %d\n", result)

}
