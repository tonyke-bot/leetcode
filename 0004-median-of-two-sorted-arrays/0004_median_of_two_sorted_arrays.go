package problem0004

// Source: https://leetcode.com/problems/median-of-two-sorted-arrays
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalSize := len(nums1) + len(nums2)
	odd := totalSize % 2
	skip := int((totalSize+odd)/2) - 1

	index1 := 0
	index2 := 0

	getNext := func() (rv int) {
		switch {
		case len(nums1) == index1:
			rv = nums2[index2]
			index2++
		case len(nums2) == index2:
			rv = nums1[index1]
			index1++
		case nums1[index1] > nums2[index2]:
			rv = nums2[index2]
			index2++
		case nums1[index1] <= nums2[index2]:
			rv = nums1[index1]
			index1++
		}
		return
	}

	for ; skip > 0; skip-- {
		getNext()
	}

	if odd == 1 {
		return float64(getNext())
	}
	return float64(getNext()+getNext()) / 2
}

/*
Test Case:
[1,3]
[2]
*/
