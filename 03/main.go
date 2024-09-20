package main

import (
	"aoc2023/utils"
	"bytes"
)

type Gear struct {
	x int
	y int
}

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

type Number struct {
	num [3]byte
	x   int
	y   int
	len int
}

func partOne(data []byte) int {
	lines := bytes.Split(data, []byte("\n"))

	num := [3]byte{}
	numIdx := 0

	nums := [10000]Number{}
	numsIdx := 0

	symbols := make([][]bool, len(lines[0]))
	for i := range symbols {
		symbols[i] = make([]bool, len(lines))
	}

	storeNum := func(x int, y int) {
		nums[numsIdx] = Number{num: num, x: x, y: y, len: numIdx}
		num[0] = 0
		num[1] = 0
		num[2] = 0
		numIdx = 0
		numsIdx++
	}

	for idy, line := range lines {
		for idx, char := range line {
			if char >= '0' && char <= '9' {
				num[numIdx] = char
				numIdx++
				continue
			} else if numIdx != 0 {
				storeNum(idx-1, idy)
			}

			if char != '.' {
				symbols[idx][idy] = true
			}
		}

		if numIdx != 0 {
			storeNum(len(lines[0])-1, idy)
		}
	}

	total := 0

	alignsWithSymbol := func(num Number) bool {
		if num.x != len(lines[0])-1 && symbols[num.x+1][num.y] {
			return true
		} else if num.x-num.len >= 0 && symbols[num.x-num.len][num.y] {
			return true
		}

		checkBounds := func(y int) bool {
			for x := num.x + 1; x >= num.x-num.len; x-- {
				if x >= len(lines[0]) || x < 0 {
					continue
				}

				if symbols[x][y] {
					return true
				}
			}

			return false
		}

		if num.y > 0 && checkBounds(num.y-1) {
			return true
		}

		if num.y < len(lines)-1 && checkBounds(num.y+1) {
			return true
		}

		return false
	}

	for _, num := range nums {
		if num.len == 0 {
			break
		}

		if alignsWithSymbol(num) {
			total += convertToInt(num.num)
		}
	}

	return total
}

func partTwo(data []byte) int {
	lines := bytes.Split(data, []byte("\n"))

	num := [3]byte{}
	numIdx := 0

	gears := [1000]Gear{}
	gearsIdx := 0

	nums := make([][]*[3]byte, len(lines[0]))
	for i := range nums {
		nums[i] = make([]*[3]byte, len(lines))
	}

	storeNum := func(x int, y int) {
		copyNum := num
		pNum := &copyNum

		for ix := x; ix >= x-numIdx+1; ix-- {
			nums[ix][y] = pNum
		}

		num[0] = 0
		num[1] = 0
		num[2] = 0
		numIdx = 0
	}

	for idy, line := range lines {
		for idx, char := range line {
			if char >= '0' && char <= '9' {
				num[numIdx] = char
				numIdx++
				continue
			} else if numIdx != 0 {
				storeNum(idx-1, idy)
			}

			if char == '*' {
				gears[gearsIdx] = Gear{x: idx, y: idy}
				gearsIdx++
			}
		}

		if numIdx != 0 {
			storeNum(len(lines[0])-1, idy)
		}
	}

	total := 0

	for i, gear := range gears {
		if i > gearsIdx {
			break
		}

		nnums := [6][3]byte{}
		nnumsIdx := 0

		if gear.x != len(lines[0])-1 && nums[gear.x+1][gear.y] != nil {
			nnums[nnumsIdx] = *nums[gear.x+1][gear.y]
			nnumsIdx++
		}

		if gear.x > 0 && nums[gear.x-1][gear.y] != nil {
			nnums[nnumsIdx] = *nums[gear.x-1][gear.y]
			nnumsIdx++
		}

		if gear.y > 0 {
			if nums[gear.x][gear.y-1] != nil {
				nnums[nnumsIdx] = *nums[gear.x][gear.y-1]
				nnumsIdx++
			} else {
				if gear.x != len(lines[0])-1 && nums[gear.x+1][gear.y-1] != nil {
					nnums[nnumsIdx] = *nums[gear.x+1][gear.y-1]
					nnumsIdx++
				}

				if gear.x > 0 && nums[gear.x-1][gear.y-1] != nil {
					nnums[nnumsIdx] = *nums[gear.x-1][gear.y-1]
					nnumsIdx++
				}
			}
		}

		if gear.y < len(lines)-1 {
			if nums[gear.x][gear.y+1] != nil {
				nnums[nnumsIdx] = *nums[gear.x][gear.y+1]
				nnumsIdx++
			} else {
				if gear.x != len(lines[0])-1 && nums[gear.x+1][gear.y+1] != nil {
					nnums[nnumsIdx] = *nums[gear.x+1][gear.y+1]
					nnumsIdx++
				}

				if gear.x > 0 && nums[gear.x-1][gear.y+1] != nil {
					nnums[nnumsIdx] = *nums[gear.x-1][gear.y+1]
					nnumsIdx++
				}
			}
		}

		if nnums[1][0] != 0 && nnums[2][0] == 0 {
			total += convertToInt(nnums[0]) * convertToInt(nnums[1])
		}
	}

	return total
}

func convertToInt(r [3]byte) int {
	var num int = 0
	if r[2] != 0 {
		num += int((r[0] - 48)) * 100
		num += int((r[1] - 48)) * 10
		num += int((r[2] - 48))
	} else if r[1] != 0 {
		num += int((r[0] - 48)) * 10
		num += int((r[1] - 48))
	} else {
		num += int((r[0] - 48))
	}
	return num
}
