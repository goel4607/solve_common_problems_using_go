package Q2_FindContainerWithMostWater

type Prac2 struct {
}

func (p2 Prac2) findMostWaterContainer(nums []int) int {
	var max int
	pL, pR := 0, len(nums)-1

	for pL < pR {
		currWidth := pR - pL

		var currHeight int
		if nums[pL] < nums[pR] {
			currHeight = nums[pL]
			pL++
		} else {
			currHeight = nums[pR]
			pR--
		}

		water := currWidth * currHeight
		if max < water {
			max = water
		}
	}

	return max
}
