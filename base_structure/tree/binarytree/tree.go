package binarytree

import "fmt"

type node struct {
	left, right *node
	val         int
}

type Tree struct {
	root *node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func New() *Tree {
	return new(Tree)
}

//前序遍历
func (t *Tree) PreOrder() {
	if t.root == nil {
		return
	}

	t.root.preOrder()
}

func (n *node) preOrder() {
	if n == nil {
		return
	}
	fmt.Println(n.val)
	n.left.preOrder()
	n.right.preOrder()
}

func preorderTraversal(root *TreeNode) []int {
	var preOrder func(n *TreeNode)
	var res []int

	preOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		res = append(res, n.Val)
		preOrder(n.Left)
		preOrder(n.Right)
	}

	preOrder(root)

	return res
}

//中序遍历
func (t *Tree) MidOrder() {
	if t.root == nil {
		return
	}

	t.root.midOrder()
}

func (n *node) midOrder() {
	if n == nil {
		return
	}
	n.left.midOrder()
	fmt.Println(n.val)
	n.right.midOrder()
}

func midorderTraversal(root *TreeNode) []int {
	var midOrder func(n *TreeNode)
	var res []int

	midOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		midOrder(n.Left)
		res = append(res, n.Val)
		midOrder(n.Right)
	}

	midOrder(root)

	return res
}

//后序遍历
func (t *Tree) PostOrder() {
	if t.root == nil {
		return
	}

	t.root.postOrder()
}

func (n *node) postOrder() {
	if n == nil {
		return
	}
	n.left.postOrder()
	n.right.postOrder()
	fmt.Println(n.val)
}

func postorderTraversal(root *TreeNode) []int {
	var postOrder func(n *TreeNode)
	var res []int

	postOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		postOrder(n.Left)
		postOrder(n.Right)
		res = append(res, n.Val)
	}

	postOrder(root)

	return res
}

//层序遍历(BFS广度优先遍历)
func (t *Tree) LevelOrder() {
	if t.root == nil {
		return
	}

	t.root.levelOrder()
}

func (n *node) levelOrder() {
	queue := make([]*node, 0)
	queue = append(queue, n)
	for len(queue) > 0 {
		n = queue[0]
		fmt.Println(n.val)
		queue = queue[1:]
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
}

// 前序遍历的非递归
func (t *Tree) PreOrderNR() {
	if t.root == nil {
		return
	}

	t.root.preOrderNR()
}

func (n *node) preOrderNR() {
	stack := make([]*node, 0)
	stack = append(stack, n)
	for len(stack) > 0 {
		n = stack[len(stack)-1]
		fmt.Println(n.val)
		stack = stack[:len(stack)-1]
		if n.right != nil {
			stack = append(stack, n.right)
		}
		if n.left != nil {
			stack = append(stack, n.left)
		}
	}
}

func preorderTraversalNR(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := make([]*TreeNode, 0)

	stack = append(stack, root)

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
	}

	return res
}

// 中序遍历的非递归
func (t *Tree) MidOrderNR() {

}

func (n *node) midOrderNR() {

}

func midorderTraversalNR(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}

	return res
}

// 后序遍历的非递归（在前序遍历的基础上，调换左右的入栈顺序，再将整个结果数据翻转）
func (t *Tree) PostOrderNR() {

}

func (n *node) postOrderNR() {

}

func postorderTraversalNR(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		if root.Left != nil {
			stack = append(stack, root.Left)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
