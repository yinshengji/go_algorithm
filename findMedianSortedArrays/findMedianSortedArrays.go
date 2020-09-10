package findMedianSortedArrays

import "github.com/kr/pretty"

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lengthOfNums1 := len(nums1)
	lengthOfNums2 := len(nums2)
	minMedianIndex := 0
	minMedianValue := 0
	maxMedianIndex := 0
	maxMedianValue := 0
	totalIndex := 1
	if (lengthOfNums1 + lengthOfNums2) % 2 == 0 {
		minMedianIndex = (lengthOfNums1 + lengthOfNums2) / 2
		maxMedianIndex = ((lengthOfNums1 + lengthOfNums2) / 2) + 1
	} else {
		minMedianIndex = ((lengthOfNums1 + lengthOfNums2) / 2) + 1
		maxMedianIndex = ((lengthOfNums1 + lengthOfNums2) / 2) + 1
	}
	if lengthOfNums1 == 0 {
		return float64(nums2[minMedianIndex - 1] + nums2[maxMedianIndex - 1]) / 2.0
	} else if lengthOfNums2 == 0 {
		return float64(nums1[minMedianIndex - 1] + nums1[maxMedianIndex - 1]) / 2.0
	}
	var nums1Head int = 0
	var nums2Head int = 0
	var currentValue int = 0
	pretty.Println(minMedianIndex, maxMedianIndex)
	for totalIndex <= maxMedianIndex {
		if nums1Head > (lengthOfNums1 - 1) {
			currentValue = nums2[nums2Head]
			nums2Head++
		} else if nums2Head > (lengthOfNums2 - 1) {
			currentValue = nums1[nums1Head]
			nums1Head++
		} else {
			if nums1[nums1Head] <= nums2[nums2Head] {
				currentValue = nums1[nums1Head]
				nums1Head++
				if nums1Head > lengthOfNums1 {
					nums2Head++
				}
			} else {
				currentValue = nums2[nums2Head]
				nums2Head++
				if nums2Head > lengthOfNums2 {
					nums1Head++
				}
			}
		}
		pretty.Println(totalIndex, currentValue)
		if totalIndex == minMedianIndex {
			minMedianValue = currentValue
		} else if totalIndex == maxMedianIndex {
			maxMedianValue = currentValue
		}
		totalIndex++
	}

	if minMedianIndex == maxMedianIndex {
		return float64(minMedianValue) / 1.0
	} else {
		return float64(minMedianValue + maxMedianValue) / 2.0
	}
}
