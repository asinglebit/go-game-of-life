package main

import "fmt"

const WIDTH = 100
const HEIGHT = 50

func main() {

	size := WIDTH * HEIGHT
	grid:= make([]int, size)

	for i := 0; i < size; i++ {
		grid[i] = i + 1
	}

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
