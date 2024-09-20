package utils

import "math"

type Vec3 struct {
	X, Y, Z float64
}

type Vec2 struct {
	X, Y float64
}

func Vec3ToVec2(v Vec3) Vec2 {
	return Vec2{X: v.X, Y: v.Y}
}

func Subtract3D(v1, v2 Vec3) Vec3 {
	return Vec3{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
}

func Cross3D(v1, v2 Vec3) Vec3 {
	return Vec3{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
}

func Dot3D(v1, v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func Subtract2D(v1, v2 Vec2) Vec2 {
	return Vec2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func Cross2D(v1, v2 Vec2) float64 {
	return v1.X*v2.Y - v1.Y*v2.X
}

func LinesIntersect2D(p1, d1, p2, d2 Vec2) (bool, Vec2) {
	crossDir := Cross2D(d1, d2)
	if math.Abs(crossDir) < 1e-6 {
		return false, Vec2{}
	}

	p1p2 := Subtract2D(p2, p1)

	t1 := Cross2D(p1p2, d2) / crossDir

	intersectPoint := Vec2{
		X: p1.X + t1*d1.X,
		Y: p1.Y + t1*d1.Y,
	}

	return true, intersectPoint
}
