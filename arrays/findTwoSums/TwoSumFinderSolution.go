package findTwoSums

type TwoSumFinderSolution1BF struct {
}

func (ts TwoSumFinderSolution1BF) findTwoSum(nums []int, target int) []int {
	//var result []int
	result := make([]int, 0)
	//make(map[int]int)

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return result
}
