package findKthLargest
/*
 * 寻找第k大的元素
 * leetcode: 215
 * 利用快排partition函数，但不进行完全排序
 */
func findKthLargest(nums []int, k int) int {
	start := 0
	length := len(nums)
	//不存在的K, 直接返回-1
	if k > length || k < 0 {
		return -1
	}
	end := length - 1
	k = length - k
	for {
		index := partition(&nums, start, end)
		if index < k {
			start = index + 1
		} else if index > k {
			end = index - 1
		} else {
			return nums[k]
		}
	}
}

func partition(originalNums *[]int, start int, end int) int {
	if start == end {
		return start
	}
	nums := *originalNums
	//确认分割值
	pivot := nums[start]
	i := start + 1
	j := end
	for {
		for nums[j] >= pivot {
			if j == start {
				break
			}
			j -= 1
		}

		for nums[i] < pivot {
			if i == end {
				break
			}
			i += 1
		}

		if i >= j {
			break
		}
		nums = swap(nums, i, j)
	}
	nums = swap(nums, j, start)
	originalNums = &nums
	return j
}

func swap(nums []int, i int, j int) []int {
	nums[i], nums[j] = nums[j], nums[i]
	return nums
}