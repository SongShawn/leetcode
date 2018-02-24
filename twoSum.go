package leetcode

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

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

// 普通解法的时间主要浪费在查找上
// 而使用hashmap可以大大降低查找的复杂度
func twoSumWithHashMap(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	var numMap = make(map[int]int)
	for i, v := range nums {
		numMap[v] = i
	}

	for i := 0; i < len(nums); i++ {
		if j, ok := numMap[target-nums[i]]; ok {
			if i == j {
				continue
			}
			return []int{i, j}
		}
	}
	return []int{}
}
