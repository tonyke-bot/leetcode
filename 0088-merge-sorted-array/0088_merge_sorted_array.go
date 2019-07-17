package problem0088

// Source: https://leetcode.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2[:n])
		return
	}
	if n == 0 {
		return
	}

	if nums1[m-1] < nums2[0] {
		copy(nums1[m:], nums2[:n])
		return
	}

	if nums2[n-1] < nums1[0] {
		copy(nums1[n:], nums1[:m])
		copy(nums1[:n], nums2[:n])
		return
	}

	pos1 := 0
	for index, num2 := range nums2 {
		for pos1 < m && nums1[pos1] < num2 {
			pos1++
		}

		if pos1 >= m {
			copy(nums1[pos1:], nums2[index:])
			return
		}

		copy(nums1[pos1+1:], nums1[pos1:])
		nums1[pos1] = num2
		m++
	}
}

/*
Test Case:
[1,2,3,0,0,0]
3
[2,5,6]
3
*/
