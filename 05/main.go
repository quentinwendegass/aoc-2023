package main

import (
	"aoc2023/utils"
	"strings"
)

type CategoryRange struct {
	src int
	dst int
	len int
}

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	seeds, categories := transformInput(data)

	minLocation := -1

	for _, seed := range seeds {
		location := seed
		for _, categorieRanges := range categories {
			location = getSrcFromCatagory(location, categorieRanges)
		}

		if minLocation == -1 || location < minLocation {
			minLocation = location
		}

	}

	return minLocation
}

// I couldn't be bothered implementing a range based algorithm to make this fast, so let's just brute force it with multiple threads.
// Takes around ~1min to execute on my old macbook.
func partTwo(data []byte) int {
	seeds, categories := transformInput(data)

	seedCatRanges := make([]CategoryRange, len(seeds)/2)

	for i := 0; i < len(seeds); i += 2 {
		seedCatRanges[i/2] = CategoryRange{src: seeds[i], dst: seeds[i], len: seeds[i+1]}
	}

	minLocations := make(chan int)

	process := func(seed CategoryRange) {
		minLocation := -1
		for i := seed.src; i < seed.src+seed.len; i++ {
			location := i
			for _, categorieRanges := range categories {
				location = getSrcFromCatagory(location, categorieRanges)
			}

			if minLocation == -1 || location < minLocation {
				minLocation = location
			}
		}

		minLocations <- minLocation
	}

	for _, seed := range seedCatRanges {
		go process(seed)
	}

	totalMinLocation := -1

	for i := 0; i < len(seedCatRanges); i++ {
		minLocation := <-minLocations

		if totalMinLocation == -1 || minLocation < totalMinLocation {
			totalMinLocation = minLocation
		}
	}

	return totalMinLocation
}

func getSrcFromCatagory(input int, categorieRanges []CategoryRange) int {
	for _, categoryRange := range categorieRanges {
		if input >= categoryRange.src && input < categoryRange.src+categoryRange.len {
			return categoryRange.dst + (input - categoryRange.src)
		}
	}

	return input
}

func transformInput(data []byte) ([]int, [][]CategoryRange) {
	splitData := strings.Split(string(data), "\n\n")

	seedStr := splitData[0]
	mapsStr := splitData[1:]

	seeds := utils.AToISlice(utils.ValuesFromLabelStringNoLabel(seedStr, ":", " "))
	categories := make([][]CategoryRange, len(mapsStr))

	for idx, mapStr := range mapsStr {
		mapStrLines := strings.Split(mapStr, "\n")[1:]

		categoryRanges := make([]CategoryRange, len(mapStrLines))
		for idy, mapStrLine := range mapStrLines {
			mapRange := utils.AToISlice(utils.ValuesFromString(mapStrLine, " "))

			categoryRanges[idy] = CategoryRange{src: mapRange[1], dst: mapRange[0], len: mapRange[2]}
		}

		categories[idx] = categoryRanges
	}

	return seeds, categories
}
