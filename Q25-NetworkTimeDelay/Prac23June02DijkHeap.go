package Q25_NetworkTimeDelay

import (
	"container/heap"
	"math"
)

// Start time: 	11:38 pm
// End time:	12:25 pm

type Prac23June02DijkHeap struct {
}

func (p Prac23June02DijkHeap) NetworkDelayTime(times [][]int, n, k int) int {
	//create adj list such that it lists all adj nodes to the specific nodes
	adjList := p.createAdjList(times, n)

	//create calc times arr and init it with math.MaxInt value
	delays := make([]int, n, n)
	for i := range delays {
		delays[i] = math.MaxInt
	}
	delays[k-1] = 0

	//create a min priority heap say minHeap and initialize it with the starting K node
	mh := &MinHeap{pos: make([]int, 0), ref: &delays}
	heap.Init(mh)

	heap.Push(mh, k-1)

	for mh.Len() > 0 {

		//iterate over the heap and pop the minimum element
		start := heap.Pop(mh).(int)

		//iterate over all the adjacent nodes to the min element
		for _, adj := range adjList[start] {
			end := adj[0]
			delay := adj[1]

			//calc the new distance by adding the orig distance to the adj node distance
			newDelay := delays[start] + delay
			if newDelay < delays[end] {

				//update the distance if the new dist is < than the one present earlier
				delays[end] = newDelay

				//add the updated nodes to the heap
				heap.Push(mh, end)
			}
		}
	}

	//iterate over the calc node times array and check if any node is still un-reachable i.e. it still has Max time
	var max int
	for _, d := range delays {
		if d == math.MaxInt {
			//if yes, then return -1
			return -1
		}

		//otherwise calc max time
		if max < d {
			max = d
		}
	}

	//return the max time
	return max
}

func (p Prac23June02DijkHeap) createAdjList(times [][]int, n int) [][][]int {
	adjList := make([][][]int, n, n)

	for _, t := range times {
		start := t[0] - 1
		end := t[1] - 1
		delay := t[2]

		adjs := adjList[start]
		if len(adjs) == 0 {
			adjs = make([][]int, 0)
		}

		adjList[start] = append(adjs, []int{end, delay})
	}

	return adjList
}

type MinHeap struct {
	pos []int
	ref *[]int
}

func (m *MinHeap) Len() int {
	return len(m.pos)
}

func (m *MinHeap) Swap(i, j int) {
	m.pos[i], m.pos[j] = m.pos[j], m.pos[i]
}

func (m *MinHeap) Less(i, j int) bool {
	if (*m.ref)[i] < (*m.ref)[j] {
		return true
	}

	return false
}

func (m *MinHeap) Push(v any) {
	m.pos = append(m.pos, v.(int))
}

func (m *MinHeap) Pop() any {
	oldLen := len(m.pos)
	v := m.pos[oldLen-1]
	m.pos = m.pos[:oldLen-1]

	return v
}
