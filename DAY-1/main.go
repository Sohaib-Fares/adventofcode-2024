package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) != 2 {
			continue
		}

		val1, err1 := strconv.Atoi(fields[0])
		val2, err2 := strconv.Atoi(fields[1])

		if err1 != nil || err2 != nil {
			log.Printf("Error parsing line: %s", line)
			continue
		}

		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of lines:", len(list1))

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})

	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	diff := 0
	for i := 0; i < len(list1); i++ {
		if list1[i]-list2[i] < 0 {
			diff += list2[i] - list1[i]
		} else {
			diff += list1[i] - list2[i]
		}
	}

	sum := 0
	for _, listOneVal := range list1 {
		multiplicity := 0
		for _, listTwoVal := range list2 {
			if listOneVal == listTwoVal {
				multiplicity++
			}
		}
		sum += multiplicity * listOneVal
	}

	fmt.Println("Total diff:", diff)
	fmt.Println("Total sum:", sum)
	fmt.Println(25574739)
}
