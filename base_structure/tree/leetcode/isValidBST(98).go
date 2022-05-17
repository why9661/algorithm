package leetcode

import "math"

/*
验证二叉搜索树
给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树

解：
根据二叉搜搜树的性质，对于树中的每个节点都有：其左子树的所有节点一定小于该节点，其右子树的所有节点一定大于该节点
*/

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min, max int) bool {
	// 基线条件
	if root == nil {
		return true
	}

	// 递归条件
	if root.Val > min && root.Val < max {
		return helper(root.Left, min, root.Val) && helper(root.Right, root.Val, max)
	}

	return false
}
