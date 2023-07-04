package Q3_FindTrappedWater

type TrappedWaterFinderPrac230601 struct {
}

// start time: 12:22 am

func (s TrappedWaterFinderPrac230601) findTrappedWater(w []int) int {
	if len(w) == 0 {
		return 0
	}
	//declare and initialize maxLeft and maxRight
	left, right := 0, len(w)-1
	maxLeft, maxRight := w[left], w[right]
	var ttlWtr int
	//iterate over the left < right
	for left <= right {
		//calc min out of maxLeft and maxRight
		if maxLeft <= maxRight {
			tw := maxLeft - w[left]
			if tw > 0 {
				ttlWtr += tw
			}

			if maxLeft < w[left] {
				maxLeft = w[left]
			}

			left++
		} else {
			tw := maxRight - w[right]
			if tw > 0 {
				ttlWtr += tw
			}

			if maxRight < w[right] {
				maxRight = w[right]
			}

			right--
		}
	}
	//end loop

	//return tw
	return ttlWtr
}
