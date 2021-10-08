package sort

/*
冒泡排序（升序，降序同理）  (每轮排序最大的元素会“冒到末尾”)
步骤：
1、比较相邻元素，如果第一个比第二个元素大则交换位置
2、对每一对相邻元素重复步骤1，一轮比较后最后的元素会是最大的元素
3、重复n轮后，排序完成

时间复杂度：
最优O(n)
平均O(n2)
最差O(n2)
空间复杂度：
O(1)
*/

func Bubblesort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
