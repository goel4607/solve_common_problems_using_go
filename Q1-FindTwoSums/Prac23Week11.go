package Q1_FindTwoSums

//start time: 8:40 pm
//end time: 8:52 pm

type Prac23Week11 struct {
}

func (p Prac23Week11) findTwoSum(nums []int, tgt int) []int {
	//check if len is <= 1 then return an empty array
	if len(nums) <= 1 {
		return []int{}
	}

	//create a map with key as num and value as index
	idxByNum := make(map[int]int)
	//iterate over the nums array
	for i, v := range nums {
		//check if tgt - nums[i] is present in map
		otherNum := tgt - v
		if idx, ok := idxByNum[otherNum]; ok {
			//return a new array with value and current values index
			return []int{idx, i}
		} else {
			//if not then add the current num in map as key and value as its index
			idxByNum[v] = i
		}
		//end of loop
	}

	//return empty array
	return []int{}
}
