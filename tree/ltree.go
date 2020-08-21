package tree

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
	return 0
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
