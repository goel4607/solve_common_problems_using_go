package Q13_FindKthLargestElement

import "sort"

type S1BF struct {
}

func (s1 S1BF) findKthLargestElement(arr []int, pos int) int {
	myArr := make([]int, 0, len(arr))
	myArr = append(myArr, arr...)
	sort.Ints(myArr)
	return myArr[len(myArr)-pos]
}
