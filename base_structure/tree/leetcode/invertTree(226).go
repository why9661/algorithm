package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left

	return root
}
