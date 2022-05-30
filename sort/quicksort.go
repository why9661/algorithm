package sort

/*
快速排序（升序）
步骤：
1、选定pivot
2、将大于pivot的元素放在pivot的右边
3、将小于pivot的元素放在pivot左边
4、分别对子序列重复步骤123
说明：
分治法（D&C）
认为长度为1的序列是有序的不用再进行排序，以此作为递归的基线条件（序列的左下标小于右下标）
每次都选取序列的第一个元素作为pivot，因此要先移动右下标（可以理解为由于在左边选取了pivot，左边就有位置可以插入）
当左下标于右下标相遇时，该次排序完成，并在此位置插入pivot
时间复杂度：
最好O(nlogn)
平均O(nlogn)
最坏O(n2)
空间复杂度：
O(nlogn)
*/

func QuickSort(arr []int) {
	left := 0
	right := len(arr) - 1
	quick(arr, left, right)
}

func quick(arr []int, left, right int) {
	// baseline
	if left < right {
		i, j := left, right
		pivot := arr[i]
		for i < j {
			for i < j && arr[j] >= pivot {
				j--
			}
			if i < j {
				arr[i] = arr[j]
			}

			for i < j && arr[i] <= pivot {
				i++
			}
			if i < j {
				arr[j] = arr[i]
			}
		}
		arr[i] = pivot
		quick(arr, left, i-1)
		quick(arr, i+1, right)
	}
}
