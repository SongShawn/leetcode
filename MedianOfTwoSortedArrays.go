package leetcode

import (
	"math"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return findMedianSortedArrays(nums1, nums2)
}

// 思路：
// A B两个数组
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	len1 := len(nums1)
	len2 := len(nums2)
	shortLen := len1
	longLen := len2
	shortNums := nums1
	longNums := nums2

	if len1 == 0 && len2 == 0 {
		return 0.0
	}

	if len1 > len2 {
		shortLen = len2
		longLen = len1
		shortNums = nums2
		longNums = nums1
	}

	totalLen := len1 + len2
	for i := 1; i < shortLen+1; i++ {
		j := (totalLen+1)/2 - i

		// i < shortLen ---> j > 0
		if i < shortLen {
			if shortNums[i-1] <= longNums[j] && longNums[j-1] <= shortNums[i] {
				if totalLen&1 == 0 { // 偶数长度
					return (math.Max(float64(shortNums[i-1]), float64(longNums[j-1])) +
						math.Min(float64(shortNums[i]), float64(longNums[j]))) / 2
				} else { // 奇数长度
					return math.Max(float64(shortNums[i-1]), float64(longNums[j-1]))
				}
			}
		} else if i == shortLen {
			if j == 0 { // 个数相等且一组数据都大于另一组数据
				if shortNums[0] >= longNums[longLen-1] {
					return (float64(longNums[longLen-1]) + float64(shortNums[0])) / 2
				} else if shortNums[longLen-1] <= longNums[0] {
					return (float64(shortNums[longLen-1]) + float64(longNums[0])) / 2
				}
			} else if shortNums[i-1] <= longNums[j] {
				if totalLen&1 == 0 { // 偶数长度
					return (math.Max(float64(shortNums[i-1]), float64(longNums[j-1])) +
						float64(longNums[j])) / 2
				} else { // 奇数长度
					return math.Max(float64(shortNums[i-1]), float64(longNums[j-1]))
				}
			} else {
				if totalLen&1 == 0 {
					return (float64(longNums[(totalLen+1)/2-1]) + math.Min(float64(shortNums[0]),
						float64(longNums[(totalLen+1)/2])))
				} else {
					return float64(longNums[(totalLen+1)/2-1])
				}
			}
		}
	}

	return 0.0
}
