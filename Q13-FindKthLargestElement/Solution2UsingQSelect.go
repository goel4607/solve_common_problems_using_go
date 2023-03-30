package Q13_FindKthLargestElement

type S2UsingQSelect struct {
}

func (s2 S2UsingQSelect) findKthLargestElement(arr []int, pos int) int {
	elmtPos := len(arr) - pos
	myArr := make([]int, 0, len(arr))
	myArr = append(myArr, arr...)

	reqdElmt := findLargestElement(myArr, 0, len(arr)-1, elmtPos)
	return reqdElmt
}

func findLargestElement(arr []int, start, end, elmtPos int) int {
	part := partition(arr, start, end)
	if elmtPos == part {
		return arr[elmtPos]
	} else if elmtPos > part {
		return findLargestElement(arr, part+1, end, elmtPos)
	} else {
		return findLargestElement(arr, start, part-1, elmtPos)
	}
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start
	j := start

	for j < end {
		if arr[j] > pivot {
			j++
		} else {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			j++
			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]
	return i
}
