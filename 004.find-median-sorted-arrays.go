package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	if len(nums1) > len(nums2) {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
	}
	halflen := (len(nums1) + len(nums2) + 1) / 2
	iMin, iMax := 0, len(nums1)

	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halflen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			var left int
			if i == 0 {
				left = nums2[j-1]
			} else if j == 0 {
				left = nums1[i-1]
			} else {
				if nums1[i-1] > nums2[j-1] {
					left = nums1[i-1]
				} else {
					left = nums2[j-1]
				}
			}
			if (len(nums1)+len(nums2))%2 == 1 {
				return float64(left)
			}
			var right int
			if i == len(nums1) {
				right = nums2[j]
			} else if j == len(nums2) {
				right = nums1[i]
			} else {
				if nums1[i] < nums2[j] {
					right = nums1[i]
				} else {
					right = nums2[j]
				}
			}
			return float64(left+right) / 2
		}
	}
	return 0.0
}
