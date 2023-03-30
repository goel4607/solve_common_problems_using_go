package Q2_FindContainerWithMostWater

type Prac23Week11 struct {
}

// start: 10:25 pm
// end: 10:42 pm
func (p Prac23Week11) findMostWaterContainer(nums []int) int {
	//take two pointers pLeft and pRight
	pLeft, pRight := 0, len(nums)-1
	//declare max area with init as zero
	var max int
	//iterate over the nums while pLeft < pRight
	for pLeft < pRight {
		//calc the h by taking the min and multiply with the diff between indexes i.e. j - i; update max if current max is >
		w := pRight - pLeft

		var h int
		//check which element is less: pLeft or pRight and move its value
		//i.e. if nums[pLeft] < nums[pRight] then pLeft++ otherwise pRight--
		if nums[pLeft] < nums[pRight] {
			h = nums[pLeft]
			pLeft++
		} else {
			h = nums[pRight]
			pRight--
		}

		area := w * h
		if max < area {
			max = area
		}
	}

	return max
}
