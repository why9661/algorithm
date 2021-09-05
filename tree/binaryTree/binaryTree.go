package binaryTree

import "fmt"

type node struct {
	left, right *node
	val         int
}

type Tree struct {
	root *node
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

//DFS深度优先遍历（前序遍历的非递归）
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
