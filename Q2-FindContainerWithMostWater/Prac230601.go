package Q2_FindContainerWithMostWater

type Prac230601 struct {
}

func (s Prac230601) findMostWaterContainer(nums []int) int {
	//dec maxLeft and maxRight
	var maxArea int
	//dec left=0 and right=len(nums)-1
	left, right := 0, len(nums)-1
	//iterate till left < right
	for left < right {
		//find diff in indexes
		d := right - left

		var area int
		//find min our of max left and max right
		if nums[left] < nums[right] {
			//calc curr area
			area = d * nums[left]
			left++
		} else {
			//calc curr area
			area = d * nums[right]
			//update left if left val < right val or vice versa
			right--
		}

		//update maxArea if current area is > maxArea
		if maxArea < area {
			maxArea = area
		}
		//loop end
	}

	return maxArea
}

func findMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
