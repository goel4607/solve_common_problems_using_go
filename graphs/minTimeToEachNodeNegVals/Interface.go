package minTimeToEachNode

type Node struct {
	start  int
	end    int
	weight int
}

type Interface interface {
	findMinTime(n int, edges []Node, start int) int
}
