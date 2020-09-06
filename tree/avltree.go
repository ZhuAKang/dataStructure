package tree

import (
	"datastructure/queue"
	"fmt"
)

// avlTNode AVL树的结点
type avlTNode struct {
	// 数据域
	element int
	// 结点所处的高度
	height int8
	// 左孩子
	left *avlTNode
	// 右孩子
	right *avlTNode
}

// AvlTree AVL树
type AvlTree struct {
	root *avlTNode
}

// InitAvlTree AVL 树的初始化函数
func InitAvlTree() *AvlTree {
	var tree AvlTree
	return &tree
}

// Insert AVL 树的插入函数
func (tree *AvlTree) Insert(element int) bool {
	// 树为空，应该先构建根节点
	if tree.root == nil {
		var node avlTNode
		node.element = element
		node.height = 1
		tree.root = &node
	} else {
		// 树不为空，插入
		tree.root = insertInNode(tree.root, element)
	}
	return true
}

// 在当前节点下插入指定元素值的节点，并返回节点插入后的当前根节点（不一定是树的根节点）
func insertInNode(node *avlTNode, element int) *avlTNode {
	// 当访问到待插入的位置的时候，生成节点
	if node == nil {
		var insertNode avlTNode
		insertNode.element = element
		node = &insertNode
		node.height = 1
	} else if node.element > element {
		// 向左插入
		node.left = insertInNode(node.left, element)
		// 左右子树的高度差为 2，(左边比右边高2)说明此时树的平衡已经被打破
		if nodeHeight(node.left)-nodeHeight(node.right) == 2 {
			if element < node.left.element {
				// 插在了左--左
				node = node.singleRotateWithLeft()
			} else {
				// 插在了左--右
				node = node.doubleRotateWithLeft()
			}
		}
	} else if node.element < element {
		// 向右插入
		node.right = insertInNode(node.right, element)
		// 左右子树的高度差为 2，(右边比左边高2)说明此时树的平衡已经被打破
		if nodeHeight(node.right)-nodeHeight(node.left) == 2 {
			if element > node.right.element {
				// 插在了右--右
				node = node.singleRotateWithRight()
			} else {
				// 插在了右--左
				node = node.doubleRotateWithRight()
			}
		}
	}
	// else 即节点中已经存在这个值了，我们就什么都不做
	// 更新节点高度信息
	node.height = Max(nodeHeight(node.left), nodeHeight(node.right)) + 1
	return node
}

// singleRotateWithLeft 以 k2 为中心，将 k2 和其左孩子执行一次向右旋转
// 此函数只有当 k2 的左孩子存在的时候调用
// 然后在 k2 和他的左孩子之间执行一次旋转
// 并更新节点高度，返回新的当前子树根
/*
		k2                k1
	   /  \    ---->     /  \
	  k1   Z            X    k2
	 /  \                   /  \
	X    Y                 Y    Z
*/
func (k2 *avlTNode) singleRotateWithLeft() *avlTNode {
	// 旋转
	k1 := k2.left
	k2.left = k1.right
	k1.right = k2
	// 更新高度
	k2.height = Max(nodeHeight(k2.left), nodeHeight(k2.right)) + 1
	k1.height = Max(nodeHeight(k1.left), nodeHeight(k2)) + 1
	return k1
}

// singleRotateWithRight 以 k2 为中心，将 k2 和其右孩子执行一次向右旋转
// 此函数只有当 k2 的右孩子存在的时候调用
// 然后在 k2 和他的右孩子之间执行一次旋转
// 并更新节点高度，返回新的当前子树根
/*
		k2                k1
	   /  \    ---->     /  \
	  Z    k1           k2   Y
	      /  \         /  \
	     X    Y       Z    X
*/
func (k2 *avlTNode) singleRotateWithRight() *avlTNode {
	// 旋转
	k1 := k2.right
	k2.right = k1.left
	k1.left = k2
	// 更新高度
	k2.height = Max(nodeHeight(k2.left), nodeHeight(k2.right)) + 1
	k1.height = Max(nodeHeight(k1.right), nodeHeight(k2)) + 1
	return k1
}

// doubleRotateWithLeft 以 k3 为中心，在 k3 左孩子存在且左孩子的右孩子存在的情况下，进行双旋转
// 插入了左孩子的右子树导致的不平衡
// 做左--右双旋转，更新节点高度并返回根
/*
		k3            		k3				        k2
	   /  \   ----->  	  /	   \	------>       /    \
      k1   D          	 k2		D			     k1     k3
	 /  \             	/  \				    /  \   /  \
	A    k2            k1	C				   A    B C    D
		/  \          /  \
	   B    C        A    B
*/
func (k3 *avlTNode) doubleRotateWithLeft() *avlTNode {
	// 先以 k1 为中心旋 k2 上去
	k3.left = k3.left.singleRotateWithRight()
	// 再以 k3 为中心旋 k2 上去
	return k3.singleRotateWithLeft()
}

// doubleRotateWithRight 以 k3 为中心，在 k3 右孩子存在且右孩子的左孩子存在的情况下，进行双旋转
// 插入了右孩子的左子树导致的不平衡
// 做右--左双旋转，更新节点高度并返回根
/*
		k3                    k3                        k2
	   /  \   ----->        /    \		----->        /    \
      A   k1               A     k2                  k3     k1
	     /  \                    /  \               /  \   /  \
	    k2   D                  B    k1            A    B C    D
	   /  \                         /  \
	  B    C                       C    D
*/
func (k3 *avlTNode) doubleRotateWithRight() *avlTNode {
	// 先以 k1 为中心，旋 k2 上去
	k3.right = k3.right.singleRotateWithLeft()
	// 再以 k3 为中心旋 k2 上去
	return k3.singleRotateWithRight()

}

// Max 返回两个值中的大者
func Max(n, m int8) int8 {
	if n > m {
		return n
	}
	return m
}

// nodeHeight 返回节点的高度（防止空指针异常）
func nodeHeight(node *avlTNode) int8 {
	if node == nil {
		return 0
	}
	return node.height
}

// LayerOrder 层序遍历，为了更好的看见树的结果
func (tree *AvlTree) LayerOrder() {
	// 如果树空，直接返回或者做别的一些操作：提示什么都行
	if tree.root == nil {
		return
	}
	// 声明一个队列
	var queue queue.Queue
	queue.InitQueue(12)
	// 根节点入队
	queue.InQueue(tree.root)
	// 队列非空的时候一直循环
	for !queue.IsEmpty() {
		top, _ := queue.OutQueue()
		topNode := top.(*avlTNode)
		fmt.Print(topNode.element)
		if topNode.left != nil {
			queue.InQueue(topNode.left)
		}
		if topNode.right != nil {
			queue.InQueue(topNode.right)
		}
	}
}
