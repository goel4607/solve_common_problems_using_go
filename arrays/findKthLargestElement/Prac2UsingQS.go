package findKthLargestElement

type Prac2UsingQS struct {
}

func (p *Prac2UsingQS) findKthLargestElement(arr []int, k int) int {
	// here we want to use quick sort (select) stopping when k is the pivot

	finalPos := len(arr) - k
	return p.quickSelect(0, len(arr)-1, finalPos, arr)
}

// 5, 4, 2, 7, 8, 3
// 5, 4, 2, 7, 8, 3 find the element which is < 3 and then replace it the first > element
// 2, 4, 5, 7, 8, 3 find the element which is < 3 and then replace it the first > element
// i = 2, j=3

// 2, 4, 5, 7, 8, 3 find the element which is < 3 and then replace it the first > element
// i = 2, in the end swap i and pivot
// 2, 3, 5, 7, 8, 4 find the element which is < 3 and then replace it the first > element
// 2, 3, 5, 7, 8, 4 find the element which is < 3 and then replace it the first > element

func (p *Prac2UsingQS) quickSelect(start, end, finalPos int, arr []int) int {
	// choose the last element as pivot and arrange it in the array
	pivot := end
	i, j := start, start

	for j < end {
		if arr[j] > arr[pivot] {
			j++
		} else {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j++
		}
	}

	arr[i], arr[pivot] = arr[pivot], arr[i]
	// consider only the portion where finalPos is present and ignore the other portion
	if finalPos == pivot {
		return arr[pivot]
	} else if finalPos < pivot {
		return p.quickSelect(start, pivot-1, finalPos, arr)
	} else {
		return p.quickSelect(pivot+1, end, finalPos, arr)
	}
	// exit when the current position is the same as final position
}
