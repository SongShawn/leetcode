package leetcode

import "testing"

func TestMedianOfTSA00(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	if 2.0 != findMedianSortedArrays(nums1, nums2) {
		t.Error(nums1, nums2)
	}

	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	if 2.5 != findMedianSortedArrays(nums1, nums2) {
		t.Error(nums1, nums2)
	}

	nums1 = []int{}
	nums2 = []int{}
	if 0 != findMedianSortedArrays(nums1, nums2) {
		t.Error(nums1, nums2)
	}

	nums1 = []int{12, 23}
	nums2 = []int{1, 2, 3, 4, 5}
	if 4 != findMedianSortedArrays(nums1, nums2) {
		t.Error(nums1, nums2)
	}

	nums1 = []int{1, 23}
	nums2 = []int{1, 2, 3, 4, 5}
	if 3 != findMedianSortedArrays(nums1, nums2) {
		t.Error(nums1, nums2)
	}

}
