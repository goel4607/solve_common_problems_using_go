package Q3_FindTrappedWater

type TrappedWaterFinderSolution2Efficient struct {
}

func (t TrappedWaterFinderSolution2Efficient) findTrappedWater(height []int) int {
	var total, maxL, maxR int

	pL, pR := 0, len(height)-1
	for pL < pR {
		var diff int
		if maxL <= maxR {
			min := maxL
			diff = min - height[pL]
			if maxL < height[pL] {
				maxL = height[pL]
			}

			if maxL <= maxR {
				pL++
			}
		} else {
			min := maxR
			diff = min - height[pR]

			if maxR < height[pR] {
				maxR = height[pR]
			}

			if maxR < maxL {
				pR--
			}
		}

		if diff > 0 {
			total += diff
		}
	}

	return total
}
