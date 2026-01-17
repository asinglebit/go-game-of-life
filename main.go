package main
import "fmt"
import "time"
import "math/rand"
const WIDTH = 94
const HEIGHT = 78
const SIZE = WIDTH * HEIGHT
func print_grid(grid []int) {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			if (grid[i * WIDTH + j] == 0) {
				fmt.Printf(". ") 
			} else {
				fmt.Printf("x ") 
			}
		}
		fmt.Print(" \n")	
    }
}
func random_initialize(grid []int) {
	for i := 0; i < SIZE; i++ {
		grid[i] = rand.Intn(2)
	}
}
func next_generation(grid_a []int, grid_b []int) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			neighbours := 0
			cells := [3]int{-1, 0, 1}
			for _, dy := range cells {
				for _, dx := range cells {
					if dx == 0 && dy == 0 { continue } 
					ny := (y + dy + HEIGHT) % HEIGHT
					nx := (x + dx + WIDTH) % WIDTH
					if (ny >= 0 && ny < HEIGHT && nx >= 0 && nx < WIDTH) {
						if grid_a[ny * WIDTH + nx] == 1 {
							neighbours += 1
						}
					}
				}
			}
			position := y * WIDTH + x
			switch neighbours {
				case 2:
					grid_b[position] = grid_a[position]
				case 3:
					grid_b[position] = 1
				default:
					grid_b[position] = 0
			}
		}
    }
}
func clearTerminal() {
    fmt.Print("\033[2J")
    fmt.Print("\033[H")
}
func main() {
	grid_a:= make([]int, SIZE)
	grid_b:= make([]int, SIZE)
	random_initialize(grid_a) 
	for {
		clearTerminal()
		print_grid(grid_a)
		next_generation(grid_a, grid_b)
		grid_a, grid_b = grid_b, grid_a
        time.Sleep(100 * time.Millisecond)
    }
}
