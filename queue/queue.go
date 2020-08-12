package queue

// 队列的数组实现，(简单顺序队列的实用价值不高) 只实现循环队列
import (
	"sync"
)

// CircularQueue 这是循环队列的结构体
// 声明结构体变量之后需要执行 InitQueue 方法初始化队列
type CircularQueue struct {

	// 队列主体
	queue []interface{}

	// 队首队尾的指示
	front, rear int

	// 队列的容量
	cap int

	// 一个读写互斥锁：用于保护栈的数S据，防止读写的冲突操作
	lock sync.RWMutex

	// 注意，即使循环队列为满状态，也有一个位置是空着的
}

// InitQueue 初始化队列
// 传入 int 类型的队列的长度
func (cqueue *CircularQueue) InitQueue(cap int) {

	if cap > 2 {
		cqueue.queue = make([]interface{}, cap)

	}
	cqueue.cap = cap
	// 传入的 cap <= 2 ，理论上这个队列是不可能存在的，所以返回的是nil
}

// OutQueue 出队操作
// 队列有值可以返回就返回出队的值和true
// 队列不满足出队条件就返回 nil 和 false
func (cqueue *CircularQueue) OutQueue() (interface{}, bool) {
	// 读之前要上锁
	cqueue.lock.Lock()
	// 读完解锁
	defer cqueue.lock.Unlock()
	// 队列非空
	if !cqueue.IsEmpty() {

		item := cqueue.queue[cqueue.front]
		// 队首指针指向下一个位置
		cqueue.front = (cqueue.front + 1) % cqueue.cap

		return item, true
	}
	return nil, false
}

// InQueue 入队操作
func (cqueue *CircularQueue) InQueue(item interface{}) bool {

	// 读之前要上锁
	cqueue.lock.Lock()
	// 读完解锁
	defer cqueue.lock.Unlock()

	if !cqueue.IsFull() {

		cqueue.queue[cqueue.rear] = item
		// 队尾指针指向下一个位置
		cqueue.rear = (cqueue.rear + 1) % cqueue.cap

		return true
	}
	return false
}

// IsEmpty 判队空操作
func (cqueue *CircularQueue) IsEmpty() bool {
	// 队首、尾指针一样的时候说明队列为空
	if cqueue.front == cqueue.rear {
		return true
	}
	return false
}

// FrontQueue 读队头元素
// 队列非空即返回队头元素和 true
// 队列为空则返回 nil 和 false
func (cqueue *CircularQueue) FrontQueue() (interface{}, bool) {
	// 读之前要上锁
	cqueue.lock.Lock()
	// 读完解锁
	defer cqueue.lock.Unlock()
	if !cqueue.IsEmpty() {
		return cqueue.queue[cqueue.front], true
	}
	return nil, false

}

// IsFull 判断队满操作
func (cqueue *CircularQueue) IsFull() bool {
	// front=（rear+1）%MaxSize
	if cqueue.front == (cqueue.rear+1)%cqueue.cap {
		return true
	}
	return false
}
