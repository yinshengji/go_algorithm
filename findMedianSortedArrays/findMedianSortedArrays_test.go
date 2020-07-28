package findMedianSortedArrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{2,3,4}
	t.Log(FindMedianSortedArrays(nums1, nums2))
}
