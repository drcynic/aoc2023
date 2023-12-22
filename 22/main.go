package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

type Pos struct {
	x, y, z int
}

type Brick struct {
	start, end    Pos
	supports      map[int]bool
	isSupportedBy map[int]bool
}

func part1() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	bricks := make([]*Brick, 0)

	for _, line := range lines {
		brick := createBrick(line)
		bricks = append(bricks, &brick)
	}

	sort.Slice(bricks, func(i, j int) bool { return bricks[i].start.z < bricks[j].start.z })

	// lookup map for speedup
	bricksByEndZ := make(map[int][]*Brick)
	for _, brick := range bricks {
		bricksByEndZ[brick.end.z] = append(bricksByEndZ[brick.end.z], brick)
	}

	moved := true
	for moved {
		moved = false
		for i, brick := range bricks {
			if canFall2(bricks, i, &bricksByEndZ) {
				prevBricksZ := bricksByEndZ[brick.end.z]
				brickIdx := slices.Index(prevBricksZ, brick)
				bricksByEndZ[brick.end.z] = slices.Delete(prevBricksZ, brickIdx, brickIdx+1)
				brick.start.z = brick.start.z - 1
				brick.end.z = brick.end.z - 1
				bricksByEndZ[brick.end.z] = append(bricksByEndZ[brick.end.z], brick)
				moved = true
			}
		}
	}

	// find supporting bricks
	for i, brick := range bricks {
		for j, b := range bricks {
			if i == j {
				continue
			}

			overlapInX := ((brick.start.x >= b.start.x && brick.start.x <= b.end.x) ||
				(brick.end.x >= b.start.x && brick.end.x <= b.end.x)) ||
				(brick.start.x <= b.start.x && brick.end.x >= b.end.x)
			if !overlapInX {
				continue
			}

			overlapInY := ((brick.start.y >= b.start.y && brick.start.y <= b.end.y) ||
				(brick.end.y >= b.start.y && brick.end.y <= b.end.y)) ||
				(brick.start.y <= b.start.y && brick.end.y >= b.end.y)
			if !overlapInY {
				continue
			}

			if brick.end.z == b.start.z-1 {
				brick.supports[j] = true
				b.isSupportedBy[i] = true
			}
		}
	}

	numDisintegratable := 0
	for _, brick := range bricks {
		if len(brick.supports) == 0 {
			numDisintegratable++
			continue
		}

		removeOk := true
		for j, _ := range brick.supports {
			if len(bricks[j].isSupportedBy) == 1 {
				removeOk = false
				break
			}
		}
		if removeOk {
			numDisintegratable++
		}
	}

	fmt.Printf("part 1: %v\n", numDisintegratable)
}

func canFall(bricks []Brick, i int) bool {
	brick := bricks[i]
	if brick.start.z == 1 {
		return false
	}
	for j, b := range bricks {
		if i == j {
			continue
		}
		overlapInX := ((brick.start.x >= b.start.x && brick.start.x <= b.end.x) ||
			(brick.end.x >= b.start.x && brick.end.x <= b.end.x)) ||
			(brick.start.x <= b.start.x && brick.end.x >= b.end.x)
		overlapInY := ((brick.start.y >= b.start.y && brick.start.y <= b.end.y) ||
			(brick.end.y >= b.start.y && brick.end.y <= b.end.y)) ||
			(brick.start.y <= b.start.y && brick.end.y >= b.end.y)
		if (overlapInX && overlapInY) && brick.end.z > b.end.z && brick.start.z-b.end.z <= 1 {
			return false
		}
	}
	return true
}

func canFall2(bricks []*Brick, i int, bricksByEndZ *map[int][]*Brick) bool {
	brick := bricks[i]
	if brick.start.z == 1 {
		return false
	}
	// get relevant bricks
	bricksToCheck := (*bricksByEndZ)[brick.start.z-1]

	for _, b := range bricksToCheck {
		if brick == b {
			continue
		}
		overlapInX := ((brick.start.x >= b.start.x && brick.start.x <= b.end.x) ||
			(brick.end.x >= b.start.x && brick.end.x <= b.end.x)) ||
			(brick.start.x <= b.start.x && brick.end.x >= b.end.x)
		overlapInY := ((brick.start.y >= b.start.y && brick.start.y <= b.end.y) ||
			(brick.end.y >= b.start.y && brick.end.y <= b.end.y)) ||
			(brick.start.y <= b.start.y && brick.end.y >= b.end.y)
		if (overlapInX && overlapInY) && brick.end.z > b.end.z && brick.start.z-b.end.z <= 1 {
			return false
		}
	}
	return true
}
func createBrick(brickLine string) Brick {
	brickString := strings.Split(brickLine, "~")
	startString := brickString[0]
	endString := brickString[1]
	startCoords := strings.Split(startString, ",")
	x, _ := strconv.Atoi(startCoords[0])
	y, _ := strconv.Atoi(startCoords[1])
	z, _ := strconv.Atoi(startCoords[2])
	startPos := Pos{x, y, z}

	endCoords := strings.Split(endString, ",")
	x, _ = strconv.Atoi(endCoords[0])
	y, _ = strconv.Atoi(endCoords[1])
	z, _ = strconv.Atoi(endCoords[2])
	endPos := Pos{x, y, z}

	// swap if start is bigger than end
	if startPos.x > endPos.x {
		startPos.x, endPos.x = endPos.x, startPos.x
	}
	if startPos.y > endPos.y {
		startPos.y, endPos.y = endPos.y, startPos.y
	}
	if startPos.z > endPos.z {
		startPos.z, endPos.z = endPos.z, startPos.z
	}

	brick := Brick{startPos, endPos, make(map[int]bool), make(map[int]bool)}
	return brick
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")
	bricks := make([]*Brick, 0)

	for _, line := range lines {
		//println(line)
		brick := createBrick(line)
		bricks = append(bricks, &brick)
	}

	sort.Slice(bricks, func(i, j int) bool { return bricks[i].start.z < bricks[j].start.z })

	bricksByEndZ := make(map[int][]*Brick)
	for _, brick := range bricks {
		bricksByEndZ[brick.end.z] = append(bricksByEndZ[brick.end.z], brick)
	}

	moved := true
	for moved {
		moved = false
		for i, brick := range bricks {
			if canFall2(bricks, i, &bricksByEndZ) {
				prevBricksZ := bricksByEndZ[brick.end.z]
				brickIdx := slices.Index(prevBricksZ, brick)
				bricksByEndZ[brick.end.z] = slices.Delete(prevBricksZ, brickIdx, brickIdx+1)
				brick.start.z = brick.start.z - 1
				brick.end.z = brick.end.z - 1
				bricksByEndZ[brick.end.z] = append(bricksByEndZ[brick.end.z], brick)
				moved = true
			}
		}
	}

	sum := 0
	originalBricks := make([]Brick, len(bricks))
	for i, brick := range bricks {
		originalBricks[i] = *brick
	}
	for idxToRemove, _ := range bricks {
		// restore original state before removing the ith brick
		bricksToCheck := make([]*Brick, len(originalBricks))
		for i, brick := range originalBricks {
			brickCopy := brick
			bricksToCheck[i] = &brickCopy
		}
		bricksToCheck = slices.Delete(bricksToCheck, idxToRemove, idxToRemove+1)

		// lookup map for speedup
		bricksByEndZ = make(map[int][]*Brick)
		for _, brick := range bricksToCheck {
			bricksByEndZ[brick.end.z] = append(bricksByEndZ[brick.end.z], brick)
		}

		foundMovers := map[*Brick]bool{}
		moved = true
		for moved {
			moved = false
			for i, brick := range bricksToCheck {
				if canFall2(bricksToCheck, i, &bricksByEndZ) {
					prevBricksZ := bricksByEndZ[brick.end.z]
					brickIdx := slices.Index(prevBricksZ, brick)
					bricksByEndZ[brick.end.z] = slices.Delete(prevBricksZ, brickIdx, brickIdx+1)
					brick.start.z = brick.start.z - 1
					brick.end.z = brick.end.z - 1
					bricksByEndZ[brick.end.z] = append(bricksByEndZ[brick.end.z], brick)
					moved = true
					foundMovers[brick] = true
				}
			}
		}

		sum += len(foundMovers)
	}

	fmt.Printf("part 2:  %v\n", sum)
}
