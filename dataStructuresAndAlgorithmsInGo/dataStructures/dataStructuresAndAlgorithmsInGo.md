# dataStructuresAndAlgorithmsInGo(数据结构与算法Go语言实现)
Golang 语言中，数据类型分为四大类：
- 1 基础类型：数字、字符串、布尔型
- 2 聚合类型：数组、结构体
        通过组合基础类型得到的复杂的数据类型。
- 3 引用类型：指针、slice、map、函数、通道
        全部都是间接的指向程序变量或者状态
- 4 接口类型

这些数据类型在声明但是没有赋值的时候均为**零值**。
## 1、数组、集合

### 1.1 数组

数组是具有**固定长度**且拥有零个或多个**相同数据类型元素**的序列。数组为**值类型**，默认情况下使用的是**值传递**。

#### 1.1.1 静态数组
- 数组的定义
var 数组名 [数组大小]数据类型
例如：```var a [5]int```
默认情况下，一个数组的元素初始值为元素类型的零值。

- 数组的使用
四种初始化方法：
    ```go
    var numArr01 [3]int = [3]int{1, 2, 3}

    var numArr02 = [3]int{5, 6, 7}
    // [...]是固定写法，"..." 出现在数组长度位置，那么数组的长度由初始化数组的元素个数决定    
    var numArr03 = [...]int{8, 9, 10}

    // 使用“：”初始化的时候为数组的固定位置赋值
    var numArr04 = [...]int{1:5, 0:6, 2:7}

    numArr05 := [...]int{1:5, 0:6, 2:7}

    ```
    两类遍历方法：
    ```go
    // 常规遍历
    for i := 0; i < len(), i++ {
        ...
    }
    // for-range 结构遍历
    for index, value := range array01 {
        ...
    }
    ```

#### 1.1.2 动态数组slice
slice 表示一个拥有相同类型元素的可变长度序列。通常写成 ```[]T``` ，其中元素类型都为 T 。切片属于**引用类型**，使用方法和数组类似。
- slice 的定义
方式一：定义一个切片，然后让切片去引用一个已经创建好的数组。
    ```go
    var arr [5]int = [...]int{1, 2, 3, 4, 5}
    var slice = arr[1:3] // slice = {2, 3, 4}
    ```
    方式二：通过 make 来创建切片。
    基本语法：var 切片名 []type = make([]type, len, cap)
    其中：type 为数据类型，len为大小, cap 为切片容量（可选，如果选了则 cap >= len）。
    ```go
    var slice []float64 = make([]float6, 5, 10)
    slice[1] = 10
    slice[2] = 20
    ```
    通过 make 方式创建切片可以指定切片的大小和容量；如果没有给切片的各个元素赋初值，就会使用默认值；通过 make 创建的切片对应的底层数组由 make 底层维护，对外不可见，只能通过slice访问。
    方式三：定义一个切片就直接指定其具体数组，使用原理类似 make 的方式。
    ```go
    var strSlice []string = []string{"tom", "jack", "mary"}
    ```
- slice 的使用方式
切片的遍历和数组一样，也有两种方式。

### 1.2 map集合
map 是散列表的引用，是一个拥有键值对元素（key--value）的无序集合。在这个集合中，键的值是唯一的，键对应的值可以通过键来获取、更新或移除。**引用类型**
键的类型：必须是可以通过 == 操作符来进行比较的数据类型，一般是int 或者 string。slice、map 和 function 不可以。
值的类型：没有限制，常用的是数字、字符串、map、结构体等。
#### 1.2.1 静态map
- map 的声明与初始化
    var 变量名 map[keyType]valueType
    ```go
    var a map[int]int
    ```
    **注意**：声明不会分配内存的，初始化需要使用 make，分配内存后才可以赋值和使用。
    ```go
    // 声明
    var a map[string]string
    // 初始化
    a = make(map[string]string, 10)
    // 使用
    a["name1"] = "tom"
    a["name2"] = "jack"
    ```
    
    方式一：类似上面代码
    方式二：类型推导
    ```go
    cities := make(map[string]string)
    cities["name1"] = "北京"
    cities["name2"] = "上海"
    cities["name3"] = "天津"
    ```
    方式三：
    ```go
    cities := map[string]string{
        "name1" : "北京"
        "name2" : "上海"
        "name3" : "天津"
    }
    ```
- map 的使用    
使用 for--range 遍历。

#### 1.2.2 动态map (slice实现)
切片的数据类型如果是 map ,map 的个数就可以动态变化了。（map 切片）
```go
// 声明一个切片，这个切片是 map 类型的
var cities []map[string]string
cities = make([]map[string]string, 2)
// 第一个切片因为是 map 类型的，所以要先 make 一下，再增加信息
cities[0] = make(map[string]string, 2)
cities[0]["安徽"] = "合肥"
cities[0]["江苏"] = "南京"
cities[1] = make(map[string]string, 2)
cities[1]["美国"] = "芝加哥"
cities[1]["英国"] = "伦敦"
```

## 2、栈、队列和链表

此节主要是线性表的一些数据结构的具体实现。

### 2.1 栈

#### 2.1.1 栈的定义
栈（stack）又名堆栈，它是一种运算受限的 **线性表**。限定仅在表尾进行插入和删除操作的线性表。这一端被称为栈顶，相对地，把另一端称为栈底。

**线性表包括：**顺序表（数组、栈、队列）、链表

![](img\8b82b9014a90f603eab7c55f3912b31bb051eda7.jpg)

#### 2.1.2 栈的相关方法
- 1) push 方法
①若TOP≥n时，则给出溢出信息，作出错处理（进栈前首先检查栈是否已满，满则溢出；不满则作②）；
②置TOP=TOP+1（栈指针加1，指向进栈地址）；
③S(TOP)=X，结束（X为新进栈的元素）；
- 2) pop 方法
①若TOP≤0，则给出下溢信息，作出错处理(退栈前先检查是否已为空栈， 空则下溢；不空则作②)；
②X=S(TOP)，（退栈后的元素赋给X）：
③TOP=TOP-1，结束（栈指针减1，指向栈顶）。
- 3) top 方法
返回栈顶的元素。
- 4) len 方法
返回栈中的当前元素的个数。
- 5) is_empty 方法
判断栈是否为空，为空则返回true，否则返回false。
#### 2.1.3 使用 Golang 实现栈
- **1 思路**
根据栈的性质以及 go 中相关数据类型的知识，选择 **切片（slice）** 作为栈的主体存储结构。考虑到栈的容量问题：这里想到设计一个可以同时满足：固定容量或者可变容量的栈,所以考虑在结构体内加入 **cap** 字段。
    ```go
    // 这是一个栈的结构体
    type stack struct {

        // 一个栈的主体 ：可以接受所有类型的 interface{} 切片
        items []interface{}

        // 一个读写互斥锁：用于保护栈的数S据，防止读写的冲突操作
        lock sync.RWMutex

        // 一个栈容量：uint 类型，如果为 0 则表示容量自动增长无上限
        cap uint

        // 一个栈高度指针
        height uint

    }
    ```
- **2 注意事项**
**2.1** 因为对栈的主体存储结构切片的读写可能存在线程干扰，所以结构体内引入一个 **读写互斥锁** 用于保护栈的数S据，防止读写的冲突操作。
**2.2** 选用可以接受所有类型的 interface{} 切片，主要是考虑到在使用的时候可以满足不同数据类型入栈的需求。当然如果要在栈内存放单一数据类型，在使用的时候注意一下就行。
**2.3** 注意包下各个方法的返回值的类型，使用时可能需要使用类型断言。
- **3 完整源码**
包 datastructure/stack 下的代码stack.go
  
    ```go
  package stack
  
    import (
        "sync"
  )
  
    // 这是一个栈的结构体
  type stack struct {
  
        // 一个栈的主体 ：可以接受所有类型的 interface{} 切片
      items []interface{}
  
        // 一个读写互斥锁：用于保护栈的数S据，防止读写的冲突操作
      lock sync.RWMutex
  
        // 一个栈容量：uint 类型，如果为 0 则表示容量自动增长无上限
      cap uint
  
        // 一个栈高度指针
      height uint
  
    }

    // 创建栈s 
    // 将此方法绑定到 Stack 这个公开的栈的 struct 上面
    // 传入一个参数 cap (无符号整型) ，表示栈的容量：
    //      传入 0 表示需要创建的栈的容量无上限、自增长
    //      传入 uint 型的非 0 整数表示创建的栈容量有限
    func NewStack(cap uint) *stack {
        var s stack
        if cap != 0 {
        // 有容量限制
            s.items = make([]interface{}, cap)
            s.cap = cap
            return &s 
            
        } else {
            //无容量限制先默认栈的容量为 2 ，后面根据需要自动增长
            s.items =make([]interface{}, 2)
            s.cap = 0
            return &s 
        }
    }

    // 判断栈是否为空，为空则返回true，否则返回false。
    func (s *stack)IsEmpty() bool {
        if s.height == 0 {
            return true
        } else {
            return false
        }
    }

    // 返回顶元素
    func (s *stack)Top() interface{} {
        // 读之前要上锁
        s.lock.Lock()
        // 读完解锁
        defer s.lock.Unlock()

        if s.height > 0 {

            return s.items[s.height - 1]
        } else {

            return nil
        }
    }

    // 进栈操作
    func (s *stack)Push(item interface{}) bool {
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

    // 出栈操作
    func (s *stack)Pop() interface{} {
        // 写之前要上锁
        s.lock.Lock()
        // 写完解锁
        defer s.lock.Unlock()
        if s.IsEmpty() {
            //栈为空，不可以执行出栈操作
            return nil
        } else {
            // 栈非空
            item := s.items[s.height - 1]
            s.items = s.items[:s.height - 1]
            s.height--
            return item
        }
    }

    ```
    main包下的测试代码：
    ```go
    package main

    import (
        "fmt"
        "datastructure/stack"
    )

    func main()  {
        // 调用 stack 包下 NewStack 方法创建一个栈
        s := stack.NewStack(4)
        fmt.Println(s)
        _ = s.Push(1)
        fmt.Println("栈顶元素是：" , s.Top())
        _ = s.Push(2)
        fmt.Println("栈顶元素是：" , s.Top())
        _ = s.Push(3)
        fmt.Println(s.IsEmpty())
        fmt.Println("栈顶元素是：" , s.Top())
        result := s.Pop()
        fmt.Println("弹出的元素是：" , result)
        fmt.Println("此时栈顶元素是：" , s.Top())
        _ = s.Pop()
        _ = s.Pop()
        result = s.Pop()
        if result == nil {
            fmt.Println("此时的栈已空，弹不出来元素了")
        } else {
            fmt.Println("弹出的元素是", result)
        }

    }
    ```

**注意**：以上代码在可导出的函数NewStack中返回的是不可导出的局部变量，虽然使用起来不会有太大的影响，但是不符合 go 程序的设计规范，需要进行修改，修改后的程序见包中的实际代码

### 2.2 队列

#### 2.2.1 队列的定义
队列是一种特殊的 **线性表** ，特殊之处在于它只允许在表的前端（front）进行删除操作，而在表的后端（rear）进行插入操作，和栈一样，队列是一种操作受限制的线性表。进行插入操作的端称为队尾，进行删除操作的端称为队头。队列中没有元素时，称为空队列。
队列的数据元素又称为队列元素。在队列中插入一个队列元素称为入队，从队列中删除一个队列元素称为出队。因为队列只允许在一端插入，在另一端删除，所以只有最早进入队列的元素才能最先从队列中删除，故队列又称为先进先出（FIFO—first in first out）线性表
- **1 顺序队列**
建立顺序队列结构必须为其静态分配或动态申请一片连续的存储空间，并设置两个指针进行管理。一个是队头指针front，它指向队头元素；另一个是队尾指针rear，它指向下一个入队元素的存储位置，如图所示
每次在队尾插入一个元素是，rear增1；每次在队头删除一个元素时，front增1。随着插入和删除操作的进行，队列元素的个数不断变化，队列所占的存储空间也在为队列结构所分配的连续空间中移动。当front=rear时，队列中没有任何元素，称为空队列。当rear增加到指向分配的连续空间之外时，队列无法再插入新元素，但这时往往还有大量可用空间未被占用，这些空间是已经出队的队列元素曾经占用过得存储单元。
顺序队列中的溢出现象：
（1） "下溢"现象：当队列为空时，做出队运算产生的溢出现象。“下溢”是正常现象，常用作程序控制转移的条件。
（2）"真上溢"现象：当队列满时，做进栈运算产生空间溢出的现象。“真上溢”是一种出错状态，应设法避免。
（3）"假上溢"现象：由于入队和出队操作中，头尾指针只增加不减小，致使被删元素的空间永远无法重新利用。当队列中实际的元素个数远远小于向量空间的规模时，也可能由于尾指针已超越向量空间的上界而不能做入队操作。该现象称为"假上溢"现象。
- **2 循环队列**
在实际使用队列时，为了使队列空间能重复使用，往往对队列的使用方法稍加改进：无论插入或删除，一旦rear指针增1或front指针增1 时超出了所分配的队列空间，就让它指向这片连续空间的起始位置。自己真从MaxSize-1增1变到0，可用取余运算rear%MaxSize和front%MaxSize来实现。这实际上是把队列空间想象成一个环形空间，环形空间中的存储单元循环使用，用这种方法管理的队列也就称为循环队列。 **除了一些简单应用之外，真正实用的队列是循环队列。** 
在循环队列中，当队列为空时，有front=rear，而当所有队列空间全占满时，也有front=rear。为了区别这两种情况，规定循环队列最多只能有MaxSize-1个队列元素，当循环队列中只剩下一个空存储单元时，队列就已经满了。因此，队列判空的条件时front=rear，而队列判满的条件时front=（rear+1）%MaxSize。队空和队满的情况如图：

#### 2.2.2 队列的相关方法

1. 出队操作

   OutQueue：初始条件: 队q 存在且非空，操作结果： 删除队首元素，并返回其值，队发生变化；

2. 入队操作

   InQueue：初始条件:队q 存在。操作结果： 对已存在的队列q，插入一个元素x 到队尾，队发生变化；

3. 判满

   IsFull：判断队列是否已满，满返回true，否则返回false；

4. 判空

   IsEmpty：判断队列是否为空，空则返回true，否则返回false；

5. 取队头元素

   FrontQueue：返回队列的队头元素且不出队。

#### 2.2.3 使用切片实现循环队列
- **1 思路**

  根据队列的性质，考虑使用切片来实现栈的主体。由存在 front 和 rear 指针以及一个同步互斥锁，所以考虑如下结构体：

  ```go
  // 队列数据结构
  type circularQueue struct {
  
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
  ```

  

- **2 注意事项**

  循环队列即使在满的状态下也有一个位置是空着不存放元素的。

- **3 完整源码**

  queue包下的 queue.go :

  ```go
  package queue
  
  // 队列的数组实现，(简单顺序队列的实用价值不高) 只实现循环队列
  import (
  	"sync"
  )
  
  // 队列数据结构
  type circularQueue struct {
  
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
  
  //	队列初始化
  //	传入队列的长度
  func NewQueue(cap int) *circularQueue {
  	var cqueue circularQueue
  
  	if cap > 2 {
  		cqueue.queue = make([]interface{}, cap)
  
  	}
  	cqueue.cap = cap
  	// 传入的 cap <= 2 ，理论上这个队列是不可能存在的，所以返回的是nil
  	return &cqueue
  
  }
  
  // 出队操作
  // 队列有值可以返回就返回出队的值和true
  // 队列不满足出队条件就返回 nil 和 false
  func (cqueue *circularQueue) OutQueue() (interface{}, bool) {
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
  
  // 入队操作 TODO
  func (cqueue *circularQueue) InQueue(item interface{}) bool {
  
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
  
  // 判队空操作
  func (cqueue *circularQueue) IsEmpty() bool {
  	// 队首、尾指针一样的时候说明队列为空
  	if cqueue.front == cqueue.rear {
  		return true
  	}
  	return false
  }
  
  // 读队头元素
  // 队列非空即返回队头元素和 true
  // 队列为空则返回 nil 和 false
  func (cqueue *circularQueue) FrontQueue() (interface{}, bool) {
  	// 读之前要上锁
  	cqueue.lock.Lock()
  	// 读完解锁
  	defer cqueue.lock.Unlock()
  	if !cqueue.IsEmpty() {
  		return cqueue.queue[cqueue.front], true
  	}
  	return nil, false
  
  }
  
  // 判断队满操作
  func (cqueue *circularQueue) IsFull() bool {
  	// front=（rear+1）%MaxSize
  	if cqueue.front == (cqueue.rear+1)%cqueue.cap {
  		return true
  	}
  	return false
  }
  
  ```

  main 包下的测试代码：

  ```go
  	// 调用 queue 包下的 NewQueeu 方法创建一个循环队列
  	q := queue.NewQueue(5)
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
  ```

  **注意**：以上代码在可导出的函数NewQueue中返回的是不可导出的局部变量，虽然使用起来不会有太大的影响，但是不符合 go 程序的设计规范，需要进行修改，修改后的程序见包中的实际代码。

### 2.3 链表

#### 2.3.1 链表的定义

**链表**是一种物理存储单元上非连续、非顺序的存储结构，数据元素的逻辑顺序是通过链表中的指针链接次序实现的。链表由一系列**结点**（链表中每一个元素称为结点）组成，结点可以在运行时动态生成。**每个结点包括两个部分：一个是存储数据元素的数据域，另一个是存储下一个结点地址的指针域。** 相比于线性表顺序结构，操作复杂。由于不必须按顺序存储，链表在插入的时候可以达到O(1)的复杂度，比另一种线性表顺序表快得多，但是查找一个节点或者访问特定编号的节点则需要O(n)的时间，而线性表和顺序表相应的时间复杂度分别是O(logn)和O(1)。
使用链表结构可以克服数组链表需要预先知道数据大小的缺点，链表结构可以充分利用计算机内存空间，实现灵活的内存动态管理。但是链表失去了数组随机读取的优点，同时链表由于增加了结点的指针域，空间开销比较大。链表最明显的好处就是，常规数组排列关联项目的方式可能不同于这些数据项目在记忆体或磁盘上顺序，数据的存取往往要在不同的排列顺序中转换。链表允许插入和移除表上任意位置上的节点，但是不允许随机存取。链表有很多种不同的类型：单向链表，双向链表以及循环链表。

**单向链表：**
![](img\单链表演示.png)

**双向链表：**

![](img\双链表演示.png)

#### 2.3.2 链表的主要方法

- 插入函数

  在链表的指定位置插入结点。

- 删除函数

  删除指定位置的链表结点。

- 查找函数

  查找指定结点在链表中的位置。

- 

#### 2.3.3 使用 Golang 链表



### 2.4 栈与队列的链表实现

#### 2.4.1 使用链表实现栈

- **1 思路**

  

- **2 注意事项**

  

- **3 完整源码**

#### 2.4.2 使用链表实现队列

- **1 思路**

  

- **2 注意事项**

  

- **3 完整源码**

## 4、树

### 4.1 二叉树

### 4.2 树

## 5、优先级队列

## 6、映射、哈希表和跳跃表

## 7、搜索树

## 8、排序

## 9、查找

## 10、图

## 11、字符串

## 12、文本处理

## 13、内存管理和 B 树