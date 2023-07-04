package binSearch

type Soln1 struct {
}

func (s Soln1) startAndEndOfTgtInSortedArr(arr []int, val int) [2]int {
	pos := s.binarySearch(arr, val, 0, len(arr)-1)
	if pos == -1 {
		return [2]int{-1, -1}
	}

	left := pos
	for left > 0 {
		leftNew := s.binarySearch(arr, val, 0, left-1)
		if leftNew == -1 {
			break
		}

		left = leftNew
	}

	right := pos
	for right < len(arr) {
		rightNew := s.binarySearch(arr, val, right+1, len(arr)-1)
		if rightNew == -1 {
			break
		}

		right = rightNew
	}

	return [2]int{left, right}
}

func (s Soln1) binarySearch(arr []int, val int, left, right int) int {
	i, j := left, right
	for i <= j {
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
