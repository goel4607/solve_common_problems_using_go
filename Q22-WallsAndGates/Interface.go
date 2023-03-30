package Q22_WallsAndGates

import "math"

const (
	WALL = iota
	GATE
	INF = math.MaxInt32 //represents infinity; the default value in each room
)

type Interface interface {
	FillRoomsWithNumHopsToNearestGate(arr [][]int)
}
