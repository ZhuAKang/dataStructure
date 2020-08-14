package linklist

import (

	// "C"
	"fmt"
	"sync"
)

// Node 为链表节点的数据结构
// 这里面 Next 和 Data 做导出
// 只是为了下面在实现队列或栈的链表实现的时候可以直接用这里面的
type Node struct {
	// 指针域
	Next *Node
	// 数据域
	Data interface{}
}

// LinkList 为链表的数据结构
type LinkList struct {
	// 链表的头节点
	head *Node
	// 链表的尾部
	tail *Node
	// 链表的长度
	len int
	// 一个读写互斥锁：用于保护数据，防止读写的冲突操作
	lock sync.RWMutex
}

// TODO: 下面所有的函数有很大的问题没有考虑到链表的尾部
// InitLinkList 初始化单链表
func (linkList *LinkList) InitLinkList() {
	var node Node
	// 这一块必须是创建的 Node 实例而不是 *Node。这样的话后面报 nil pointer 错误
	// 这一块是必须先创建结构体的变量实体，创建指针的话不指向指定结构体变量实体的话就是nil pointer错误
	linkList.lock.Lock()
	defer linkList.lock.Unlock()
	linkList.head = &node
	linkList.tail = &node
	linkList.len = 0
}

// // InitCirclList 初始化循环链表
// func InitCirclList() *Node {

// 	return nil
// }

// Length 查询链表长度
func (linkList *LinkList) Length() int {
	return linkList.len
}

// SerachInLinkList 单链表上查询（只查离链表头最近的一个的前一个结点）
// 传入一个 interface 类型的值，查询链上是否有此值
// 有的话返回这个值所在结点的前一个结点的指针及 true ，否则返回 nil 和 false
func (linkList *LinkList) SerachInLinkList(item interface{}) (*Node, bool) {
	for i := 0; i < linkList.len; i++ {
		currentNode := linkList.head
		// 都是 interface 类型，比较值相等不能使用 “ == ” 了
		// 可以使用 reflect 包下的 func DeepEqual(a1, a2 interface{}) bool
		// TODO: 比较这一块有问题，使用上面的也不太行，可能要用反射包里的东西
		// reflect.ValueOf(currentNode.Next.Data) == reflect.ValueOf(item)
		if currentNode.Next != nil &&
			currentNode.Next.Data == item {
			return currentNode, true
		}
	}
	return nil, false
}

// DeleteInLinkList 单链表上的删除（只删除离链表头最近的一个）
// TODO: 删除的逻辑要注意，注意删除的是链表尾部
func (linkList *LinkList) DeleteInLinkList(item interface{}) bool {
	// 修改前要加锁
	linkList.lock.Lock()
	defer linkList.lock.Unlock()
	beforeDeleteNode, _ := linkList.SerachInLinkList(item)
	if beforeDeleteNode.Next != nil {
		// 获取待删除结点指针
		deleteNode := beforeDeleteNode.Next
		// 从链上删除
		beforeDeleteNode.Next = deleteNode.Next
		// 删除的是链表尾
		if deleteNode == linkList.tail {
			linkList.tail = beforeDeleteNode
		}

		// TODO: 此处可能需要内存释放，考虑使用 cgo
		linkList.len--
		return true
	}
	return false

}

// InsertIntoLinkList 单链表上的插入操作
// 参数: node 为待插入的节点；position 为节点插入的位置，从0开始
// 返回值: true 表示插入成功，否则插入失败。失败可能是由于插入的位置不对
// 插入位置，从 1 开始，0 位置放的是头结点
func (linkList *LinkList) InsertIntoLinkList(node *Node, position int) bool {
	// 修改前要加锁
	linkList.lock.Lock()
	defer linkList.lock.Unlock()

	// 插入链表尾部
	if position == linkList.Length()+1 {
		linkList.tail.Next = node
		// 链表尾部必须指向后面（重要）
		linkList.tail = node
		linkList.len++
		return true
	} else {
		if position > 0 && position <= linkList.Length() {

			beforInsertNode := linkList.head
			// 找到待插入的前一个结点
			for i := 1; i < position; i++ {
				beforInsertNode = beforInsertNode.Next
			}
			// 断链插入
			node.Next = beforInsertNode.Next
			beforInsertNode.Next = node
			linkList.len++
			return true
		}

	}
	return false
}

// ShowList 打印链表
func (linkList *LinkList) ShowList() {
	curNode := linkList.head
	for i := 0; i < linkList.len; i++ {
		curNode = curNode.Next
		fmt.Println(curNode.Data)
	}
}
