package findGatePath

import "math"

const (
	WALL = iota
	GATE
	INF = math.MaxInt32
)

type Interface interface {
	fillRoomsWithNumHopsToNearsestGate(arr [][]int)
}
