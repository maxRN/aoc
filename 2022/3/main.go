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

func sameLetter(a, b string) string {
	for _, lettsA := range strings.Split(a, "") {
		for _, lettsB := range strings.Split(b, "") {
			if lettsA == lettsB {
				return lettsA
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
	var rucksacks []Rucksack
	var total int
	for _, rucksack := range rucksackStrings {
		if rucksack == "" {
			continue
		}
		newRucksack := Rucksack{comp1: rucksack[0 : len(rucksack)/2], comp2: rucksack[len(rucksack)/2:]}
		newRucksack.same = sameLetter(newRucksack.comp1, newRucksack.comp2)
		value := intForLett(newRucksack.same)
		total += value
		rucksacks = append(rucksacks, newRucksack)
	}

	// fmt.Println(rucksacks)
	fmt.Println(total)

}
