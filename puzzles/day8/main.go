package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

const (
	ySize   = 6
	xSize   = 50
	pRect   = "rect "
	pRotRow = "rotate row y="
	pRotCol = "rotate column x="
)

func main() {
	grid := new([ySize][xSize]bool)
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		parseAndExec(line, grid)
	}
	count := printAndCountGrid(grid)
	fmt.Println(count)
}

func parseAndExec(line string, grid *[ySize][xSize]bool) {
	if strings.HasPrefix(line, pRect) {
		line = line[len(pRect):]
		AB := strings.Split(line, "x")
		a, _ := strconv.Atoi(AB[0])
		b, _ := strconv.Atoi(AB[1])
		rect(b, a, grid)
	} else if strings.HasPrefix(line, pRotRow) {
		line = line[len(pRotRow):]
		AB := strings.Split(line, " by ")
		a, _ := strconv.Atoi(AB[0])
		b, _ := strconv.Atoi(AB[1])
		rotRow(a, b, grid)
	} else if strings.HasPrefix(line, pRotCol) {
		line = line[len(pRotCol):]
		AB := strings.Split(line, " by ")
		a, _ := strconv.Atoi(AB[0])
		b, _ := strconv.Atoi(AB[1])
		rotCol(a, b, grid)
	}
}

func rect(Y, X int, grid *[ySize][xSize]bool) {
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			grid[y][x] = true
		}
	}
}

func rotCol(col, count int, grid *[ySize][xSize]bool) {
	for c := 0; c < count; c++ {
		overFlow := grid[ySize-1][col]
		for y := ySize - 2; y >= 0; y-- {
			grid[y+1][col] = grid[y][col]
		}
		grid[0][col] = overFlow
	}
}

func rotRow(row, count int, grid *[ySize][xSize]bool) {
	for c := 0; c < count; c++ {
		overFlow := grid[row][xSize-1]
		for x := xSize - 2; x >= 0; x-- {
			grid[row][x+1] = grid[row][x]
		}
		grid[row][0] = overFlow
	}
}

func printAndCountGrid(grid *[ySize][xSize]bool) (count int) {
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if grid[y][x] {
				count++
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	return count
}
