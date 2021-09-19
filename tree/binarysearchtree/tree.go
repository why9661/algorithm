package binarysearchtree

import "fmt"

/*
二叉搜索树or二叉查找树

特点：
对二叉搜索树中的每一个节点都有
1、若其左子树不为空，那么左子树中每个节点的值都不大于该节点值
2、若其右子树不为空，那么右子树中每个节点的值都不大于该节点值
*/

type Node struct {
	Val   int
	Times int
	Left  *Node
	Right *Node
}

type BSTree struct {
	Root *Node
}

func New() *BSTree {
	return new(BSTree)
}

// 添加值
func (t *BSTree) Add(value int) {
	if t.Root == nil {
		newNode := &Node{Val: value, Times: 1}
		t.Root = newNode
		return
	}

	t.Root.add(value)
}

func (n *Node) add(value int) {
	if value < n.Val {
		if n.Left == nil {
			newNode := &Node{Val: value, Times: 1}
			n.Left = newNode
		} else {
			n.Left.add(value)
		}
	} else if value > n.Val {
		if n.Right == nil {
			newNode := &Node{Val: value, Times: 1}
			n.Right = newNode
		} else {
			n.Right.add(value)
		}
	} else if value == n.Val {
		n.Times++
	}
}

// 查找树中最大值
func (t *BSTree) FindMax() *Node {
	if t.Root == nil {
		return nil
	}

	return t.Root.findMax()
}

func (n *Node) findMax() *Node {
	if n.Right == nil {
		return n
	} else {
		return n.Right.findMax()
	}
}

// 查找树中最小值
func (t *BSTree) FindMin() *Node {
	if t.Root == nil {
		return nil
	}

	return t.Root.findMin()
}

func (n *Node) findMin() *Node {
	if n.Left == nil {
		return n
	} else {
		return n.Left.findMin()
	}
}

// 查找指定节点
func (t *BSTree) Find(value int) *Node {
	if t.Root == nil {
		return nil
	}

	return t.Root.find(value)
}

func (n *Node) find(value int) *Node {
	if value == n.Val {
		return n
	} else if value < n.Val {
		if n.Left == nil {
			return nil
		} else {
			return n.Left.find(value)
		}
	} else if value > n.Val {
		if n.Right == nil {
			return nil
		} else {
			return n.Right.find(value)
		}
	}

	return nil
}

// 查找指定结点的父结点
func (t *BSTree) FindParent(value int) *Node {
	if t.Root == nil {
		return nil
	}

	if t.Root.Val == value {
		return nil
	}

	return t.Root.findParent(value)
}

func (n *Node) findParent(value int) *Node {
	if value < n.Val {
		if n.Left == nil {
			return nil
		} else {
			if n.Left.Val == value {
				return n
			} else {
				return n.Left.findParent(value)
			}
		}
	} else if value > n.Val {
		if n.Right == nil {
			return nil
		} else {
			if n.Right.Val == value {
				return n
			} else {
				return n.Right.findParent(value)
			}
		}
	}

	return nil
}

// 删除值
// 删除有三种情况：
// 1、删除的结点没有子树即叶子结点，直接删除
// 2、删除的结点有一个子树，用子树替换该结点
// 3、删除的结点有两个子树，用左子树的最大结点（或右子树的最小结点）替换该结点
func (t *BSTree) Delete(value int) {
	//空树直接返回
	if t.Root == nil {
		return
	}

	node := t.Find(value)
	//要删除的结点不存在，直接返回
	if node == nil {
		return
	}

	parent := t.FindParent(value)

	//一：删除的节点没有子树
	if node.Left == nil && node.Right == nil {
		//判断parent来确定删除的节点是否为根节点
		if parent == nil {
			//是根节点的话，将根节点置为nil即可
			t.Root = nil
			return
		} else if parent.Left == node {
			//要删除的节点是父亲节点的左孩子，那么将父亲节点的左子树置为nil即可
			parent.Left = nil
			return
		} else if parent.Right == node {
			//要删除的节点是父亲节点的右孩子，那么将父亲节点的右子树置为nil即可
			parent.Right = nil
			return
		}
	}

	//二、删除的节点有一个子树
	//删除的节点有左子树
	if node.Left != nil && node.Right == nil {
		//判断parent来确定要删除的节点是否为根节点
		if parent == nil {
			//是根节点的话那么直接将根节点置为要删除节点的左子树
			t.Root = node.Left
			return
		} else if parent.Left == node {
			//要删除的节点是父亲节点的左孩子，那么将父亲节点的左子树置为要删除节点的左子树即可
			parent.Left = node.Left
			return
		} else if parent.Right == node {
			//要删除的节点是父亲节点的右孩子，那么将父亲节点的右子树置为要删除节点的左子树即可
			parent.Right = node.Left
			return
		}
	}
	//删除的节点只有右子树，同理
	if node.Left == nil && node.Right != nil {
		//判断parent来确定要删除的节点是否为根节点
		if parent == nil {
			//是根节点的话那么直接将根节点置为要删除节点的左子树
			t.Root = node.Right
			return
		} else if parent.Left == node {
			//要删除的节点是父亲节点的左孩子，那么将父亲节点的左子树置为要删除节点的右子树即可
			parent.Left = node.Right
			return
		} else if parent.Right == node {
			//要删除的节点是父亲节点的右孩子，那么将父亲节点的右子树置为要删除节点的右子树即可
			parent.Right = node.Right
			return
		}
	}

	//三、删除的节点有左子树和右子树
	if node.Right != nil && node.Left != nil {
		//找到”替换节点“（这里使用右子树的最小节点）
		minNode := node.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}

		//删除该节点
		t.Delete(minNode.Val)

		//判断要删除的节点是否为根节点
		if parent == nil {
			//是根节点，用”替换节点“覆盖根节点
			t.Root.Val = minNode.Val
			t.Root.Times = minNode.Times
		} else {
			//不是根节点
			node.Val = minNode.Val
			node.Times = minNode.Times
		}
	}
}

// 顺序遍历(由于二叉搜索树的特性，使用中序遍历即可得到递增序列)
func (t *BSTree) MidOrder() {
	t.Root.midOrder()
}

func (n *Node) midOrder() {
	if n == nil {
		return
	}
	n.Left.midOrder()
	fmt.Println(n.Val)
	n.Right.midOrder()
}
