package tree

import (
	"datastructure/stack"
	"strconv"
)

// ETNode 表达式树上的结点
type ETNode struct {
	// 数据域，使用 string 存储表达式值和操作符等
	element byte
	// 左指针
	left *ETNode
	// 右指针
	right *ETNode
}

// ExpTree 表达式树
type ExpTree struct {
	// 树的根节点
	root *ETNode
}

// InitExpTree 表达式树的初始化函数，传入 string 类型的表达式，返回表达式树的指针
func InitExpTree(expression string) *ExpTree {
	if expression == "" {
		return nil
	}
	// 对传入的字符串按照 byte 一个字节一个字节的读出来判断，
	// 数字就放入一个栈，操作符就弹出两个栈顶组成新的子树再放入子树的根
	var tNodeStack stack.Lstack
	tNodeStack.InitStack(0)
	for i := 0; i < len(expression); i++ {
		var node ETNode
		node.element = expression[i]
		switch expression[i] {
		// 如果是数字则入栈
		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
			tNodeStack.Push(node)
		case 42, 43, 45, 47:
			// 出栈两个子树然后组成新的子树，再放入栈内
			rightNumber, _ := tNodeStack.Pop()
			right := rightNumber.(ETNode)
			leftNumber, _ := tNodeStack.Pop()
			left := leftNumber.(ETNode)
			node.left = &left
			node.right = &right
			tNodeStack.Push(node)
		}
	}
	// 此时栈顶即为最后的根
	tree, _ := tNodeStack.Pop()
	etree := tree.(ETNode)
	var expTree ExpTree
	expTree.root = &etree
	return &expTree
}

// Compute 计算表达式树的最终数值的入口函数
// 可以采用递归遍历的方式去计算
func (expTree *ExpTree) Compute() int {
	if expTree == nil {
		return 0
	}
	return computeNode(expTree.root)
}

// computeNode 采用中序遍历的递归方式去求表达式树的值
func computeNode(node *ETNode) int {
	// 非叶子节点
	if node != nil {
		if node.left != nil && node.right != nil {
			left := computeNode(node.left)
			right := computeNode(node.right)
			switch node.element {
			case 42:
				return left * right
			case 43:
				return left + right
			case 45:
				return left - right
			case 47:
				return left / right
			}
		} else {
			// 字符串单个读出来是 byte，但是并没有找到单 byte（内存ASCII码）转 int 的
			// 所以就先转为字符串，再从字符串转回数值
			str := string(node.element)
			value, _ := strconv.Atoi(str)
			return value
		}
	}
	// 叶子节点
	return 0
}
