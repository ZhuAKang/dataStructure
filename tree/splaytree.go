package tree

import (
	"datastructure/queue"
	"fmt"
)

// SPTNode 伸展树的结点结构
type SPTNode struct {
	// 数据域
	element int
	// 父、左孩子、右孩子指针
	fa, left, right *SPTNode
}

// SPTree 伸展树
type SPTree struct {
	// 根节点
	root *SPTNode
}

// InitSPTree 初始化一棵伸展树
func InitSPTree() *SPTree {
	var tree SPTree
	return &tree
}

// Splay 伸展操作
// 对树上的数据域的值为 element 的结点进行伸展操作
// 返回树
func Splay(element int, tree *SPTree) *SPTree {
	// 树为空
	if tree == nil {
		return nil
	}
	// 查找树上有没有这个结点
	result := search(element, tree.root)
	if result == nil {
		return nil
	}
	tree.root = splayN(result)
	return tree
}

// SplayN 对树上某一结点进行伸展
func splayN(node *SPTNode) *SPTNode {
	// 说明是根节点，不存在父节点，那就直接不用动了
	if node.fa == nil {
		return node
	}
	// 父节点存在，此时就要开始判断了
	father := node.fa
	// 父节点以上没有结点了，说明父节点是根节点
	// 执行一次左旋或者右旋
	if father.fa == nil {
		// 当前节点在父节点的左侧，执行一次右旋
		if father.left == node {
			node = singleRoateWithLeft(father)
		} else {
			// 当前节点子啊父节点的右侧，执行一次左旋
			node = singleRoateWithRight(father)
		}
		return node
	}
	// 父节点的父节点不为空，即当前状态是一字或之子（此处需要递归）
	grandFather := father.fa
	if grandFather.left == father && father.left == node {
		// 处于左左状态，一字旋转，向右旋
		father = singleRoateWithLeft(grandFather)
		node = singleRoateWithLeft(father)
	} else if grandFather.right == father && father.right == node {
		// 处于右右状态，一字旋转，向左旋
		father = singleRoateWithRight(grandFather)
		node = singleRoateWithRight(father)

	} else if grandFather.left == father && father.right == node {
		// 处于左右状态，之子旋转
		// 待上升节点先在父节点处左旋上升
		grandFather.left = singleRoateWithRight(father)
		node = singleRoateWithLeft(grandFather)
	} else {
		// 处于右左状态，之子旋转
		grandFather.right = singleRoateWithLeft(father)
		node = singleRoateWithRight(grandFather)
	}
	// 递归继续调用
	return splayN(node)

}

// singleRoateWithLeft 在当前节点（存在左子树）执行一次右旋操作
/*
		k2                k1
	   /  \    ---->     /  \
	  k1   Z            X    k2
	 /  \                   /  \
	X    Y                 Y    Z
*/
func singleRoateWithLeft(k2 *SPTNode) *SPTNode {
	k1 := k2.left
	k2.left = k1.right
	if k2.left != nil {
		k2.left.fa = k2
	}
	k1.right = k2
	k1.fa = k2.fa
	k2.fa = k1
	if k1.fa != nil {
		// 父亲节点存在，则应该更新其左或右孩子指针
		// 之前的父亲节点指向的是 k2 ,现在改成 k1
		if k1.fa.left == k2 {
			k1.fa.left = k1
		} else {
			k1.fa.right = k1
		}

	}
	return k1
}

// singleRoateWithRight 在当前节点（存在右子树）执行一次左旋操作
/*
		k2                k1
	   /  \    ---->     /  \
	  Z    k1           k2   Y
	      /  \         /  \
	     X    Y       Z    X
*/
func singleRoateWithRight(k2 *SPTNode) *SPTNode {
	k1 := k2.right
	k2.right = k1.left
	// 更新父节点信息
	if k2.right != nil {
		k2.right.fa = k2
	}
	k1.left = k2
	k1.fa = k2.fa
	k2.fa = k1
	if k1.fa != nil {
		// 父亲节点存在，则应该更新其左或右孩子指针
		// 之前的父亲节点指向的是 k2 ,现在改成 k1
		if k1.fa.left == k2 {
			k1.fa.left = k1
		} else {
			k1.fa.right = k1
		}

	}
	return k1
}

// search 在子树上查找数据域的值为 element 的结点，
// 查找到就返回该节点指针，查找失败返回nil
// 重要：这个函数只是查找有没有，这样方便在 伸展的时候进行操作，
// 并不是查找并伸展的操作，这个操作另外有个函数 Find
func search(element int, subtree *SPTNode) *SPTNode {
	if subtree == nil {
		return nil
	} else if subtree.element == element {
		return subtree
	} else if subtree.element > element {
		sL := search(element, subtree.left)
		if sL != nil {
			return sL
		}
	} else {
		sR := search(element, subtree.right)
		if sR != nil {
			return sR
		}
	}
	return nil
}

// LayerOrder 层序遍历，为了方便观察树的形状
func (tree *SPTree) LayerOrder() {
	if tree == nil {
		return
	}
	var queue queue.Queue
	_ = queue.InitQueue(10)
	// 根节点入队
	queue.InQueue(tree.root)
	for !queue.IsEmpty() {
		top, _ := queue.OutQueue()
		topNode := top.(*SPTNode)
		fmt.Print(topNode.element)
		if topNode.left != nil {
			queue.InQueue(topNode.left)
		}
		if topNode.right != nil {
			queue.InQueue(topNode.right)
		}
	}
	fmt.Println()
}

// PreOrderByRec 先序遍历的递归实现的入口函数,层序不能唯一确定树形，加一个
func (tree *SPTree) PreOrderByRec() {
	// 如果树空，直接返回或者做别的一些操作：提示什么都行
	if tree.root == nil {
		return
	}
	preOrder(tree.root)
}

// preOrderByRecursion 先序遍历的递归实现
func preOrder(node *SPTNode) {
	if node != nil {
		// 先触发当前结点操作
		fmt.Print(node.element)
		// 再转向操作左右子树
		preOrder(node.left)
		preOrder(node.right)
	}

}

// Find 在伸展树上查找元素 element
// 找到了就进行伸展并返回伸展后树指针和 true
// 没找到了返回树指针和 false
func Find(element int, tree *SPTree) (*SPTree, bool) {
	if tree == nil {
		return tree, false
	}
	// 查找有没有这个节点
	node := search(element, tree.root)
	if node == nil {
		return tree, false
	}
	tree = Splay(element, tree)
	return tree, true
}

// Insert 插入函数，向伸展树中插入一个值为 element 的结点
// 插入成功返回 true，插入失败返回 false
func (tree *SPTree) Insert(element int) bool {
	if tree == nil {
		return false
	}
	var node SPTNode
	node.element = element
	// 树中没根节点，插入的作为根节点
	if tree.root == nil {
		tree.root = &node
		return true
	}
	// 现在可以判断并插入了
	// 指示当前访问节点(插入位置)
	p := tree.root
	beforeNode := p.fa
	// 一直访问下去找到待插入位置
	for p != nil {
		// 查到了树上已有该节点了
		if p.element == element {
			return false
		}
		if p.element > element {
			beforeNode = p
			p = p.left
		} else {
			beforeNode = p
			p = p.right
		}
	}
	// 查到了插入位置，在 beforeNode 的孩子位置
	if beforeNode.element > element {
		// 插在左孩子位置
		beforeNode.left = &node
		node.fa = beforeNode
	} else {
		// 插在右孩子位置
		beforeNode.right = &node
		node.fa = beforeNode
	}
	tree.root = splayN(&node)
	return true
}

// Delete 删除伸展树上节点值为 element 的节点
// 删除成功返回 true ，删除失败返回false
func (tree *SPTree) Delete(element int) bool {
	node := search(element, tree.root)
	// 树上不存在待删除的节点，删除失败
	if node == nil {
		return false
	}
	father := node.fa
	// 左右孩子都不为空，则选择右边最小的元素作为替代
	if node.left != nil && node.right != nil {
		// 找到待换上去的节点 P （p 有两种可能，有右孩子或没有右孩子是叶子节点）
		p := node.right
		for p.left != nil {
			p = p.left
		}
		// 删除
		node.element = p.element
		if p.fa.left == p {
			p.fa.left = p.right
		}
		if p.fa.right == p {
			p.fa.right = p.right
		}
		// 删除成功，对父节点进行一次伸展
		if father != nil {
			tree = Splay(father.element, tree)
		}
	} else {
		// 有一个孩子或者零个孩子
		// 只有左孩子
		if node.left != nil {
			node.element = node.left.element
			node.left = nil
			// 删除成功，对父节点进行一次伸展
			if father != nil {
				tree = Splay(father.element, tree)
			}
		} else if node.right != nil {
			// 只有右孩子
			node.element = node.right.element
			node.right = nil
			// 删除成功，对父节点进行一次伸展
			if father != nil {
				tree = Splay(father.element, tree)
			}
		} else {
			// 没有孩子
			if father != nil {
				if father.left == node {
					father.left = nil
				} else {
					father.right = nil
				}
				// 删除成功，对父节点进行一次伸展
				tree = Splay(father.element, tree)
			} else {
				// 没有孩子也没有父节点，说明这树上就一个节点（根节点），且要删除它
				tree.root = nil
			}
		}
	}
	return true
}
