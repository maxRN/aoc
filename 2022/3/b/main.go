package main

import (
	"fmt"
	"os"
	"strings"
)

type Rucksack struct {
	comp1 string
	comp2 string
	same  string
}

func sameBadge(a, b, c string) string {
	for _, lettsA := range strings.Split(a, "") {
		for _, lettsB := range strings.Split(b, "") {
			for _, lettsC := range strings.Split(c, "") {
				if (lettsA == lettsB) && (lettsA == lettsC) {
					return lettsA
				}
			}
		}
	}
	panic("wtf")
}

func intForLett(same string) int {
	runeSame := []rune(same)[0]
	// fmt.Printf("string %s is rune %d", same, runeSame)

	if int(runeSame) >= 97 {
		return int(runeSame) - 96
	} else {
		return int(runeSame) - 38
	}

}

func main() {
	input, _ := os.ReadFile("input.tx")
	rucksackStrings := strings.Split(string(input), "\n")
	var rucksackThrees [][]string
	var temp []string
	for i, rucksack := range rucksackStrings {
		if rucksack == "" {
			continue
		}
		temp = append(temp, rucksack)
		if i%3 == 2 {
			rucksackThrees = append(rucksackThrees, temp)
			temp = nil
		}
	}

	var total int
	for _, rucksackThree := range rucksackThrees {
		badge := sameBadge(rucksackThree[0], rucksackThree[1], rucksackThree[2])
		value := intForLett(badge)
		total += value
	}

	fmt.Println(total)
}
