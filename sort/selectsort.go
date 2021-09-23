package sort

/*
选择排序（升序，降序同理）  （每轮排序都会“选择”一个最小的值）
步骤：
1、首先在未排序序列中找到最小元素，存放到排序序列的起始位置
2、再从剩余未排序元素中继续寻找最小元素，然后放到未排序序列起始位置
3、重复第二步，直到所有元素均排序完毕

时间复杂度：
最优O(n2)
平均O(n2)
最差O(n2)
空间复杂度：
O(1)
*/

func Selectsort(arr []int) {
	var minIndex int
	for i := 0; i < len(arr)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
