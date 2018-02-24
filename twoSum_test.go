package leetcode

import (
	"testing"
)

func StringSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestTwoNum00(t *testing.T) {

	except := []int{0, 1}
	result := twoSum([]int{2, 7, 11, 15}, 9)

	if !StringSliceEqual(except, result) {
		t.Errorf("failed. except=%v, result=%v",
			except, result)
	}
	return
}

//[-3,4,3,90]
func TestTwoNum01(t *testing.T) {

	except := []int{0, 2}
	result := twoSum([]int{-3, 4, 3, 90}, 0)

	if !StringSliceEqual(except, result) {
		t.Errorf("failed. except=%v, result=%v",
			except, result)
	}
	return
}
