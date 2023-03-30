package Q3_FindTrappedWater

type TrappedWaterFinderS3Att struct {
}

func (s3 TrappedWaterFinderS3Att) findTrappedWater(w []int) int {
	if len(w) <= 2 {
		return 0
	}

	var tw, maxL, maxR int
	left, right := 0, len(w)-1

	for left <= right {
		if maxL <= maxR {
			if maxL > 0 {
				d := maxL - w[left]
				if d > 0 {
					tw += d
				}
			}

			if maxL < w[left] {
				maxL = w[left]
			}

			left++
		} else {
			if maxR > 0 {
				d := maxR - w[right]
				if d > 0 {
					tw += d
				}
			}

			if maxR < w[right] {
				maxR = w[right]
			}

			right--
		}
	}

	return tw
}
