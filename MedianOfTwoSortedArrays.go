package leetcode

// 思路：
// 以较长数组为基准，取中间值longMed
// 以二分法在较短数组中找到最后一个小于longMed的值的索引i
// 短数组[0,i]会放到长数组的左边，(i,n)放到长数组的右边
// 根据左右两边添加数据个数的差值修正longMed的索引即可
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0.0
}