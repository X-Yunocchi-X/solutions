package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	findNthSmallestValue := func(k int) int {
		var result int
		n, m := len(nums1), len(nums2)
		l1, l2 := 0, 0
		for {
			index := min(k/2-1, n-l1-1, m-l2-1)
			if l1 >= n {
				result = nums2[l2+k-1]
				break
			}
			if l2 >= m {
				result = nums1[l1+k-1]
				break
			}
			if k == 1 {
				result = min(nums1[l1], nums2[l2])
				break
			}
			if nums1[l1+index] <= nums2[l2+index] {
				l1 = l1 + index + 1
			} else {
				l2 = l2 + index + 1
			}
			k -= index + 1
		}
		return result
	}
	result := 0.0
	length := (len(nums1) + len(nums2))
	if length&1 == 0 {
		result = float64(findNthSmallestValue(length/2)+findNthSmallestValue(length/2+1)) / 2
	} else {
		result = float64(findNthSmallestValue(length/2 + 1))
	}
	return result
}
