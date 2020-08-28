package tree

// Search 二叉搜索树的查找函数
// 查到了返回该节点指针，没查到返回 nil
func Search(tree *BTree, value int) *TNode {
	if tree.root == nil {
		return nil
	}
	// 非递归查找
	// 指示访问节点
	p := tree.root
	for p != nil {
		if p.element > value {
			p = p.left
		} else if p.element < value {
			p = p.right
		} else {
			return p
		}
	}
	// 没查到
	return nil
}

// Insert 向一棵二叉搜索树中插入一个值
func Insert(tree *BTree, value int) bool {
	if tree.root == nil {
		return false
	}

	// 指示访问节点
	p := tree.root
	// 指示要插入的那个节点
	var beforeNode *TNode
	for p != nil {
		if p.element > value {
			beforeNode = p
			p = p.left
		} else if p.element < value {
			beforeNode = p
			p = p.right
		} else {
			// 树上已经有元素了
			return false
		}
	}
	// 构造节点
	var node TNode
	node.SetElement(value)
	if value > beforeNode.element {
		// 插在右边
		beforeNode.SetRight(&node)
	} else {
		// 插在左边
		beforeNode.SetLeft(&node)
	}
	return true
}

// Delete 在二叉搜索树中删除一个元素 （入口函数）
// 删除成功返回 true ，失败（节点元素不存在或者其他因素）返回 false
func Delete(tree *BTree, value int) bool {
	if tree.root == nil {
		return false
	}
	// 查找有咩有这个元素
	delNode := Search(tree, value)
	// 没有返回 false
	if delNode == nil {
		return false
	}
	// 树上存在此节点，删除
	_ = deleteNode(tree.root, value)
	return true
}

// deleteNode这是删除的递归函数
func deleteNode(node *TNode, value int) *TNode {
	if node == nil {
		return nil
	} else if value < node.element {
		// 向左，递归
		node.left = deleteNode(node.left, value)
	} else if value > node.element {
		// 向右，递归
		node.right = deleteNode(node.right, value)
		// 从这儿往下说明找到了删除的元素
	} else if node.left != nil && node.right != nil {
		// 左右孩子都存在
		// 使用右侧最小的替代
		temp := FindMin(node.right)
		node.SetElement(temp.element)
		node.right = deleteNode(node.right, temp.element)
	} else {
		// 就一个或者零个孩子
		if node.left == nil {
			// 左侧没有孩子
			// 直接放入右孩子
			node = node.right
		} else if node.right == nil {
			// 右侧没孩子
			// 直接放左侧的孩子
			node = node.left
		}
	}
	return node
}

// FindMin 查找子树最小的元素
func FindMin(node *TNode) *TNode {
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
}
