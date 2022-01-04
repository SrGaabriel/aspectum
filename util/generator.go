package util

import (
	"fmt"
	"github.com/SrGaabriel/aspectum/types"
	"math/rand"
	"time"
)

func GenerateRandomly(height int, width int) types.Maze {
	maze := types.CreateEmptyMaze(height, width)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			maze.Points[types.NewPoint(x, y)] = rand.Intn(2) == 0
		}
	}
	return maze
}

func GenerateProcedurally(height int, width int) types.Maze {
	maze := types.CreateFullMaze(height, width, true)
	x := rand.Intn(maze.Width)
	y := rand.Intn(maze.Height)
	startingPoint := types.NewPoint(x, y)
	maze.Points[startingPoint] = false

	frontiers := getCellsAroundPoint(maze, startingPoint, true)
	for len(frontiers) > 1 {
		fmt.Println("Frontiers Length:", len(frontiers), frontiers)
		randomFrontier := frontiers[rand.Intn(len(frontiers)):][0]
		randomFrontierNeighbors := getCellsAroundPoint(maze, randomFrontier, false)

		if len(randomFrontierNeighbors) > 0 {
			randomFrontierRandomNeighbor := randomFrontierNeighbors[rand.Intn(len(randomFrontierNeighbors))]
			connectDistantCells(maze, randomFrontier, randomFrontierRandomNeighbor)
		}
		frontiers = append(frontiers, getCellsAroundPoint(maze, randomFrontier, true)...)
		frontiers = remove(frontiers, randomFrontier)
		time.Sleep(1000000000)
	}
	return maze
}

func connectDistantCells(maze types.Maze, cell types.Point, neighbor types.Point) {
	maze.Points[types.NewPoint(cell.X+neighbor.X, cell.Y+neighbor.Y)] = false
}

func remove(cells []types.Point, cell types.Point) []types.Point {
	var newItems []types.Point
	for _, i := range cells {
		if i != cell {
			newItems = append(newItems, cell)
		}
	}
	return newItems
}

func getCellsAroundPoint(maze types.Maze, point types.Point, occupied bool) []types.Point {
	frontiers := []types.Point{
		types.NewPoint(point.X+2, point.Y),
		types.NewPoint(point.X-2, point.Y),
		types.NewPoint(point.X, point.Y+2),
		types.NewPoint(point.X, point.Y-2),
	}
	var validFrontiers []types.Point
	for _, frontierPoint := range frontiers {
		if maze.FitsMaze(frontierPoint) && occupied == maze.Points[frontierPoint] {
			validFrontiers = append(validFrontiers, frontierPoint)
		}
	}
	return validFrontiers
}
