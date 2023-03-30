package Q2_FindContainerWithMostWater

type MostWaterContainerFinderS3Prac struct {
}

func (s3 MostWaterContainerFinderS3Prac) findMostWaterContainer(nums []int) int {
	var maxW int

	lt, rt := 0, len(nums)-1
	for lt < rt {
		diff := rt - lt

		var min int
		if nums[lt] < nums[rt] {
			min = nums[lt]
			lt++
		} else {
			min = nums[rt]
			rt--
		}

		area := diff * min
		if maxW < area {
			maxW = area
		}
	}

	return maxW
}
