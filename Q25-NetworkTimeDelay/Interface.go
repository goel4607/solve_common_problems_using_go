package Q25_NetworkTimeDelay

type Node struct {
	Start  int
	End    int
	Weight int
}

type NodeWeight struct {
	index  int
	weight int
}

type Interface interface {
	FindMinTime(n, k int, times []Node) int
}
