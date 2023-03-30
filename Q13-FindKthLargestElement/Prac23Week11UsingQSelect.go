package Q13_FindKthLargestElement

type Prac23Week11UsingQSelect struct {
}

func (p *Prac23Week11UsingQSelect) findKthLargestElement(arr []int, k int) int {
	return p.findKthLargestElementUsingQSelect(arr, 0, len(arr)-1, len(arr)-k)
}

func (p *Prac23Week11UsingQSelect) findKthLargestElementUsingQSelect(arr []int, start, end, k int) int {
	i, j := start, end
	pvt := arr[j]
	curr := i
	for curr < end {
		if arr[curr] >= pvt {
			curr++
		} else {
			arr[i], arr[curr] = arr[curr], arr[i]
			i++
			curr++
		}
	}

	arr[i], arr[j] = arr[j], arr[i]
	if i == k {
		return arr[k]
	} else if i < k {
		return p.findKthLargestElementUsingQSelect(arr, i+1, j, k)
	} else {
		return p.findKthLargestElementUsingQSelect(arr, start, i-1, k)
	}
}
