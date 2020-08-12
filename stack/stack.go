package stack

import (
	"sync"
)

// Stack 这是一个栈的结构体
// 声明 Stack 变量之后需要执行 InitStack 操作初始化
type Stack struct {

	// 一个栈的主体 ：可以接受所有类型的 interface{} 切片
	items []interface{}

	// 一个读写互斥锁：用于保护栈的数S据，防止读写的冲突操作
	lock sync.RWMutex

	// 一个栈容量：uint 类型，如果为 0 则表示容量自动增长无上限
	cap uint

	// 一个栈高度指针
	height uint
}

// InitStack 创建栈s
// 将此方法绑定到 Stack 这个公开的栈的 struct 上面
// 传入一个参数 cap (无符号整型) ，表示栈的容量：
//      传入 0 表示需要创建的栈的容量无上限、自增长
//      传入 uint 型的非 0 整数表示创建的栈容量有限
func (s *Stack) InitStack(cap uint) {
	if cap != 0 {

		// 有容量限制
		s.items = make([]interface{}, cap)
		s.cap = cap

	} else {

		//无容量限制先默认栈的容量为 2 ，后面根据需要自动增长
		s.items = make([]interface{}, 2)
		s.cap = 0
	}
}

// IsEmpty 判断栈是否为空，为空则返回true，否则返回false。
func (s *Stack) IsEmpty() bool {
	if s.height == 0 {
		return true
	}

	return false
}

// Top 返回顶元素
func (s *Stack) Top() (interface{}, bool) {
	// 读之前要上锁
	s.lock.Lock()
	// 读完解锁
	defer s.lock.Unlock()

	if s.height > 0 {

		return s.items[s.height-1], true
	}

	return nil, false
}

// Push 进栈操作
func (s *Stack) Push(item interface{}) bool {
	// 写之前要上锁
	s.lock.Lock()
	// 写完解锁
	defer s.lock.Unlock()
	if s.cap == 0 {
		// 表示是无限制自增长的栈
		s.items = append(s.items, item)
		return true
	} else if s.height < s.cap {
		// 有容量限制，且未满
		s.items[s.height] = item
		// 栈高增 1
		s.height++
		return true
	} else {
		// 栈内已满
		return false
	}

}

// Pop 出栈操作
func (s *Stack) Pop() (interface{}, bool) {
	// 写之前要上锁
	s.lock.Lock()
	// 写完解锁
	defer s.lock.Unlock()
	if !s.IsEmpty() {
		// 栈非空
		item := s.items[s.height-1]
		s.items = s.items[:s.height-1]
		s.height--
		return item, true
	}
	//栈为空，不可以执行出栈操作
	return nil, false
}
