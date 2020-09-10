package findMedianSortedArrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	nums1 := []int{1,2,3}
	nums2 := []int{2,4,5}
	t.Log(FindMedianSortedArrays(nums1, nums2))
}
