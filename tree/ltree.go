package tree

import (
	"datastructure/stack"
)

// TNode 树的结点结构体
type TNode struct {
	// 左孩子指针
	left *TNode
	// 右孩子指针
	right *TNode
	// 数据域
	element int
}

// BTree 树的结构体
type BTree struct {
	// 根节点
	root *TNode
}

// InitTree 传入一个结点指针作为树的根节点，初始化一棵树
func (tree *BTree) InitTree(node *TNode) {
	tree.root = node
}

// GetRoot 获取树的根节点
func (tree *BTree) GetRoot() *TNode {
	return tree.root
}

// SetLeft 设置结点的左孩子
func (node *TNode) SetLeft(left *TNode) {
	node.left = left
}

// SetRight 设置结点的左孩子
func (node *TNode) SetRight(right *TNode) {
	node.right = right
}

// SetElement 设置结点的值
func (node *TNode) SetElement(element int) {
	node.element = element
}

// GetDepth 返回树的深度（高度）
func (tree *BTree) GetDepth() int {
	// 使用栈来完成，这一块的栈就可以用我们之前写的了
	var stack stack.Lstack
	// 初始化栈，栈容量自增长
	stack.InitStack(0)
	// 树的深度，初始化为栈的初始高度 0
	maxDepth := stack.Height()

	// 树的根节点入栈，作为栈底元素
	stack.Push(tree.root)

	for {
		// 获取栈顶元素，看看有没有左孩子，有就入栈（判断：更新高度），然后继续下一层循环
		// 如果没有左孩子，看看有没有右孩子，有就入栈（判断：更新高度），然后继续下一层循环
		// 如果左右都没有孩子，弹出栈顶元素，然后继续循环

		current, ok := stack.Top()

		// TODO: 逻辑有点问题，可能要有一个访问标志位，来判断左边的结点是否已经被访问了
		// 栈内仍有元素，栈顶仍然可以取
		if ok {
			// 类型断言，转换成 TNode 型
			currentNode, ok := current.(TNode)
			if ok {
				// 类型转换成功，可以下一步判断了
				// 左孩子存在，入栈
				if currentNode.left != nil {
					stack.Push(currentNode.left)
					// 判断高度，现在高的话高度就增加
					continue

					// 右孩子存在，入栈
				} else if currentNode.right != nil {
					stack.Push(currentNode.right)
					continue
				} else {
					// 左右都不存在，弹出当前栈顶并重来循环
					_, _ = stack.Pop()
					continue
				}
			}
		}
		// 循环退出条件：栈内空了
		if stack.IsEmpty() {
			break
		}
	}
	// 左右结点都访问过了就弹出
	return maxDepth
}

// IsCompleteBTree 判断当前二叉树是否为完全二叉树
func (tree *BTree) IsCompleteBTree() bool {
	return false
}

// IsFullBTree 判断当前二叉树是否为满二叉树
func (tree *BTree) IsFullBTree() bool {
	return false
}

// PrintBTree 打印显示当前的二叉树
func (tree *BTree) PrintBTree() {

}
