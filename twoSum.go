package leetcode

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		a := []int{i}
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				a = append(a, j)
				return a
			}
		}
	}
	return []int{}
}
