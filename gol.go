package main

import (
	"fmt"
	"time"
	"math/rand/v2"
)

const WIDTH int = 50
const HEIGHT int = 25
var grid [HEIGHT][WIDTH]int

type Seed int64
const (
	Random Seed = iota
	Glider
	Blinker
)

type Axis int64
const (
	X Axis = iota
	Y
)

func print_grid() {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if grid[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func set_initial (s Seed) {
	switch s {
	case Random:
		for i := 0; i < HEIGHT; i++ {
			for j := 0; j < WIDTH; j++ {
				grid[i][j] = rand.IntN(2)
			}
		}
	case Glider:
		grid[0][1] = 1
		grid[1][2] = 1
		grid[2][0] = 1
		grid[2][1] = 1
		grid[2][2] = 1
	case Blinker:
		grid[0][1] = 1
		grid[1][1] = 1
		grid[2][1] = 1
	default:
		panic("Invalid initial condition")
	}
}

func fix_bounds(val int, axis Axis) int {
	if axis == X {
		if val < 0 {
			return WIDTH + val
		} else if val >= WIDTH {
			return val - WIDTH
		}
	} else {
		if val < 0 {
			return HEIGHT + val
		} else if val >= HEIGHT {
			return val - HEIGHT
		}
	}

	return val
}

func count_neighbors(x int, y int) int {
	count := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			fx := fix_bounds(x + j, X)
			fy := fix_bounds(y + i, Y)
			if grid[fy][fx] == 1 {
				count += grid[fy][fx]
			}
		}
	}
	return count
}

func update_grid() {
	var new_grid [HEIGHT][WIDTH]int = grid

	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			neighbors := count_neighbors(j, i)
			if grid[i][j] == 1 {
				if neighbors < 2 || neighbors > 3 {
					new_grid[i][j] = 0
				}
			} else {
				if neighbors == 3 {
					new_grid[i][j] = 1
				}
			}
		}
	}
	grid = new_grid
}

func run(total_time int, intervall int) {
	timer := 0

	for timer < total_time {
		clear()
		update_grid()
		print_grid()
		timer += intervall
		time.Sleep(time.Duration(intervall) * time.Millisecond)
	}

}


func clear() {
	fmt.Print("\033[H\033[2J")
}

func hide() {
	fmt.Print("\033[?25l")
}

func main() {
	hide()
	clear()
	set_initial(Random)
    print_grid()

    run(5000, 100)
}
