package Q3_FindTrappedWater

type FindTrappedWaterPrac2 struct {
}

func (p FindTrappedWaterPrac2) findTrappedWater(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var water int

	pL, pR := 0, len(arr)-1
	maxL, maxR := arr[pL], arr[pR]
	for pL <= pR {
		if maxL <= maxR {
			if maxL > 0 && maxL > arr[pL] {
				water += maxL - arr[pL]
			}

			if maxL < arr[pL] {
				maxL = arr[pL]
			}

			pL++
		} else {
			if maxR > 0 && maxR > arr[pR] {
				water += maxR - arr[pR]
			}

			if maxR < arr[pR] {
				maxR = arr[pR]
			}

			pR--
		}
	}

	return water
}
