package arr

/*
旋转数组
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数
*/

func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

//暴力解法
//func rotate(nums []int, k int)  {
//	k = k % len(nums)
//
//	for k > 0 {
//		temp := nums[len(nums)-1]
//		for i := len(nums)-1; i > 0; i-- {
//			nums[i] = nums[i-1]
//		}
//		nums[0] = temp
//		k--
//	}
//}
