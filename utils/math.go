package utils

import "math"

func GCD(a int, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM(values []int) int {
	lcm := values[0]

	for i := 1; i < len(values); i++ {
		lcm = values[i] * lcm / GCD(values[i], lcm)
	}

	return lcm
}

func ShoelaceArea(vertices [][]int) float64 {
	n := len(vertices)
	area := 0.0
	for i := 0; i < n; i++ {
		x1, y1 := vertices[i][0], vertices[i][1]
		x2, y2 := vertices[(i+1)%n][0], vertices[(i+1)%n][1]
		area += float64(x1*y2 - x2*y1)
	}
	return math.Abs(area) / 2.0
}

func CountBoundaryPoints(vertices [][]int) int {
	n := len(vertices)
	boundaryPoints := 0
	for i := 0; i < n; i++ {
		x1, y1 := vertices[i][0], vertices[i][1]
		x2, y2 := vertices[(i+1)%n][0], vertices[(i+1)%n][1]
		dx := int(math.Abs(float64(x2 - x1)))
		dy := int(math.Abs(float64(y2 - y1)))
		boundaryPoints += GCD(dx, dy)
	}
	return boundaryPoints
}

func CountPointsInsideArea(area float64, boundaryPoints int) int {
	// Pick's theorem
	return int(area - float64(boundaryPoints)/2 + 1)
}

func Multiply(values ...int) int {
	total := 1
	for _, val := range values {
		total *= val
	}
	return total
}
