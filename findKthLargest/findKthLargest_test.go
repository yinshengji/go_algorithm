package findKthLargest

import (
	"testing"
)

func TestFindKthLargest(t *testing.T)  {
	//普通
	nums := []int{2, 1, 3, 6, 2, 1, 4}
	k := 2
	n := findKthLargest(nums, 5)
	if n != k {
		t.Fail()
	}
	//全等
	nums = []int{3, 3, 3, 3, 3, 3, 3}
	k = 3
	n = findKthLargest(nums, 3)
	if n != k {
		t.Fail()
	}
	//k值错误
	nums = []int{3, 3, 3}
	k = -1
	n = findKthLargest(nums, 4)
	if n != k {
		t.Fail()
	}
}