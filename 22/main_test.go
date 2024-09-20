package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestCreateBrick(t *testing.T) {
	brick := createBrick("2,0,2~1,2,3")

	assert.Equal(t, brick, &Brick{y: 2, height: 2, x: 1, width: 2, z: 0, depth: 3, supports: make([]*Brick, 0), isSupportedBy: make([]*Brick, 0)})
}

func TestSortBricksByHeight(t *testing.T) {
	bricks := []*Brick{{y: 4}, {y: 2}, {y: 6}, {y: 1}}

	sortBricksByHeight(bricks)

	assert.Equal(t, bricks[0].y, 1)
	assert.Equal(t, bricks[1].y, 2)
	assert.Equal(t, bricks[2].y, 4)
	assert.Equal(t, bricks[3].y, 6)
}

func TestIsColliding(t *testing.T) {
	brick1 := createBrick("1,0,1~1,2,1")
	brick2 := createBrick("0,0,2~2,0,2")

	assert.Equal(t, areBricksColliding(brick1, brick2), false)

	brick2 = createBrick("0,0,1~2,0,1")

	assert.Equal(t, areBricksColliding(brick1, brick2), true)
}
