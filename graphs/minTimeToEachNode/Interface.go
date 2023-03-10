package minTimeToEachNode

type Node struct {
	start  int
	end    int
	weight int
}

type Interface interface {
	findMinTime(n, start int, adjList []Node) int
}
