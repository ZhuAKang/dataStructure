package stack

import (
	"datastructure/linklist"
	"sync"
)

// Lstack 这是链表实现的栈的数据结构
type Lstack struct {
	// 栈的主体结构：链表
	stack linklist.LinkList

	// 一个读写互斥锁：用于保护栈的数S据，防止读写的冲突操作
	lock sync.RWMutex

	// 一个栈容量：int 类型，如果是复数则默认是无上限的
	cap int

	// 栈高度指针（就不用了，链表的长度就是栈的高度）
	// height uint

}

// InitStack 创建栈s
// 将此方法绑定到 Lstack 这个公开的栈的 struct 上面
// 传入一个参数 cap (整型) ，表示栈的容量：
//      传入 = 0 表示需要创建的栈的容量无上限、自增长
//      传入 > 0 的非 0 整数表示创建的栈容量有限
//      传入 < 0 的，参数错误，创建失败
func (s *Lstack) InitStack(cap int) bool {
	if cap >= 0 {
		s.stack.InitLinkList()
		s.cap = cap
		return true
	}
	return false
}

// IsEmpty 判断栈是否为空，为空则返回true，否则返回false。
func (s *Lstack) IsEmpty() bool {
	if s.stack.Length() == 0 {
		return true
	}
	return false
}

// Height 栈的高度
func (s *Lstack) Height() int {
	return s.stack.Length()
}

// Top 返回顶元素，也即返回链表的尾部
func (s *Lstack) Top() (interface{}, bool) {
	// 读之前要上锁
	s.lock.Lock()
	// 读完解锁
	defer s.lock.Unlock()

	if s.stack.Length() > 0 {
		node, ok := s.stack.SearchByID(s.stack.Length())
		if ok {
			return node.Data, true
		}
	}
	return nil, false
}

// Push 进栈操作
func (s *Lstack) Push(item interface{}) bool {
	// 写之前要上锁
	s.lock.Lock()
	// 写完解锁
	defer s.lock.Unlock()
	// 创建节点
	var node linklist.Node
	node.Data = item
	// 进栈之前先判断
	if s.cap == 0 || s.Height() < s.cap {
		// 无上限的栈或者未满，直接就入栈了（放在链表尾部）
		ok := s.stack.InsertIntoLinkList(&node, s.Height()+1)
		if ok {
			return true
		}
		return false
	}
	// 满了，放不下了
	return false

}

// Pop 出栈操作
func (s *Lstack) Pop() (interface{}, bool) {
	// 写之前要上锁
	s.lock.Lock()
	// 写完解锁
	defer s.lock.Unlock()
	if !s.IsEmpty() {
		// 栈非空，链表尾部节点删除（出栈）
		item, ok := s.stack.DeleteSP(s.Height())
		if ok {
			return item, true
		}
	}
	//栈为空，不可以执行出栈操作
	return nil, false
}
