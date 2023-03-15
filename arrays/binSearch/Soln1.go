package binSearch

type Soln1 struct {
}

func (s Soln1) searchIndexInSortedArray(arr []int, val int) int {
	i, j := 0, len(arr)-1
	for i < j {
		mid := (i + j) / 2
		if val == arr[mid] {
			return mid
		}

		if val < arr[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}

	return -1
}
