package queue

// Linked list Queue
// 队列的链表实现(　TODO:　待完成链表之后再来完成这一块)
import (
	"datastructure/linklist"
	"sync"
)

// Queue 这是简单队列的结构体
// 声明结构体变量之后需要执行 InitQueue 方法初始化队列
type Queue struct {

	// 队列主体
	queue linklist.LinkList

	// 队首队尾的指示可以不用了，单链表有 head 和 tail
	// front, rear int

	// 队列的容量
	cap int

	// 一个读写互斥锁：用于保护栈的数S据，防止读写的冲突操作
	lock sync.RWMutex
}

// InitQueue 初始化队列
// 传入 int 类型的队列的长度
func (queue *Queue) InitQueue(cap int) bool {

	if cap > 1 {
		queue.queue.InitLinkList()
		queue.cap = cap
		return true
	}
	// 传入的 cap <= 1 ，理论上这个队列是不可能存在的
	return false
}

// OutQueue 出队操作
// 队列有值可以返回就返回出队的值和true
// 队列不满足出队条件就返回 nil 和 false
func (queue *Queue) OutQueue() (interface{}, bool) {
	// 读之前要上锁
	queue.lock.Lock()
	// 读完解锁
	defer queue.lock.Unlock()
	// 直接调用
	if queue.queue.Length() == 0 {
		return nil, false
	}
	return queue.queue.DeleteSP(1)
}

// InQueue 入队操作
func (queue *Queue) InQueue(item interface{}) bool {

	// 读之前要上锁
	queue.lock.Lock()
	// 读完解锁
	defer queue.lock.Unlock()
	// 入队之前要判断队是否满
	if queue.queue.Length() < queue.cap {
		var inNode linklist.Node
		inNode.Data = item
		// 入队并返回入队操作的结果
		return queue.queue.InsertIntoLinkList(&inNode, queue.queue.Length()+1)
	}
	return false
}

// IsEmpty 判队空操作
func (queue *Queue) IsEmpty() bool {
	// 队首、尾指针一样的时候说明队列为空
	if queue.queue.Length() == 0 {
		return true
	}
	return false
}

// TODO: 有毛病了，读不出来啊，因为在 linkList 里面，头是隐藏的信息
// 也可以改成导出类型，不过这就得不偿失了
// 倒是可以在 link_list 里再写一个函数用来导出指定位置的结点的元素：(其实可以是 func search(location int) interface{})

// FrontQueue 读队头元素
// 队列非空即返回队头元素和 true
// 队列为空则返回 nil 和 false
func (queue *Queue) FrontQueue() (interface{}, bool) {
	// 读之前要上锁
	queue.lock.Lock()
	// 读完解锁
	defer queue.lock.Unlock()
	node, ok := queue.queue.SearchByID(1)
	if ok {
		return node.Data, ok
	}
	return nil, ok

}

// IsFull 判断队满操作
func (queue *Queue) IsFull() bool {
	if queue.cap == queue.queue.Length() {
		return true
	}
	return false
}
