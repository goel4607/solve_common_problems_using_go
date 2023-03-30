package Q1_FindTwoSums

type TwoSumFinderS2Efficient struct {
}

func (s TwoSumFinderS2Efficient) findTwoSum(nums []int, target int) []int {
	r := make([]int, 0)

	// mapping of other number with other number's index
	otherNumVsIdxMap := make(map[int]int, 0)

	// iterate over the nums
	// for each num
	for i, n := range nums {
		//  check if map contains this number then return that number's index and this number's index
		if v, ok := otherNumVsIdxMap[n]; ok {
			r = append(r, v, i)
			break
		} else { // 	if not then store the diff between target and this number in the map with its value as this number index
			otherNumVsIdxMap[target-n] = i
		}
	}

	return r
}
