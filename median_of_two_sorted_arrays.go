package zj_lc

//There are two sorted arrays nums1 and nums2 of size m and n respectively.
//
//Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
//
//You may assume nums1 and nums2 cannot be both empty.
//
//Example 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//The median is 2.0
//Example 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//The median is (2 + 3)/2 = 2.5

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	//如果只有一个，找total/2，否则找total/2-1，total/2
	//每次循环如果sign==total/2或者sign==total/2-1，则说明已经找到中位数

	if total == 0 {
		return 0
	}

	if total == 1 {
		if len(nums1) > 0 {
			return float64(nums1[0])
		}
		return float64(nums2[0])
	}

	var i, j, sign, amount int

	f := func(n int) {
		if sign == total/2 {
			amount += n
		}

		if total%2 == 0 {
			if sign == total/2-1 {
				amount += n
			}
		}
	}

	for {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				f(nums1[i])
				i++
			} else {
				f(nums2[j])
				j++
			}
		} else if i < len(nums1) {
			f(nums1[i])
			i++
		} else {
			f(nums2[j])
			j++
		}

		if sign == total/2 {
			break
		}

		sign++
	}

	if total%2 != 0 {
		return float64(amount)
	}

	return float64(amount) / 2
}
