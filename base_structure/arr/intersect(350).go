package arr

import "sort"

/*
两个数组的交集
给定两个数组，编写一个函数来计算它们的交集。
*/

//常规解法：用map保存一个数组中元素的出现次数
//func intersect(arr1 []int, arr2 []int) []int {
//	m := map[int]int{}
//
//	var result []int
//
//	for _, v := range arr1 {
//		m[v] += 1
//	}
//
//	for _, v := range arr2 {
//		if m[v] > 0 {
//			result = append(result, v)
//			m[v]--
//		}
//	}
//
//	return result
//}

//如果两个数组是有序的，可以使用“双指针”
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	var result []int

	i, j := 0, 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}

	return result
}
