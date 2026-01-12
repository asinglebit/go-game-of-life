package main

import "fmt"

func main() {
    rows := 5
    cols := 5

    // Create a 2D grid
    grid := make([][]int, rows)
    for i := 0; i < rows; i++ {
        grid[i] = make([]int, cols)
        for j := 0; j < cols; j++ {
            grid[i][j] = i*cols + j + 1
        }
    }

    // Print the grid
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Printf("%3d ", grid[i][j])
        }
        fmt.Println()
    }
}