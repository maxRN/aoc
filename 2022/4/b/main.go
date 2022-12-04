package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	a int
	b int
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	var inputPairs [][]Pair
	for _, pairing := range strings.Split(string(inputFile), "\n") {
		if pairing == "" {
			continue
		}
		pair := strings.Split(pairing, ",")
		pair1Tmp := strings.Split(pair[0], "-")
		pair2Tmp := strings.Split(pair[1], "-")
		pair1A, _ := strconv.Atoi(pair1Tmp[0])
		pair1B, _ := strconv.Atoi(pair1Tmp[1])
		pair2A, _ := strconv.Atoi(pair2Tmp[0])
		pair2B, _ := strconv.Atoi(pair2Tmp[1])
		pair1 := Pair{a: pair1A, b: pair1B}
		pair2 := Pair{a: pair2A, b: pair2B}
		tmpPairSlice := []Pair{pair1, pair2}
		inputPairs = append(inputPairs, tmpPairSlice)
	}

	var total int
	for _, elfPair := range inputPairs {
		if elfPair[0].a >= elfPair[1].a && elfPair[0].b <= elfPair[1].b {
			total++
		} else if elfPair[1].a >= elfPair[0].a && elfPair[1].b <= elfPair[0].b {
			total++
		} else if elfPair[0].a >= elfPair[1].a && elfPair[0].a <= elfPair[1].b {
			total++
		} else if elfPair[0].b >= elfPair[1].a && elfPair[0].b <= elfPair[1].b {
			total++
		} else if elfPair[1].a >= elfPair[0].a && elfPair[1].a <= elfPair[0].b {
			total++
		} else if elfPair[1].b >= elfPair[0].a && elfPair[1].b <= elfPair[0].b {
			total++
		}

	}

	fmt.Println(total)
}
