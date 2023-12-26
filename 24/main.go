package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/go-gl/mathgl/mgl64"
)

type Hailstone struct {
	pos, vel mgl64.Vec3
}

type HailstoneInt struct {
	pos, vel [3]int
}

func main() {
	fileContent, err := ioutil.ReadFile("input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileContent)
	lines := strings.Split(text, "\n")

	hailstones := make([]Hailstone, 0)
	hailstones2 := make([]HailstoneInt, 0)
	for _, line := range lines {
		//println(line)
		split := strings.Split(line, "@")
		posString := split[0]
		velString := split[1]

		hailstones = append(hailstones, Hailstone{str2Vec3(posString), str2Vec3(velString)})
		hailstones2 = append(hailstones2, HailstoneInt{str2int3(posString), str2int3(velString)})
	}

	part1(hailstones)
	part2(hailstones2)
}

func part1(hailstones []Hailstone) {
	sum := 0
	minRange, maxRange := 200000000000000, 400000000000000
	for i, h1 := range hailstones {
		for j := i + 1; j < len(hailstones); j++ {
			h2 := hailstones[j]
			if i == j {
				continue
			}
			p1 := mgl64.Vec2{h1.pos[0], h1.pos[1]}
			p2 := mgl64.Vec2{h1.pos[0] + h1.vel[0], h1.pos[1] + h1.vel[1]}
			p3 := mgl64.Vec2{h2.pos[0], h2.pos[1]}
			p4 := mgl64.Vec2{h2.pos[0] + h2.vel[0], h2.pos[1] + h2.vel[1]}
			intersects, p := lineLine2d(p1, p2, p3, p4)
			if intersects && p.X() >= float64(minRange) && p.X() <= float64(maxRange) && p.Y() >= float64(minRange) && p.Y() <= float64(maxRange) {
				if p.Sub(p1).Dot(p2.Sub(p1)) >= 0 && p.Sub(p3).Dot(p4.Sub(p3)) >= 0 {
					sum++
				}
			}
		}
	}

	println("part1:", sum)
}

func str2Vec3(s string) mgl64.Vec3 {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(strings.TrimSpace(split[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(split[1]))
	z, _ := strconv.Atoi(strings.TrimSpace(split[2]))
	return mgl64.Vec3{float64(x), float64(y), float64(z)}
}

func str2int3(s string) [3]int {
	split := strings.Split(s, ",")
	x, _ := strconv.Atoi(strings.TrimSpace(split[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(split[1]))
	z, _ := strconv.Atoi(strings.TrimSpace(split[2]))
	return [3]int{(x), (y), (z)}
}

func standardForm(p1, p2 mgl64.Vec2) (float64, float64) {
	dyba := p2.Y() - p1.Y()
	dxba := p2.X() - p1.X()
	a := dyba / dxba
	b := -p1.X()*a + p1.Y()
	return a, b
}

func lineLine2d(p1, p2, p3, p4 mgl64.Vec2) (bool, mgl64.Vec2) {
	a1, b1 := standardForm(p1, p2)
	a2, b2 := standardForm(p3, p4)
	if a1*b2-a2*b1 == 0 {
		return false, mgl64.Vec2{}
	}
	cx := a1 - a2
	cy := b2 - b1
	x := cy / cx
	y := a1*x + b1
	return true, mgl64.Vec2{x, y}
}

func part2(hailstones []HailstoneInt) {
	// Generate SageMath script
	fmt.Println()
	fmt.Println("var('x y z vx vy vz t1 t2 t3')")
	fmt.Println("eq1 = x + (vx * t1) == ", hailstones[0].pos[0], " + (", hailstones[0].vel[0], " * t1)")
	fmt.Println("eq2 = y + (vy * t1) == ", hailstones[0].pos[1], " + (", hailstones[0].vel[1], " * t1)")
	fmt.Println("eq3 = z + (vz * t1) == ", hailstones[0].pos[2], " + (", hailstones[0].vel[2], " * t1)")
	fmt.Println("eq4 = x + (vx * t2) == ", hailstones[1].pos[0], " + (", hailstones[1].vel[0], " * t2)")
	fmt.Println("eq5 = y + (vy * t2) == ", hailstones[1].pos[1], " + (", hailstones[1].vel[1], " * t2)")
	fmt.Println("eq6 = z + (vz * t2) == ", hailstones[1].pos[2], " + (", hailstones[1].vel[2], " * t2)")
	fmt.Println("eq7 = x + (vx * t3) == ", hailstones[2].pos[0], " + (", hailstones[2].vel[0], " * t3)")
	fmt.Println("eq8 = y + (vy * t3) == ", hailstones[2].pos[1], " + (", hailstones[2].vel[1], " * t3)")
	fmt.Println("eq9 = z + (vz * t3) == ", hailstones[2].pos[2], " + (", hailstones[2].vel[2], " * t3)")
	fmt.Println("print(solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9],x,y,z,vx,vy,vz,t1,t2,t3))")
	fmt.Println()

	// another approach: turn into linear equations and solve with gaussian elimination. but using
	// github.com/alex-ant/gomath did only work for example and failed for real input cause of precision
	// issues, so rechecked with sage math and this worked. So principally this should work with solving
	// by simply doing gaussian elimination
	getCoeffs := func(h0, h1 HailstoneInt, idx0, idx1 int) (int, int, int, int, int) {
		hx0, hy0 := h0.pos[idx0], h0.pos[idx1]
		hvx0, hvy0 := h0.vel[idx0], h0.vel[idx1]
		hx1, hy1 := h1.pos[idx0], h1.pos[idx1]
		hvx1, hvy1 := h1.vel[idx0], h1.vel[idx1]
		rhs := -hx0*hvy0 + hy0*hvx0 + hx1*hvy1 - hy1*hvx1
		//return coefficients for [x y vx vy rhs] or [y z vy vz rhs] with idx0=1, idx1=2
		return (hvy1 - hvy0), (hvx0 - hvx1), (hy0 - hy1), (hx1 - hx0), (rhs)
	}

	a0, b0, c0, d0, rhs0 := getCoeffs(hailstones[0], hailstones[1], 0, 1)
	a1, b1, c1, d1, rhs1 := getCoeffs(hailstones[0], hailstones[2], 0, 1)
	a2, b2, c2, d2, rhs2 := getCoeffs(hailstones[0], hailstones[3], 0, 1)
	a3, b3, c3, d3, rhs3 := getCoeffs(hailstones[0], hailstones[4], 0, 1)
	a4, b4, c4, d4, rhs4 := getCoeffs(hailstones[1], hailstones[2], 1, 2)
	a5, b5, c5, d5, rhs5 := getCoeffs(hailstones[1], hailstones[3], 1, 2)

	fmt.Println()
	fmt.Println("var('x y z vx vy vz')")
	fmt.Printf("eq1 = (%v * x) + (%v * y) + (%v * vx) + (%v * vy) == %v\n", a0, b0, c0, d0, rhs0)
	fmt.Printf("eq2 = (%v * x) + (%v * y) + (%v * vx) + (%v * vy) == %v\n", a1, b1, c1, d1, rhs1)
	fmt.Printf("eq3 = (%v * x) + (%v * y) + (%v * vx) + (%v * vy) == %v\n", a2, b2, c2, d2, rhs2)
	fmt.Printf("eq4 = (%v * x) + (%v * y) + (%v * vx) + (%v * vy) == %v\n", a3, b3, c3, d3, rhs3)
	fmt.Printf("eq5 = (%v * y) + (%v * z) + (%v * vy) + (%v * vz) == %v\n", a4, b4, c4, d4, rhs4)
	fmt.Printf("eq6 = (%v * y) + (%v * z) + (%v * vy) + (%v * vz) == %v\n", a5, b5, c5, d5, rhs5)
	fmt.Println("print(solve([eq1,eq2,eq3,eq4,eq5,eq6],x,y,z,vx,vy,vz))")
	fmt.Println()

	fmt.Println("Part 2: ", 331109811422259+312547020340291+118035075297081) //761691907059631
}
