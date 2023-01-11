package findTrappedWater

type TrappedWaterFinderSolution1BF struct {
}

func (t TrappedWaterFinderSolution1BF) findTrappedWater(height []int) int {
	var water, maxLeft int

	//maxLeft, maxRight := make([]int, 0, len(height)), make([]int, 0, len(height))
	// iterate over the height
	for i := range height {
		n := height[i]
		// select the first number say n
		// find the max left and max right say mxLeft and mxRight
		var maxRight int
		for j := i + 1; j < len(height); j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}
		if i > 0 && height[i-1] > maxLeft {
			maxLeft = height[i-1]
		}
		// find the min of them: min (mxLeft, mxRight) say min
		min := maxLeft
		if maxLeft > maxRight {
			min = maxRight
		}
		// if min > 0 then calc the diff between min and this number n say the result as d

		if min > 0 && min-n > 0 {
			d := min - n
			// if d > 0 then it tells how much height is trapped there and add this water
			// select the next number
			water += d

			//fmt.Printf("Trapped height at [%d]=%d\n", i, d)
		}
	}

	// return water
	return water
}
