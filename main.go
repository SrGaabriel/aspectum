package main

import (
	"fmt"
	"github.com/SrGaabriel/aspectum/types"
	"github.com/SrGaabriel/aspectum/util"
)

func main() {
	maze := util.GenerateProcedurally(16, 16)
	for x := 0; x < maze.Width; x++ {
		for y := 0; y < maze.Height; y++ {
			point := maze.Points[types.NewPoint(x, y)]
			if point {
				fmt.Print("⬛")
			} else {
				fmt.Print("⬜")
			}
		}
		fmt.Println("")
	}
}
