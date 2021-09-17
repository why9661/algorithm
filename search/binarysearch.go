package search

/*
二分查找（要求待查找的序列必须是有序的）
*/

func binarySearch(arr []int, target int) int {
	var low = 0
	var high = len(arr) - 1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] > target {
			high = mid - 1
		}
		if arr[mid] < target {
			low = mid + 1
		}
		if arr[mid] == target {
			return mid
		}
	}

	return -1
}
