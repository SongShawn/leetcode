package leetcode

import (
	//"expvar"

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

	var maxLeft, minRight float64
	totalLen := len1 + len2
	if shortLen == 0 {
		minRight = float64(longNums[totalLen/2])
		if totalLen&1 == 0 {
			maxLeft = float64(longNums[totalLen/2-1])
		} else {
			maxLeft = minRight
		}
	}

	for i := 1; i < shortLen+1; i++ {
		j := (totalLen+1)/2 - i

		// i < shortLen ---> j > 0
		if i < shortLen {
			if shortNums[i-1] <= longNums[j] && longNums[j-1] <= shortNums[i] {
				maxLeft = math.Max(float64(shortNums[i-1]), float64(longNums[j-1]))
				if totalLen&1 == 0 { // 偶数长度
					minRight = math.Min(float64(shortNums[i]), float64(longNums[j]))
				} else { // 奇数长度
					minRight = maxLeft
				}
				break
			}
		} else if i == shortLen {
			if shortNums[i-1] > longNums[j] {
				maxLeft = float64(longNums[(totalLen+1)/2-1])
				if totalLen&1 == 0 { // 偶数长度
					if (totalLen+1)/2 == longLen {
						minRight = float64(shortNums[0])
					} else {
						minRight = math.Min(float64(shortNums[0]), float64(longNums[(totalLen+1)/2]))
					}
				} else { // 奇数长度
					minRight = maxLeft
				}
			} else {
				if j == 0 {
					maxLeft = float64(shortNums[i-1])
					minRight = float64(longNums[0])
				} else {
					maxLeft = math.Max(float64(shortNums[i-1]), float64(longNums[j-1]))
					if totalLen&1 == 0 {
						minRight = float64(longNums[j])
					} else {
						minRight = maxLeft
					}
				}
			}
		}

	}

	return (maxLeft + minRight) / 2
}
