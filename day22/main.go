package main

import (
	"math"
	"regexp"

	"github.com/varbrad/advent-of-code-2021/utils"
)

func main() {
	input := utils.ReadInputToList("day22/input")
	utils.Day(22, "Reactor Reboot").Input(input).Part1(Day22Part1).Part2(Day22Part2)
}

func Day22Part1(input []string) int {
	return solve(input, true)
}

func Day22Part2(input []string) int {
	return solve(input, false)
}

func solve(input []string, boundCheck bool) int {
	cuboids := map[cuboid]int{}

	for _, line := range input {
		ins := parseLine(line)
		if boundCheck && !ins.inBounds {
			continue
		}
		newCuboids := map[cuboid]int{}
		for key, value := range cuboids {
			ok, subCuboid := makeSubCuboid(key, ins.cuboid)
			if ok {
				newCuboids[subCuboid] -= value
			}
		}
		if ins.mode == "on" {
			newCuboids[ins.cuboid]++
		}
		for k, v := range newCuboids {
			cuboids[k] += v
		}
	}

	sum := 0
	for key, sgn := range cuboids {
		sum += key.volume() * sgn
	}

	return sum
}

type vec3d struct {
	x int
	y int
	z int
}

type cuboid struct {
	start vec3d
	end   vec3d
}

func makeCuboid(x0, x1, y0, y1, z0, z1 int) cuboid {
	return cuboid{
		vec3d{x0, y0, z0},
		vec3d{x1, y1, z1},
	}
}

func makeSubCuboid(key cuboid, new cuboid) (ok bool, c cuboid) {
	ix0 := utils.MaxInteger(new.start.x, key.start.x)
	ix1 := utils.MinInteger(new.end.x, key.end.x)
	iy0 := utils.MaxInteger(new.start.y, key.start.y)
	iy1 := utils.MinInteger(new.end.y, key.end.y)
	iz0 := utils.MaxInteger(new.start.z, key.start.z)
	iz1 := utils.MinInteger(new.end.z, key.end.z)

	if ix0 <= ix1 && iy0 <= iy1 && iz0 <= iz1 {
		return true, makeCuboid(ix0, ix1, iy0, iy1, iz0, iz1)
	}

	return false, c
}

func (c cuboid) volume() int {
	dx := c.end.x - c.start.x + 1
	dy := c.end.y - c.start.y + 1
	dz := c.end.z - c.start.z + 1

	return dx * dy * dz
}

type instruction struct {
	mode     string
	inBounds bool
	cuboid   cuboid
}

var regex = regexp.MustCompile(`^(on|off) x=(-?\d+)..(-?\d+),y=(-?\d+)..(-?\d+),z=(-?\d+)..(-?\d+)$`)

func parseLine(line string) instruction {
	result := regex.FindStringSubmatch(line)

	c := makeCuboid(
		utils.ToInteger(result[2]),
		utils.ToInteger(result[3]),
		utils.ToInteger(result[4]),
		utils.ToInteger(result[5]),
		utils.ToInteger(result[6]),
		utils.ToInteger(result[7]),
	)

	return instruction{
		mode:     result[1],
		inBounds: isInBounds(c),
		cuboid:   c,
	}
}

func isInBounds(c cuboid) bool {
	for _, vec := range []vec3d{c.start, c.end} {
		for _, v := range []int{vec.x, vec.y, vec.z} {
			if math.Abs(float64(v)) > 50 {
				return false
			}
		}
	}
	return true
}
