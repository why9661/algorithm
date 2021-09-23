package sort

/*
插入排序
步骤：
1、认为i前面都是有序序列，将i与前面序列元素依次比较，比i大则向后移，然后将i插入到相应位置
2、重复步骤1

时间复杂度：
最好O(n)
平均O(n2)
最坏O(n2)
空间复杂度：
O(1)
*/

func Insertsort(arr []int) {
	var preIndex int
	var cur int
	for i := 1; i < len(arr); i++ {
		cur = arr[i]
		preIndex = i - 1
		for preIndex >= 0 && cur < arr[preIndex] {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = cur
	}
}
