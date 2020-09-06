package tree

import (
	"fmt"
	"testing"
)

func TestGetDepth(t *testing.T) {
	// 这是树的入口
	var btree BTree
	/*
		 		1
			  /   \
			 2     3
			/ \   / \
		   4   5 6   7
			  /
			 8
	*/
	var node1, node2, node3, node4, node5, node6, node7, node8 TNode
	node1.SetElement(1)
	node2.SetElement(2)
	node3.SetElement(3)
	node4.SetElement(4)
	node5.SetElement(5)
	node6.SetElement(6)
	node7.SetElement(7)
	node8.SetElement(8)
	node1.SetLeft(&node2)
	node1.SetRight(&node3)
	node2.SetLeft(&node4)
	node2.SetRight(&node5)
	node3.SetLeft(&node6)
	node3.SetRight(&node7)
	node5.SetLeft(&node8)

	btree.InitTree(&node1)

	fmt.Println("循环遍历-树的高度是", btree.GetDepth())
	fmt.Println("循环遍历-树的宽度是", btree.GetWidth())
	fmt.Println("递归遍历-树的高度是", btree.GetDepthRec())
	fmt.Print("先序遍历的递归法：")
	btree.PreOrderByRec()
	fmt.Println()
	fmt.Print("先序遍历的循环法：")
	btree.PreOrderByCircle()
	fmt.Println()
	fmt.Print("中序遍历的递归法：")
	btree.InOrderByRec()
	fmt.Println()
	fmt.Print("中序遍历的循环法：")
	btree.InOrderByCircle()
	fmt.Println()
	fmt.Print("后序遍历的递归法：")
	btree.PostOrderByRec()
	fmt.Println()
	fmt.Print("后序遍历的循环法：")
	btree.PostOrderByCircle()
	fmt.Println()
	fmt.Print("层序遍历的循环法：")
	btree.LayerOrder()
	fmt.Println()
	fmt.Println("这棵树上共有多少节点？", btree.GetTreeNodeNumber())
	fmt.Println("这棵树是完全二叉树吗？", btree.IsCompleteBTree())
	fmt.Println("这棵树是满二叉树吗？", btree.IsFullBTree())

}

func TestBSTree(t *testing.T) {
	// 这是树的入口
	var btree BTree
	/*
		 		6
			  /   \
			 2     8
			/ \
		   1   5
			  /
			 3
	*/
	var node1, node2, node3, node5, node6, node8 TNode
	node1.SetElement(1)
	node2.SetElement(2)
	node3.SetElement(3)
	node5.SetElement(5)
	node6.SetElement(6)
	node8.SetElement(8)

	node6.SetLeft(&node2)
	node6.SetRight(&node8)
	node2.SetRight(&node5)
	node2.SetLeft(&node1)
	node5.SetLeft(&node3)

	btree.InitTree(&node6)

	// 待查元素
	ele := 3
	result := Search(&btree, ele)
	if result == nil {
		fmt.Println("查找失败，树上不存在：", ele)
	} else {
		fmt.Println("查找成功，树上存在：", result.element)
	}

	ok := Insert(&btree, 9)
	if ok {
		fmt.Print("插入9成功，层序遍历结果：")
		btree.LayerOrder()
		fmt.Println()
	} else {
		fmt.Println("插入失败")
	}
	ok = Insert(&btree, 4)
	if ok {
		fmt.Print("插入4成功，层序遍历结果：")
		btree.LayerOrder()
		fmt.Println()
	} else {
		fmt.Println("插入失败")
	}
	ok = Insert(&btree, 7)
	if ok {
		fmt.Print("插入7成功，层序遍历结果：")
		btree.LayerOrder()
		fmt.Println()
	} else {
		fmt.Println("插入失败")
	}
	fmt.Println("此时树高为：", btree.GetDepth())
	Delete(&btree, 2)
	fmt.Print("删除2之后，层序遍历结果：")
	btree.LayerOrder()
	fmt.Println()
}

func TestExpTree(t *testing.T) {
	// 后缀表达式
	exp := "42*22*3+*"
	expTree := InitExpTree(exp)
	value := expTree.Compute()
	fmt.Println("表达式的最终计算结果是：", value)
}

func TestAvlTree(t *testing.T) {
	// 初始化一棵二叉平衡树
	tree := InitAvlTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(6)
	tree.Insert(4)
	tree.LayerOrder()
	fmt.Println()
	tree.Insert(1)
	tree.LayerOrder()
	fmt.Println()
}
