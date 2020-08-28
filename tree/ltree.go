package tree

import (
	"datastructure/queue"
	"datastructure/stack"
	"fmt"
	"math"
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
// TODO: 还有递归的方法，待定
func (tree *BTree) GetDepth() int {
	if tree.root == nil {
		return 0
	}

	// 使用栈来完成，这一块的栈就可以用我们之前写的了
	var stack stack.Lstack
	// 初始化栈，栈容量自增长
	stack.InitStack(0)
	// 树的深度，初始化为栈的初始高度 0
	maxDepth := stack.Height()

	// 树的根节点入栈，作为栈底元素
	stack.Push(tree.root)

	// 左右孩子访问标志位，指向刚刚出栈的那个节点
	var accessed *TNode

	for {

		current, ok := stack.Top()

		// 栈内仍有元素，栈顶仍然可以取
		if ok {
			// 类型断言，转换成 *TNode 型
			currentNode, ok := current.(*TNode)
			if ok {
				// 类型转换成功，可以下一步判断了
				// 左孩子存在且左右孩子均未被访问过，入栈（左孩子未被访问入栈可以理解，右孩子没被访问主要是后序遍历思想，先左后右最后中间，
				// 当刚刚出栈的节点是当前栈顶的右孩子的时候，说明左孩子已经访问过了，不需要再访问了）
				if currentNode.left != nil && accessed != currentNode.left && accessed != currentNode.right {
					stack.Push(currentNode.left)
					// 判断高度，现在高的话高度就增加
					if maxDepth < stack.Height() {
						maxDepth = stack.Height()
					}
					continue
				} else if currentNode.right != nil && accessed != currentNode.right {
					// 右孩子存在且未被访问，入栈
					stack.Push(currentNode.right)
					// 判断高度，现在高的话高度就增加
					if maxDepth < stack.Height() {
						maxDepth = stack.Height()
					}
					continue
				} else {
					// 左右都不存在，弹出当前栈顶并重来循环
					_, _ = stack.Pop()
					// 标记当前出栈的节点
					accessed = currentNode
					continue
				}
			}
		}
		// 循环退出条件：栈内空了或者上面栈的 Top 没有元素了，又或者类型转换失败了（栈内存的不是*TNode 类型）
		if stack.IsEmpty() || !ok {
			break
		}
	}
	// 左右结点都访问过了就弹出
	return maxDepth
}

// GetDepthRec 二叉树的高度的递归实现的入口函数
func (tree *BTree) GetDepthRec() int {
	if tree.root == nil {
		return 0
	}
	return GetDepthByRecursion(tree.root)
}

// GetDepthByRecursion 二叉树的高度的递归实现
func GetDepthByRecursion(node *TNode) int {
	if node != nil {
		left := GetDepthByRecursion(node.left)
		right := GetDepthByRecursion(node.right)
		if left < right {
			return right + 1
		}
		return left + 1
	}
	return 0
}

// GetWidth 二叉树的宽度
func (tree *BTree) GetWidth() int {
	if tree.root == nil {
		return 0
	}
	// 声明一个队列，这里使用的是之前写的使用链表实现的队列
	var queue queue.Queue
	// 初始化队列长度，因为是链表实现的队列，且队里只存这一层加上少部分上一层的节点，所以不用设太大
	queue.InitQueue(12)

	queue.InQueue(tree.root)
	// count 每层节点数，width 宽度
	count := 1
	width := 1
	// 队列非空
	for !queue.IsEmpty() {
		// 临时保存下一层的节点数
		var size = 0
		for i := 0; i < count; i++ {
			// 出栈
			curNode, _ := queue.OutQueue()
			node := curNode.(*TNode)
			if node.left != nil {
				queue.InQueue(node.left)
				size++
			}
			if node.right != nil {
				queue.InQueue(node.right)
				size++
			}
		}
		// 下一层没有节点了，退出
		if size == 0 {
			break
		}
		if size > width {
			width = size
		}
		// 重新开始下一层了
		count = size
	}
	return width
}

// --------------------树形判断--------------------------

// IsCompleteBTree 判断当前二叉树是否为完全二叉树(使用层序遍历思想)
func (tree *BTree) IsCompleteBTree() bool {
	// 树为空则直接退出
	if tree.root == nil {
		return false
	}
	// 声明队列并进行初始化
	var queue queue.Queue
	queue.InitQueue(12)

	// 根节点入队
	queue.InQueue(tree.root)

	// 开始循环判断(只要队列非空，就一直循环下去直到队列为空或者触发退出机制)
	for !queue.IsEmpty() {
		curNode, _ := queue.OutQueue()
		node := curNode.(*TNode)
		// 左右孩子节点都存在
		if node.left != nil && node.right != nil {
			queue.InQueue(node.left)
			queue.InQueue(node.right)
		} else if node.left == nil && node.right != nil {
			// 左孩子为空右孩子不为空，则不是完全二叉树
			return false
		} else {
			// 左右孩子都为空或者左孩子为空右孩子不为空，
			// 则队列中这个节点之后的所有节点都是叶子节点，才能使得这棵树是完全二叉树
			for !queue.IsEmpty() {
				curNode, _ = queue.OutQueue()
				node := curNode.(*TNode)
				if node.left != nil || node.right != nil {
					return false
				}
			}
		}
	}
	return true
}

// IsFullBTree 判断当前二叉树是否为满二叉树(国内定义的满二叉树)
// 得到树高和节点的个数，判断是否满足：2^(树高) - 1 = 节点个数
func (tree *BTree) IsFullBTree() bool {
	if tree.root == nil {
		return true
	}
	// 如果2^(树高) - 1 != 节点个数，则不是满二叉树
	if math.Pow(2, float64(tree.GetDepth()))-1 != float64(tree.GetTreeNodeNumber()) {
		return false
	}
	return true
}

// GetTreeNodeNumber 遍历计算树上的节点总数(使用层序)
func (tree *BTree) GetTreeNodeNumber() int {
	if tree.root == nil {
		return 0
	}
	var count int
	// 声明队列并进行初始化
	var queue queue.Queue
	queue.InitQueue(12)

	// 根节点入队
	queue.InQueue(tree.root)
	count++
	for !queue.IsEmpty() {
		curNode, _ := queue.OutQueue()
		node := curNode.(*TNode)
		if node.left != nil {
			queue.InQueue(node.left)
			count++
		}
		if node.right != nil {
			queue.InQueue(node.right)
			count++
		}
	}
	return count
}

// PrintBTree 打印显示当前的二叉树
func (tree *BTree) PrintBTree() {

}

// ------先序遍历的实现------

// PreOrderByCircle 先序遍历的循环实现
func (tree *BTree) PreOrderByCircle() {
	// 如果树空，直接返回或者做别的一些操作：提示什么都行
	if tree.root == nil {
		return
	}
	// 声明一个栈，树根入栈，然后开始循环（条件：栈非空）
	// 弹出栈顶元素，进行操作，然后将右孩子与左孩子依次入栈（如果右或者左存在的话）
	var stack stack.Lstack
	stack.InitStack(0)

	// 树根入栈
	stack.Push(tree.root)
	for !stack.IsEmpty() {
		// 出栈并进行相应操作
		topNode, _ := stack.Pop()
		node := topNode.(*TNode)
		fmt.Print(node.element)
		// 右孩子存在的话入栈
		if node.right != nil {
			stack.Push(node.right)
		}
		// 左孩子存在的话入栈
		if node.left != nil {
			stack.Push(node.left)
		}
	}
	return
}

// PreOrderByRec 先序遍历的递归实现的入口函数
func (tree *BTree) PreOrderByRec() {
	// 如果树空，直接返回或者做别的一些操作：提示什么都行
	if tree.root == nil {
		return
	}
	preOrderByRecursion(tree.root)
}

// preOrderByRecursion 先序遍历的递归实现
func preOrderByRecursion(node *TNode) {
	if node != nil {
		// 先触发当前结点操作
		fmt.Print(node.element)
		// 再转向操作左右子树
		preOrderByRecursion(node.left)
		preOrderByRecursion(node.right)
	}

}

// ------中序遍历的实现------

// InOrderByCircle 中序遍历的循环实现
func (tree *BTree) InOrderByCircle() {
	// 如果树空，直接返回或者做别的一些操作：提示什么都行
	if tree.root == nil {
		return
	}
	var stack stack.Lstack
	stack.InitStack(0)

	// TODO: 要设置一个访问标志位，看看出栈的那个结点是从左边回来的还是从右边回来的
	var beforeNode *TNode
	// 指示访问的当前位置
	p := tree.root

	stack.Push(tree.root)

	for !stack.IsEmpty() {

		// 左孩子存在就入栈
		if p.left != nil && p.left != beforeNode {
			stack.Push(p.left)
			p = p.left
		} else {
			// 没有左孩子，出栈并执行相关操作
			// 先保存之前出栈的结点
			beforeNode = p
			// 栈顶出栈并执行相关操作
			topNode, _ := stack.Pop()
			p = topNode.(*TNode)
			fmt.Print(p.element)
			if p.right != nil {
				p = p.right
				stack.Push(p)
			}

		}

	}
}

// InOrderByRec 中序遍历的递归实现的入口函数
func (tree *BTree) InOrderByRec() {
	// 如果树空，直接返回或者做别的一些操作：提示什么都行
	if tree.root == nil {
		return
	}
	inOrderByRecursion(tree.root)
}

// inOrderByRecursion 中序遍历的递归实现
func inOrderByRecursion(node *TNode) {
	if node != nil {
		// 先转向操作左子树
		inOrderByRecursion(node.left)
		// 再触发当前结点操作
		fmt.Print(node.element)
		// 再转向操作左子树
		inOrderByRecursion(node.right)
	}
}

// ------后序遍历的实现------

// PostOrderByCircle 后序遍历的循环实现
func (tree *BTree) PostOrderByCircle() {
	// 如果树空，直接返回或者做别的一些操作：提示什么都行
	if tree.root == nil {
		return
	}
	// 声明个栈，用来存结点
	var stack stack.Lstack
	stack.InitStack(0)
	// 访问标识，看看刚刚出栈的结点是不是栈顶的孩子
	var beforeNode *TNode

	// 根节点先入栈
	stack.Push(tree.root)
	for !stack.IsEmpty() {
		// 获取栈顶元素，有左或右孩子，则左右孩子入栈，其中右孩子先入栈
		topNode, _ := stack.Top()
		node := topNode.(*TNode)
		// 当前栈顶结点的左孩子存在且刚刚出栈的结点不是此节点的左右孩子，证明左孩子未被访问过
		if node.left != nil && beforeNode != node.right && beforeNode != node.left {
			stack.Push(node.left)
		} else if node.right != nil && beforeNode != node.right {
			// 栈顶结点的右孩子存在且刚刚出栈的结点不是这个结点的右孩子，就证明当前栈顶结点的右孩子还未被访问过
			stack.Push(node.right)
		} else {
			// 左右孩子均不存在或者左右孩子以及都被访问过了
			curNode, _ := stack.Pop()
			node = curNode.(*TNode)
			beforeNode = node
			fmt.Print(node.element)
		}
	}
}

// PostOrderByRec 后序遍历的递归实现的入口函数
func (tree *BTree) PostOrderByRec() {
	// 如果树空，直接返回或者做别的一些操作：提示什么都行
	if tree.root == nil {
		return
	}
	postOrderByRecursion(tree.root)
}

// postOrderByRecursion 后序遍历的递归实现
func postOrderByRecursion(node *TNode) {
	if node != nil {
		// 先转向操作左右子树
		postOrderByRecursion(node.left)
		postOrderByRecursion(node.right)
		// 再触发当前结点操作
		fmt.Print(node.element)
	}
}

// ------层序遍历的实现------

// LayerOrder 层序遍历的循环实现（好像没有递归的）
func (tree *BTree) LayerOrder() {
	// 如果树空，直接返回或者做别的一些操作：提示什么都行
	if tree.root == nil {
		return
	}
	// 声明一个队列，这里使用的是之前写的使用链表实现的队列
	var queue queue.Queue
	// 初始化队列长度，因为是链表实现的队列，且队里只存这一层加上少部分上一层的节点，所以不用设太大
	queue.InitQueue(7)

	queue.InQueue(tree.root)
	// 队列非空的时候一直循环
	for !queue.IsEmpty() {
		// 队头元素出队
		out, _ := queue.OutQueue()
		node := out.(*TNode)
		fmt.Print(node.element)
		if node.left != nil {
			queue.InQueue(node.left)
		}
		if node.right != nil {
			queue.InQueue(node.right)
		}
	}

}
