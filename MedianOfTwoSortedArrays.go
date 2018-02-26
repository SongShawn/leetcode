package leetcode

import (
	"math"
)

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
		var j int
		if totalLen&1 == 0 {
			j = (totalLen)/2 - i
		} else {
			j = (totalLen-1)/2 - i
		}

		if j < longLen && i < shortLen {
			if shortNums[i-1] <= longNums[j] && longNums[j-1] <= shortNums[i] {
				if totalLen&1 == 0 { // 偶数长度
					return (math.Max(float64(shortNums[i-1]), float64(longNums[j-1])) +
						math.Min(float64(shortNums[i]), float64(longNums[j]))) / 2
				} else { // 基数长度
					return math.Max(float64(shortNums[i-1]), float64(longNums[j-1]))
				}
			}
		} else if i < shortLen {
			if longNums[j-1] <= shortNums[i] {
				if totalLen&1 == 0 { // 偶数长度
					return (math.Max(float64(shortNums[i-1]), float64(longNums[j-1])) +
						float64(shortNums[i])) / 2
				} else { // 基数长度
					return math.Max(float64(shortNums[i-1]), float64(longNums[j-1]))
				}
			}
		} else if j < longLen {
			if totalLen&1 == 0 { // 偶数长度
				return (math.Max(float64(shortNums[i-1]), float64(longNums[j-1])) +
					float64(longNums[j])) / 2
			} else { // 基数长度
				return math.Max(float64(shortNums[i-1]), float64(longNums[j-1]))
			}
		}

	}

	return 0.0
}
