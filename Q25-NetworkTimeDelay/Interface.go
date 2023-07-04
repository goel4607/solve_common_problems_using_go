package Q25_NetworkTimeDelay

//type Node struct {
//	Start  int
//	End    int
//	Weight int
//}

type NodeWeight struct {
	index  int
	weight int
}

type Interface interface {
	NetworkDelayTime(times [][]int, n, k int) int
}
