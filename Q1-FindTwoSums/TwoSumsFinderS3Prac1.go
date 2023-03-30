package Q1_FindTwoSums

type TwoSumFinderS3 struct {
}

func (s3 TwoSumFinderS3) findTwoSum(nums []int, tgt int) []int {
	if len(nums) <= 1 {
		return []int{}
	}

	// make a map with int as key and index as value
	numVsIdx := make(map[int]int, 0)

	// iterate over the arr
	for i, v := range nums {
		// calc diff between nums[i] and target and see if that diff is in map
		d := tgt - v
		// if not then store that number with index as value
		if j, ok := numVsIdx[d]; ok {
			return []int{j, i}
		} else { // else we found our two sum; return map index and this index
			numVsIdx[v] = i
		}
	}

	return []int{}
}
