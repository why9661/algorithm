package sort

/*
希尔排序
希尔排序是把记录按下标的一定增量分组，对每组使用插入排序算法排序，随着增量逐渐减少，每组包含的关键词越来越多，当增量减至1时，整个文件恰被分成一组。
相较于直接插入排序，多了"增量"和"分组"。或者说直接插入排序是增量为1的希尔排序。
*/

func ShellSort(arr []int) {
	length := len(arr)

	if length <= 1 {
		return
	}

	// initial gap
	gap := length / 2

	for gap > 0 {
		// each group sorted by insert sort
		for i := 0; i < gap; i++ {
			shellSort(arr, i, gap)
		}
		gap /= 2
	}
}

func shellSort(arr []int, start, gap int) {
	var cur int
	var preIndex int

	for i := start + gap; i < len(arr); i += gap {
		cur = arr[i]
		preIndex = i - gap
		for preIndex >= 0 && cur < arr[preIndex] {
			arr[preIndex+gap] = arr[preIndex]
			preIndex -= gap
		}
		arr[preIndex+gap] = cur
	}
}
