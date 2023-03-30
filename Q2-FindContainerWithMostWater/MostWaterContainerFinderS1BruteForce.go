package Q2_FindContainerWithMostWater

import "math"

type MostWaterContainerFinderS1BF struct {
}

func (s MostWaterContainerFinderS1BF) findMostWaterContainer(nums []int) int {
	var maxArea int

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			minHeight := int(math.Min(float64(nums[i]), float64(nums[j])))
			width := j - i

			area := minHeight * width

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
