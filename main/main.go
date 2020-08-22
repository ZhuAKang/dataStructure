package main

import (
	"datastructure/linklist"
	"datastructure/queue"
	"datastructure/stack"
	"fmt"
	"reflect"
)

func main() {

	// 声明一个 stack.Stack 的结构体变量
	var s stack.Stack
	// 调用 stack 包下 InitStack 方法初始化一个栈
	s.InitStack(4)

	fmt.Println(s)
	_ = s.Push(1)
	top, _ := s.Top()
	fmt.Println("栈顶元素是：", top)
	_ = s.Push(2)
	top, _ = s.Top()
	fmt.Println("栈顶元素是：", top)
	_ = s.Push(3)
	fmt.Println("此时栈为空吗？", s.IsEmpty())
	top, _ = s.Top()
	fmt.Println("栈顶元素是：", top)
	result, _ := s.Pop()
	fmt.Println("弹出的元素是：", result)
	top, _ = s.Top()
	fmt.Println("此时栈顶元素是：", top)
	_, _ = s.Pop()
	_, _ = s.Pop()
	result, _ = s.Pop()
	if result == nil {
		fmt.Println("此时的栈已空，弹不出来元素了")
	} else {
		fmt.Println("弹出的元素是", result)
	}

	// 声明一个 queue.CircularQueue 的结构体变量
	var q queue.CircularQueue
	// 调用 queue 包下的 InitQueue 方法初始化循环队列
	q.InitQueue(5)

	fmt.Println(q)
	_ = q.InQueue(1)
	_ = q.InQueue(2)
	_ = q.InQueue(3)
	resultQ, _ := q.FrontQueue()
	fmt.Println("此时的队头元素是：", resultQ)
	_ = q.InQueue(4)
	for i := 0; i < 5; i++ {
		resultQ, ok := q.OutQueue()
		if ok {
			fmt.Printf("第 %d 次出队成功，出队元素 %v \n", i+1, resultQ)
		} else {
			fmt.Printf("第 %d 次出队失败\n", i)
		}
	}

	// 声明一个 linklist.LinkList 的结构体变量
	var list linklist.LinkList
	// 初始化链表
	(&list).InitLinkList()
	// 初始化 Node 结点
	var node1 linklist.Node
	node1.Data = 1
	var node2 linklist.Node
	node2.Data = 2
	var node3 linklist.Node
	node3.Data = 3
	var node4 linklist.Node
	node4.Data = 4
	// 放入 list 内
	list.InsertIntoLinkList(&node1, 1)
	list.InsertIntoLinkList(&node3, 1)
	list.InsertIntoLinkList(&node2, 1)
	list.InsertIntoLinkList(&node4, 1)

	// fmt.Println(&list)
	// fmt.Println(unsafe.Sizeof(node1.Next))
	// fmt.Println(unsafe.Sizeof(node1.Data))
	// fmt.Println(unsafe.Sizeof(list))

	// list.ShowList()
	// node, ok := list.SerachByValue(2)
	// fmt.Println("3在结点", node, ok)

	// 声明一个 queue.Queue 的结构体变量
	var lQ queue.Queue
	// 调用 queue 包下的 InitQueue 方法初始化循环队列
	lQ.InitQueue(5)

	fmt.Println(lQ)
	_ = lQ.InQueue(1)
	_ = lQ.InQueue(2)
	_ = lQ.InQueue(3)
	resultQ, _ = lQ.FrontQueue()
	fmt.Println("队列的链表实现：此时的队头元素是：", resultQ)
	_ = lQ.InQueue(4)
	for i := 0; i < 5; i++ {
		resultQ, ok := lQ.OutQueue()
		if ok {
			fmt.Printf("队列的链表实现：第 %d 次出队成功，出队元素 %v \n", i+1, resultQ)
		} else {
			fmt.Printf("队列的链表实现：第 %d 次出队失败\n", i)
		}
	}

	// 声明一个 stack.Lstack 的结构体变量
	var Ls stack.Lstack
	// 调用 stack 包下 InitStack 方法初始化一个栈
	Ls.InitStack(4)

	fmt.Println(Ls)
	_ = Ls.Push(1)
	LsTop, _ := Ls.Top()
	fmt.Println("栈的链表实现：栈顶元素是：", LsTop)
	fmt.Println("栈的链表实现：栈高度是：", Ls.Height())
	_ = Ls.Push(2)
	LsTop, _ = Ls.Top()
	fmt.Println("栈的链表实现：栈顶元素是：", LsTop)
	_ = Ls.Push(3)
	fmt.Println("此时栈为空吗？", Ls.IsEmpty())
	fmt.Println("栈的链表实现：栈高度是：", Ls.Height())
	LsTop, _ = Ls.Top()
	fmt.Println("栈的链表实现：栈顶元素是：", LsTop)
	LsResult, _ := Ls.Pop()
	fmt.Println("栈的链表实现：弹出的元素是：", LsResult)
	LsTop, _ = Ls.Top()
	fmt.Println("栈的链表实现：此时栈顶元素是：", LsTop)
	_, _ = Ls.Pop()
	_, _ = Ls.Pop()
	LsResult, _ = Ls.Pop()
	if LsResult == nil {
		fmt.Println("栈的链表实现：此时的栈已空，弹不出来元素了")
	} else {
		fmt.Println("栈的链表实现：弹出的元素是", LsResult)
	}

	var a interface{}
	var b interface{}
	a = 8
	b = 8
	fmt.Println("两个是否相等", reflect.ValueOf(a) == reflect.ValueOf(b))
	// A := a.(reflect.ValueOf(a).Kind())
	// B := b.(reflect.ValueOf(b).Kind())

}
