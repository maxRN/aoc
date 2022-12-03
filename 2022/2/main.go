package main

import (
	"fmt"
	"os"
	"strings"
)

type MU struct {
	m string
	o string
}

func getPointForRPS(rps string) int {
	switch rps {
	case "A":
		return 1
	case "B":
		return 2
	default:
		return 3
	}
}

func determineRPS(opp, outcome string) int {
	var ourChoice string
	switch outcome {
	case "Y":
		ourChoice = opp
	case "X":
		if opp == "A" {
			ourChoice = "C"
		} else if opp == "B" {
			ourChoice = "A"
		} else {
			ourChoice = "B"
		}
	default:
		if opp == "A" {
			ourChoice = "B"
		} else if opp == "B" {
			ourChoice = "C"
		} else {
			ourChoice = "A"
		}

	}

	return getPointForRPS(ourChoice)
}

func (mU *MU) getPoints() (total int) {
	switch mU.m {
	case "X":
		total += 0
	case "Y":
		total += 3
	default:
		total += 6
	}

	total += determineRPS(mU.o, mU.m)
	// determine score based on outcome
	if mU.m == mU.o {
		total += 3
		return total
	}

	if iWin(mU.m, mU.o) {
		total += 6
	} else {
		total += 0
	}

	return total

}

func iWin(m, o string) bool {
	if (m == "A" && o == "C") || (m == "B" && o == "A") || (m == "C" && o == "B") {
		return true
	}

	return false
}

func main() {
	input, _ := os.ReadFile("input.txt")
	var matchUps []MU
	for _, line := range strings.Split(string(input), "\n") {
		// ignore last line
		if len(line) == 0 {
			continue
		}

		choices := strings.Split(line, " ")

		matchUp := MU{m: choices[1], o: choices[0]}
		matchUps = append(matchUps, matchUp)
	}

	var total int
	for _, matchup := range matchUps {
		total += matchup.getPoints()
	}

	fmt.Println(total)

}
