package main

import (
	"fmt"
	"math"
	"strings"

	intcode "github.com/RichardGejji/AdventOfCode2019/Intcode"
)

type point struct {
	x int64
	y int64
}

var gridSize = int64(500)

var grid [][]rune

func readGrid() {
	var output int64
	grid = [][]rune{}

	//Run code
	for i := int64(0); i < gridSize; i++ {
		row := make([]rune, gridSize, gridSize)
		for j := int64(0); j < gridSize; j++ {
			A := intcode.Computer{}
			A.Initialize(inputStr)
			A.Input = make(chan int64, 2)
			A.Output = make(chan int64, 2)
			go func() {
				A.Run()
			}()
			A.Input <- j
			A.Input <- i
			output = <-A.Output
			if output == 1 {
				row[j] = '#'
			} else {
				row[j] = '.'
			}
		}
		grid = append(grid, row)
	}
}
func readGridTest() {
	grid = [][]rune{}
	for _, row := range strings.Split(testGrid, "\n") {
		grid = append(grid, []rune(row))
	}
}

func printGrid() {
	for i := 0; i < len(grid); i++ {
		fmt.Printf("%s\n", string(grid[i]))
	}
}

//get the average vector value for tractor
func findVectorsFromGrid() (float64, float64) {
	bestSlopeU := float64(0)
	bestSlopeV := float64(-999999999)
	for i := 10; i < len(grid); i++ {
		for j := 1; j < len(grid); j++ {
			if grid[i][j] != '.' {
				slope := -float64(i) / float64(j)
				//fmt.Printf("Found slope %v at j=%d, i=%d \n", slope, j, i)
				if slope < bestSlopeU {
					bestSlopeU = slope
				}
				break
			}
		}

		//Dont calculate # is on RHS
		if grid[i][len(grid[i])-1] != '.' {
			continue
		}

		for j := len(grid[i]) - 1; j > 0; j-- {
			if grid[i][j] != '.' {
				slope := -float64(i) / float64(j)
				if slope > bestSlopeV {
					bestSlopeV = slope
				}
				break
			}

		}

	}
	return bestSlopeU, bestSlopeV
}

func abs(a float64) float64 {
	if a >= 0 {
		return a
	}
	return -a
}

func findGrid(uSlope, vSlope float64, gridLength float64) {
	for ux := float64(10); ux < 1000; ux++ {
		uy := math.Ceil(ux * uSlope)

		//Check if grid within two slopes
		vx := ux + gridLength - 1
		vy := math.Floor(vx * vSlope)
		if uy+gridLength-1 <= vy {
			fmt.Printf("Found grid at x=%v\n", ux)
			fmt.Printf("Upper point is at x=%v, y=%v\n", ux, vy)
			fmt.Printf("Score is: %v\n", 10000*ux-vy)
			return
		}
	}

}

func main() {
	//Fill out grid
	readGrid()
	//readGridTest()
	gridSize = int64(len(grid))
	printGrid()
	slopeU, slopeV := findVectorsFromGrid()
	fmt.Printf("Found valued %v %v\n", slopeU, slopeV)
	//findGrid(slopeU, slopeV, float64(10))
	findGrid(slopeU, slopeV, float64(100))

}

const testGrid = `#.......................................
.#......................................
..##....................................
...###..................................
....###.................................
.....####...............................
......#####.............................
......######............................
.......#######..........................
........########........................
.........#########......................
..........#########.....................
...........##########...................
...........############.................
............############................
.............#############..............
..............##############............
...............###############..........
................###############.........
................#################.......
.................########OOOOOOOOOO.....
..................#######OOOOOOOOOO#....
...................######OOOOOOOOOO###..
....................#####OOOOOOOOOO#####
.....................####OOOOOOOOOO#####
.....................####OOOOOOOOOO#####
......................###OOOOOOOOOO#####
.......................##OOOOOOOOOO#####
........................#OOOOOOOOOO#####
.........................OOOOOOOOOO#####
..........................##############
..........................##############
...........................#############
............................############
.............................###########`

const inputStr = `109,424,203,1,21102,1,11,0,1106,0,282,21101,18,0,0,1105,1,259,2102,1,1,221,203,1,21101,0,31,0,1105,1,282,21101,0,38,0,1105,1,259,20101,0,23,2,21202,1,1,3,21101,0,1,1,21101,57,0,0,1105,1,303,1202,1,1,222,21002,221,1,3,21001,221,0,2,21102,1,259,1,21101,0,80,0,1105,1,225,21101,0,175,2,21102,1,91,0,1106,0,303,2101,0,1,223,21001,222,0,4,21102,259,1,3,21101,225,0,2,21102,1,225,1,21102,1,118,0,1105,1,225,21002,222,1,3,21101,70,0,2,21101,0,133,0,1105,1,303,21202,1,-1,1,22001,223,1,1,21102,1,148,0,1105,1,259,2102,1,1,223,21002,221,1,4,21002,222,1,3,21102,24,1,2,1001,132,-2,224,1002,224,2,224,1001,224,3,224,1002,132,-1,132,1,224,132,224,21001,224,1,1,21101,195,0,0,105,1,109,20207,1,223,2,21002,23,1,1,21101,0,-1,3,21102,1,214,0,1106,0,303,22101,1,1,1,204,1,99,0,0,0,0,109,5,2102,1,-4,249,21202,-3,1,1,22102,1,-2,2,21201,-1,0,3,21101,0,250,0,1106,0,225,21201,1,0,-4,109,-5,2105,1,0,109,3,22107,0,-2,-1,21202,-1,2,-1,21201,-1,-1,-1,22202,-1,-2,-2,109,-3,2105,1,0,109,3,21207,-2,0,-1,1206,-1,294,104,0,99,21202,-2,1,-2,109,-3,2106,0,0,109,5,22207,-3,-4,-1,1206,-1,346,22201,-4,-3,-4,21202,-3,-1,-1,22201,-4,-1,2,21202,2,-1,-1,22201,-4,-1,1,22101,0,-2,3,21101,343,0,0,1105,1,303,1105,1,415,22207,-2,-3,-1,1206,-1,387,22201,-3,-2,-3,21202,-2,-1,-1,22201,-3,-1,3,21202,3,-1,-1,22201,-3,-1,2,21201,-4,0,1,21101,0,384,0,1105,1,303,1105,1,415,21202,-4,-1,-4,22201,-4,-3,-4,22202,-3,-2,-2,22202,-2,-4,-4,22202,-3,-2,-3,21202,-4,-1,-2,22201,-3,-2,1,21201,1,0,-4,109,-5,2106,0,0`
