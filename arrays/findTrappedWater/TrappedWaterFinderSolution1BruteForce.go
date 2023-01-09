package findTrappedWater

type TrappedWaterFinderSolution1BF struct {
}

func (t TrappedWaterFinderSolution1BF) findTrappedWater(water []int) int {
	var trappedWater, maxLeft int

	//maxLeft, maxRight := make([]int, 0, len(water)), make([]int, 0, len(water))
	// iterate over the water
	for i := range water {
		n := water[i]
		// select the first number say n
		// find the max left and max right say mxLeft and mxRight
		var maxRight int
		for j := i + 1; j < len(water); j++ {
			if water[j] > maxRight {
				maxRight = water[j]
			}
		}
		if i > 0 && water[i-1] > maxLeft {
			maxLeft = water[i-1]
		}
		// find the min of them: min (mxLeft, mxRight) say min
		min := maxLeft
		if maxLeft > maxRight {
			min = maxRight
		}
		// if min > 0 then calc the diff between min and this number n say the result as d

		if min > 0 && min-n > 0 {
			d := min - n
			// if d > 0 then it tells how much water is trapped there and add this trappedWater
			// select the next number
			trappedWater += d

			//fmt.Printf("Trapped water at [%d]=%d\n", i, d)
		}
	}

	// return trappedWater
	return trappedWater
}
