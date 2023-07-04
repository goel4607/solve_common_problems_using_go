package Q1_FindTwoSums

type TwoSumsFinder230601 struct {
}

func (s TwoSumsFinder230601) findTwoSum(nums []int, target int) []int {
	//create a map with key as num and value as index
	idxByNum := make(map[int]int, 0)
	//iterate over the nums
	for i := 0; i < len(nums); i++ {
		//find the other number by subtracting the target from the current number say on = target - nums[i]
		on := target - nums[i]
		//check if the other number is present in map
		if v, ok := idxByNum[on]; ok {
			//if yes, then return the other number's mapped index and the current index i
			return []int{v, i}
		}
		//if no, then add the current number as key and its index i in the map
		idxByNum[nums[i]] = i
		// loop end
	}

	// return empty array to indicate that the other number is not found
	return []int{}
}
