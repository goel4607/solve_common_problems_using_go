package Q2_FindContainerWithMostWater

type MostWaterContainerFinderS2Efficient struct {
}

func (s MostWaterContainerFinderS2Efficient) findMostWaterContainer(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	var max int

	i := 0
	j := len(nums) - 1

	for i < j {
		var minHeight, width int
		if nums[i] <= nums[j] {
			minHeight = nums[i]
			width = j - i

			i++
		} else {
			minHeight = nums[j]
			width = j - i

			j--
		}

		area := minHeight * width
		if area > max {
			max = area
		}
	}

	return max
}
