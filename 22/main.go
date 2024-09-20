package main

import (
	"aoc2023/utils"
	"slices"
	"strconv"
	"strings"
)

type Brick struct {
	x, y, z              int
	width, height, depth int
	supports             []*Brick
	isSupportedBy        []*Brick
}

// Loved this one!
func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	bricks := createBricksFromInput(data)
	simulateBricksFalling(bricks)

	removableBricks, _ := splitBricksByRemovable(bricks)

	return len(removableBricks)
}

func partTwo(data []byte) int {
	bricks := createBricksFromInput(data)
	simulateBricksFalling(bricks)
	_, unremovableBricks := splitBricksByRemovable(bricks)

	totalFallingBricks := 0

	for _, brick := range unremovableBricks {
		queue := utils.CreateQueue[*Brick]()
		queue.Push(brick.supports...)

		removedBricks := []*Brick{brick}

		for queue.Len() > 0 {
			topBrick := queue.Pop()

			if !hasBrickSupport(topBrick, removedBricks) {
				totalFallingBricks++
				removedBricks = append(removedBricks, topBrick)
				queue.PushNoDuplicate(topBrick.supports...)
			}
		}
	}

	return totalFallingBricks
}

func hasBrickSupport(brick *Brick, removedBricks []*Brick) bool {
	for _, supportedBy := range brick.isSupportedBy {
		if !slices.Contains(removedBricks, supportedBy) {
			return true
		}
	}
	return false
}

func splitBricksByRemovable(bricks []*Brick) (removable, unremovable []*Brick) {
	removableBricks := make([]*Brick, 0)
	unremovableBricks := make([]*Brick, 0)

	for _, brick := range bricks {
		canRemove := true
		for _, supportedBrick := range brick.supports {
			if len(supportedBrick.isSupportedBy) == 1 {
				canRemove = false
				break
			}
		}

		if canRemove {
			removableBricks = append(removableBricks, brick)
		} else {
			unremovableBricks = append(unremovableBricks, brick)
		}
	}

	return removableBricks, unremovableBricks
}

func simulateBricksFalling(bricks []*Brick) {
	maxHeight := 0
	fallenBricks := make([]*Brick, 0)

	sortBricksByHeight(bricks)

	for _, brick := range bricks {
		if brick.y > maxHeight {
			brick.y = maxHeight
		}

		collision := false

		for !collision {
			if brick.y == 0 {
				break
			}

			brick.y--

			for _, fallenBrick := range fallenBricks {
				if areBricksColliding(brick, fallenBrick) {
					collision = true
					brick.isSupportedBy = append(brick.isSupportedBy, fallenBrick)
					fallenBrick.supports = append(fallenBrick.supports, brick)
				}
			}

			if collision {
				brick.y++
			}
		}

		brickHightPoint := brick.height + brick.y
		if brickHightPoint > maxHeight {
			maxHeight = brickHightPoint
		}
		fallenBricks = append(fallenBricks, brick)
	}
}

func createBricksFromInput(data []byte) []*Brick {
	brickInputs := strings.Split(string(data), "\n")

	bricks := make([]*Brick, len(brickInputs))

	for i, input := range brickInputs {
		bricks[i] = createBrick(input)
	}

	return bricks
}

func createBrick(data string) *Brick {
	coordinates := strings.Split(data, "~")

	position1 := strings.Split(coordinates[0], ",")
	position2 := strings.Split(coordinates[1], ",")

	x1 := atoi(position1[0])
	y1 := atoi(position1[2])
	z1 := atoi(position1[1])

	x2 := atoi(position2[0])
	y2 := atoi(position2[2])
	z2 := atoi(position2[1])

	x := min(x1, x2)
	y := min(y1, y2)
	z := min(z1, z2)
	width := max(x1, x2) - x + 1
	height := max(y1, y2) - y + 1
	depth := max(z1, z2) - z + 1

	supports := make([]*Brick, 0)
	isSupportedBy := make([]*Brick, 0)

	return &Brick{x: x, y: y, z: z, width: width, height: height, depth: depth, supports: supports, isSupportedBy: isSupportedBy}
}

func sortBricksByHeight(bricks []*Brick) {
	slices.SortStableFunc(bricks, func(b1, b2 *Brick) int {
		if b1.y < b2.y {
			return -1
		} else if b1.y > b2.y {
			return 1
		}
		return 0
	})
}

func areBricksColliding(brick1, brick2 *Brick) bool {
	xCollision := brick1.x < brick2.x+brick2.width && brick1.x+brick1.width > brick2.x
	yCollision := brick1.y < brick2.y+brick2.height && brick1.y+brick1.height > brick2.y
	zCollision := brick1.z < brick2.z+brick2.depth && brick1.z+brick1.depth > brick2.z

	return xCollision && yCollision && zCollision
}

func atoi(str string) int {
	val, err := strconv.Atoi(str)

	if err != nil {
		panic("input data is in the wrong format")
	}

	return val
}
