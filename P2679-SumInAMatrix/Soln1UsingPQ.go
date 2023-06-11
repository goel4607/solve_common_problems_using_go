package P679_SumInAMatrix

import (
	"container/heap"
	"fmt"
)

type Soln1UsingPQ struct {
}

func (s Soln1UsingPQ) matrixSum(nums [][]int) int {
	pqs := make([]*MaxPQ, len(nums), len(nums))

	//iterate over the nums
	for i, row := range nums {
		fmt.Println(i, len(row))
		//make a new matrix with priority queue for each row say pqs
		pq := &MaxPQ{}
		for _, v := range row {
			heap.Push(pq, v)
		}

		pqs[i] = pq
	}

	//declare sum int
	var sum int

	for i := 0; i < len(nums[0]); i++ {
		//create a new pq for this row say cPQ
		c := &MaxPQ{}

		//iterate over the pqs
		for _, pq := range pqs {
			//if len(*pq) < 1 {
			//	break
			//}
			//fmt.Println(i,pq.Len())
			//pop the max number of each pq and add this number to cPQ
			m := heap.Pop(pq).(int)
			heap.Push(c, m)
			//pop the max number of cPQ add it into sum
		}

		//if len(*c) > 0 {
		sum += heap.Pop(c).(int)
		//}
	}

	return sum
}

type MaxPQ []int

func (pq *MaxPQ) Push(n any) {
	*pq = append(*pq, n.(int))
}

func (pq *MaxPQ) Pop() any {
	oldLen := len(*pq)
	v := (*pq)[oldLen-1]
	*pq = (*pq)[:oldLen-1]

	return v
}

func (pq *MaxPQ) Less(i, j int) bool {
	if (*pq)[i] < (*pq)[j] {
		return false // max pq
	}

	return true
}

func (pq *MaxPQ) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq MaxPQ) Len() int {
	return len(pq)
}
