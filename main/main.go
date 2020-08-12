package main

import (
	"datastructure/queue"
	"datastructure/stack"
	"fmt"
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
	fmt.Println(s.IsEmpty())
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
			fmt.Printf("第 %d 次出队失败", i)
		}
	}
}
