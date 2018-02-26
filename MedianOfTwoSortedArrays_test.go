package leetcode

import "testing"

func TestMedianOfTSA00(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	t.Logf("ret %f\n", findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
}
