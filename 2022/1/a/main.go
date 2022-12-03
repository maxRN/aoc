package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.ReadFile("input.txt")
	calories := strings.Split(string(file), "\n")
	var finishedCals [][]string
	var tmpCal []string
	for _, value := range calories {
		tmpCal = append(tmpCal, value)
		if len(value) == 0 {
			finishedCals = append(finishedCals, tmpCal)
			tmpCal = nil
		}
	}

	var calsOnly []int
	var tmpInt int
	for _, cals := range finishedCals {

		for _, calo := range cals {
			numVal, _ := strconv.Atoi(calo)
			tmpInt += numVal
		}
		calsOnly = append(calsOnly, tmpInt)
		tmpInt = 0
	}

	sort.Slice(calsOnly, func(i, j int) bool { return calsOnly[i] > calsOnly[j] })

	var total int
	for _, val := range calsOnly[0:3] {
		total += val
	}

	fmt.Println(total)
}
