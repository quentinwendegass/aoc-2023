package main

import (
	"aoc2023/utils"
	"strings"
)

type Line struct {
	pos utils.Vec3
	vel utils.Vec3
}

type BoundingBox struct {
	minX, minY, maxX, maxY float64
}

func main() {
	utils.WithAOC(partOne, partTwo, utils.DefaultDataLoader)
}

func partOne(data []byte) int {
	lines := convertInputToLines(data)

	box := BoundingBox{
		minX: 200000000000000,
		minY: 200000000000000,
		maxX: 400000000000000,
		maxY: 400000000000000,
	}

	return countIntersectionsInArea(lines, box)
}

// Solved with the help of reddit. Don't like this one since it doesn't work on the example input.
func partTwo(data []byte) int {
	lines := convertInputToLines(data)

	rockVel := findVelocities(lines)
	rockPos := calculateRockPosition(lines[0], lines[1], rockVel)

	return int(rockPos.X) + int(rockPos.Y) + int(rockPos.Z)
}

func convertInputToLines(data []byte) []Line {
	lineStrings := strings.Split(string(data), "\n")

	lines := make([]Line, len(lineStrings))

	for i, line := range lineStrings {
		l := strings.Split(line, "@")

		positions := strings.Split(l[0], ",")
		velocities := strings.Split(l[1], ",")

		posX, posY, posZ := utils.ToInt(strings.TrimSpace(positions[0])), utils.ToInt(strings.TrimSpace(positions[1])), utils.ToInt(strings.TrimSpace(positions[2]))
		velX, velY, velZ := utils.ToInt(strings.TrimSpace(velocities[0])), utils.ToInt(strings.TrimSpace(velocities[1])), utils.ToInt(strings.TrimSpace(velocities[2]))

		lines[i] = Line{pos: utils.Vec3{X: float64(posX), Y: float64(posY), Z: float64(posZ)}, vel: utils.Vec3{X: float64(velX), Y: float64(velY), Z: float64(velZ)}}
	}

	return lines
}

func calculateRockPosition(l1, l2 Line, rockVel utils.Vec3) utils.Vec3 {
	rockPos := utils.Vec3{}

	mA := (l1.vel.Y - rockVel.Y) / (l1.vel.X - rockVel.X)
	mB := (l2.vel.Y - rockVel.Y) / (l2.vel.X - rockVel.X)
	cA := l1.pos.Y - (mA * l1.pos.X)
	cB := l2.pos.Y - (mB * l2.pos.X)
	rockPos.X = (cB - cA) / (mA - mB)
	rockPos.Y = mA*rockPos.X + cA
	time := (rockPos.X - l1.pos.X) / (l1.vel.X - rockVel.X)
	rockPos.Z = l1.pos.Z + (l1.vel.Z-rockVel.Z)*time

	return rockPos
}

func findVelocities(lines []Line) utils.Vec3 {
	possibleXVelocities, possibleYVelocities, possibleZVelocities := make([]int, 0), make([]int, 0), make([]int, 0)

	for _, line1 := range lines {
		for _, line2 := range lines {
			if line1 == line2 {
				continue
			}

			possibleXVelocities = findPossibleVelocities(int(line1.pos.X), int(line2.pos.X), int(line1.vel.X), int(line2.vel.X), possibleXVelocities)
			possibleYVelocities = findPossibleVelocities(int(line1.pos.Y), int(line2.pos.Y), int(line1.vel.Y), int(line2.vel.Y), possibleYVelocities)
			possibleZVelocities = findPossibleVelocities(int(line1.pos.Z), int(line2.pos.Z), int(line1.vel.Z), int(line2.vel.Z), possibleZVelocities)
		}
	}

	if len(possibleXVelocities) != 1 || len(possibleYVelocities) != 1 || len(possibleZVelocities) != 1 {
		panic("invalid input data")
	}

	return utils.Vec3{X: float64(possibleXVelocities[0]), Y: float64(possibleYVelocities[0]), Z: float64(possibleZVelocities[0])}
}

func findPossibleVelocities(pos1, pos2, vel1, vel2 int, allPossibleVelocities []int) []int {
	if vel1 == vel2 {
		possibleVelocities := make([]int, 0)
		for i := -500; i <= 500; i++ {
			if i != vel1 && (pos1-pos2)%(i-vel1) == 0 {
				possibleVelocities = append(possibleVelocities, i)
			}
		}

		if len(allPossibleVelocities) == 0 {
			return possibleVelocities
		} else {
			return utils.FindCommonElements(allPossibleVelocities, possibleVelocities)
		}
	}

	return allPossibleVelocities
}

func isPointInBoundingBox(point utils.Vec2, box BoundingBox) bool {
	return point.X >= box.minX && point.X <= box.maxX &&
		point.Y >= box.minY && point.Y <= box.maxY
}

func isPointInPast(point utils.Vec2, line1, line2 Line) bool {
	return (line1.vel.X > 0 && point.X < line1.pos.X) || (line1.vel.X <= 0 && point.X > line1.pos.X) ||
		(line1.vel.Y > 0 && point.Y < line1.pos.Y) || (line1.vel.Y <= 0 && point.Y > line1.pos.Y) ||
		(line2.vel.X > 0 && point.X < line2.pos.X) || (line2.vel.X <= 0 && point.X > line2.pos.X) ||
		(line2.vel.Y > 0 && point.Y < line2.pos.Y) || (line2.vel.Y <= 0 && point.Y > line2.pos.Y)
}

func countIntersectionsInArea(lines []Line, box BoundingBox) int {
	count := 0
	seenPairs := make(map[[2]int]bool)

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if i == j {
				continue
			}
			if seenPairs[[2]int{i, j}] || seenPairs[[2]int{j, i}] {
				continue
			}

			intersect, point := utils.LinesIntersect2D(utils.Vec3ToVec2(lines[i].pos), utils.Vec3ToVec2(lines[i].vel), utils.Vec3ToVec2(lines[j].pos), utils.Vec3ToVec2(lines[j].vel))

			if intersect && isPointInBoundingBox(point, box) && !isPointInPast(point, lines[i], lines[j]) {
				count++
				seenPairs[[2]int{i, j}] = true
			}
		}
	}

	return count
}
