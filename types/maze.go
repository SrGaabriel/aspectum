package types

type Maze struct {
	Height int
	Width  int
	Points map[Point]bool
}
type Point struct {
	X int
	Y int
}

func CreateEmptyMaze(height int, width int) Maze {
	return Maze{Height: height, Width: width, Points: make(map[Point]bool)}
}

func CreateFullMaze(height int, width int, state bool) Maze {
	maze := CreateEmptyMaze(height, width)
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			maze.Points[NewPoint(x, y)] = state
		}
	}
	return maze
}

func NewPoint(x int, y int) Point {
	return Point{X: x, Y: y}
}

func (maze Maze) FitsMaze(point Point) bool {
	return point.X >= 0 && point.X < maze.Width && point.Y >= 0 && point.Y < maze.Height
}
