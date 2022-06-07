package sort

/*
归并排序
步骤：
1、将序列分为若干组，最终为每个数字为一组
2、将分好的若干分组两两合并，保证合并后的序列是有序的
3、重复步骤2，直到只剩下一组即排序完成

时间复杂度：
最好O(nlogn)
平均O(nlogn)
最坏O(nlogn)
空间复杂度：
O(n)
*/

func MergeSort(arr []int) []int {
	length := len(arr)
	// baseline
	if length < 2 {
		return arr
	}

	mid := length / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(arr1 []int, arr2 []int) []int {
	var result []int

	for len(arr1) != 0 && len(arr2) != 0 {
		if arr1[0] < arr2[0] {
			result = append(result, arr1[0])
			arr1 = arr1[1:]
		} else {
			result = append(result, arr2[0])
			arr2 = arr2[1:]
		}
	}

	if len(arr1) != 0 {
		result = append(result, arr1...)
	}

	if len(arr2) != 0 {
		result = append(result, arr2...)
	}

	return result
}

func MergeSortV2(arr []int) []int {
	// 先split然后merge
	length := len(arr)
	// baseline
	if length < 2 {
		return arr
	}

	mid := length / 2

	left := MergeSortV2(arr[:mid])
	right := MergeSortV2(arr[mid:])

	return mergeV2(left, right)
}

func mergeV2(a, b []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}

	if i < len(a) {
		result = append(result, a[i:]...)
	}

	if j < len(b) {
		result = append(result, b[j:]...)
	}

	return result
}
