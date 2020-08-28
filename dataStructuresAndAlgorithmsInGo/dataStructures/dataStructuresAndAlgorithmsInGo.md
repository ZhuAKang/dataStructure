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

  <img src="img\单链表的插入.jpg" style="zoom:80%;" />

- 删除函数

  删除指定位置的链表结点。

  ![](img\单链表的删除.jpg)

- 查找函数

  查找指定结点在链表中的位置。由于链表数据存取的不连续性，不能随机存取。所以在查找的时候只能顺序查找，时间复杂度在O(n)。

#### 2.3.3 使用 Golang 链表

​	**注：**在这里，我尝试去建立一个通用的单链表结构体并实现相关的查询、删除方法。但是在编写查找函数的时候遇见了很大的问题：想要这个结构体通用，则其接受的类型就设为了 interface ，但是两个**原始数据相等但是不相同**的 interface 类型如何比较是否是相等的？尝试使用reflect或者其他的一些办法没有解决比较值的这个问题，不能比较是否相等那就也不能在这个单链表上删除特定元素值的结点了。所以这一块卡住了，但是其他的函数：例如初始化、查询长度、插入函数。查找、删除这两个函数这一块的问题待以后解决吧。或者有大佬可以来给我指点一下。

​	代码在 *linklist/link_list.go*

```go
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

```

main下的测试代码如下：

```go
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

	list.ShowList()
	node, ok := list.SerachInLinkList(2)
	fmt.Println("3在结点", node, ok)
```



### 2.4 栈与队列的链表实现

因为上面的链表功能不够完善，部分函数存在问题，不能够满足栈和队列的实现需求，所以在上面链表的基础之上新增两个函数如下：

```go
// DeleteSP 删除从头结点开始的第 XX 个结点
// 也即删除指定序号的结点
// 删除成功返回删除的结点的值以及true；否则返回 nil 和 false
func (linkList *LinkList) DeleteSP(position int) (interface{}, bool) {
	// 要考虑好，删除的是链表上的第一个还是最后一个又或者是第一个和最后一个指向的是同一个结点
	if position > 0 {
		// 删除第一个且链表就一个元素
		if position == 1 && linkList.Length() == 1 {
			deleteNode := linkList.head.Next
			linkList.tail = linkList.head
			linkList.len--
			return deleteNode.Data, true
		} else if position == linkList.len {
			// 删除最后一个
			deleteNode := linkList.tail
			beforeNode := linkList.head
			// 找到最后一个结点的前一个结点
			for i := 1; i < linkList.len; i++ {
				beforeNode = beforeNode.Next
			}
			// 删除
			linkList.tail = beforeNode
			beforeNode.Next = nil
			linkList.len--
			return deleteNode.Data, true
		} else {
			// 删除这之间的结点
			beforeDeleteNode := linkList.head
			// deleteNode := linkList.head.Next
			// 找到待删结点的前一个结点
			for i := 1; i < position; i++ {
				beforeDeleteNode = beforeDeleteNode.Next
			}
			// 删除
			deleteNode := beforeDeleteNode.Next
			beforeDeleteNode.Next = deleteNode.Next
			linkList.len--
			return deleteNode.Data, true
		}

	}
	return nil, false
}

// SearchByID 根据位置查找结点
// 查找成功则返回查找到的结点以及 true ，否则返回 nil 和 false
func (linkList *LinkList) SearchByID(position int) (*Node, bool) {

	if position <= linkList.Length() && position > 0 {
		loNode := linkList.head
		for i := 1; i < position; i++ {
			loNode = loNode.Next
		}
		return loNode.Next, true
	}
	return nil, false
}
```



#### 2.4.1 使用链表实现栈

- **1 思路**

  

- **2 注意事项**

  栈的结构体中栈高度可以不需要了，链表的长度就是栈的高度。

- **3 完整源码**

  **stack/link_list_stack.go**

  ```go
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
  
  ```

  **main.go中的测试代码：**

  ```go
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
  ```
  

#### 2.4.2 使用链表实现队列

- **1 思路**

  

- **2 注意事项**

  原始队列里面会有 front, rear  用于指示队首与队尾的位置，而现在可以不用了，单链表有 head 和 tail 。

- **3 完整源码**

  **queue/link_list_queue.go**

  ```go
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
  
  ```
  
  
  
  **main.go中的测试代码：**

  ```go
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
  }
  ```
  
  ### 2.5 注意：
  
  以上涉及链表中结点删除的操作均未考虑内存释放的问题，默认就是交给 go 自己的垃圾回收机制管理了。如果要自己做内存释放的话可能要用CGO去做哦。

  

## 4、树

​	上一部分讲的顺序表和链表的实现都不够令人满意：不是检索速度快、就是易于插入新的节点，但是不能同时具有这两个优点。但是树却可以同时具有以上两个特性，且在其上的大部分操作的运行时间平均是 O(log N) 。

### 4.1 二叉树

**以下的代码均在 tree 包下。**

#### 4.1.1 二叉树的定义及特性

一棵二叉树（binary tree）是由结点的有限集合组成的，这个集合或者为空或者由一个根节点（root）以及两棵不相交的二叉树组成，这两棵子树分别称为当前根节点的左子树（left subtree）与右子树（right subtree）。而这两棵子树的根节点有分别成为当前根节点的子节点。

<img src="img/二叉树.jpg" style="zoom:50%;" />

##### a) 完全二叉树

从根节点开始，每一层从左向右填充。

<img src="img/完全二叉树.jpg" style="zoom:50%;" />

一棵深度为k的有n个结点的二叉树，对树中的结点按从上至下、从左到右的顺序进行编号，如果编号为i（1≤i≤n）的结点与满二叉树中编号为i的结点在二叉树中的位置相同，则这棵二叉树称为完全二叉树。

一棵高度为 d 的完全二叉树除了 d - 1 层外，其他每一层都是满的。

##### b) 满二叉树

对于满二叉树国内外的定义是不一样的。这里我们按照国内的来吧。

**国内定义：**一个二叉树，如果每一个层的结点数都达到最大值，则这个二叉树就是满二叉树。也就是说，如果一个二叉树的层数为K，且结点总数是(2^k) -1 ，则它就是满二叉树。

![](img/国内满二叉树.png)

**国外定义：**如果一棵二叉树的结点要么是叶子结点，要么它有两个子结点，这样的树就是满二叉树

![](img/国外满二叉树.png)

#### 4.1.2 二叉树的主要实现方法

##### A）指针实现

使用指针实现二叉树，每个结点存储两个字节的的指针和一个本结点的存储值。
<img src="img/树的节点构造图.PNG" style="zoom:40%;" />
下面是一棵树的演示：
![](img/二叉树的指针实现.png)

golang 实现的结构体如下：

```go
// TNode 树的结点结构体
type TNode struct {
	// 左孩子指针
	left *TNode
	// 右孩子指针
	right *TNode
	// 数据域
	element int
}

// Tree 树的结构体
type Tree struct {
	// 根节点
	root *TNode
}
```

**注：**这里为了避免出现上面链表在查找时候的问题，即 interface 间值比较的问题，就简单一点，data 的类型取 int 得了。（如果之后可以解决这个问题，那以后再改）

##### B）数组实现

使用数组存储二叉树有利有弊，一般来说都是在数组中存入有规律易寻找的树类型，例如完全二叉树或满二叉树。

假设在完全二叉树中，逐层而下、从左到右，结点的位置完全由其序号决定。则可以采用数组有效的存储二叉树的数据，把每一个数据存放在其结点对应序号的位置上。

![二叉树的数组表示](img\二叉树的数组表示.png)

如上面的数组所示，如果数组存储的是完全二叉树（按照宽度优先遍历存储），则其对应的树形就应该是：

<img src="img\数组所代表的二叉树.png" alt="数组所代表的二叉树" style="zoom:80%;" />

如果上面数组存储的的二叉树是按照先序、后序、中序遍历的结点顺序存储的，则代表的可能是不同的二叉树树形。这个后面会讲。

#### 4.1.3 二叉树的遍历、查找等相关函数

##### 4.1.3.1 先序遍历

即在遍历树的时候，先输出根节点，在依序输出左右子树。首先访问根结点然后遍历左子树，最后遍历右子树。在遍历左、右子树时，仍然先访问根结点，然后遍历左子树，最后遍历右子树，如果二叉树为空则返回。

以上面的数组对应的二叉树为例，先序遍历的输出结果应该是：124895367

先序遍历的实现可以采用循环法也可以采用递归法：

1. **循环实现：**

   ```go
   // PreOrderByCircle 先序遍历的循环实现
   func (tree *BTree) PreOrderByCircle() {
   	// 声明一个栈，树根入栈，然后开始循环（条件：栈非空）
   	// 弹出栈顶元素，进行操作，然后将右孩子与左孩子依次入栈（如果右或者左存在的话）
   
   	var stack stack.Lstack
   	stack.InitStack(0)
   
   	// 树根入栈
   	stack.Push(tree.root)
   	for !stack.IsEmpty() {
   		// 出栈并进行相应操作
   		topNode, _ := stack.Pop()
   		node := topNode.(*TNode)
   		fmt.Print(node.element)
   		// 右孩子存在的话入栈
   		if node.right != nil {
   			stack.Push(node.right)
   		}
   		// 左孩子存在的话入栈
   		if node.left != nil {
   			stack.Push(node.left)
   		}
   	}
   	return
   }
   ```

   

2. **递归实现**

   ```go
   // PreOrderByRec 先序遍历的递归实现的入口函数
   func (tree *BTree) PreOrderByRec() {
   	preOrderByRecursion(tree.root)
   }
   
   // preOrderByRecursion 先序遍历的递归实现
   func preOrderByRecursion(node *TNode) {
   	if node != nil {
   		// 先触发当前结点操作
   		fmt.Print(node.element)
   		// 再转向操作左右子树
   		preOrderByRecursion(node.left)
   		preOrderByRecursion(node.right)
   	}
   
   }
   ```

   

##### 4.1.3.2 中序遍历

中序遍历首先遍历左子树，然后访问根结点，最后遍历右子树。若二叉树为空则结束返回，否则继续先遍历左子树......

以上面的数组对应的二叉树为例，中序遍历的输出结果应该是：849251637

先序遍历的实现可以采用循环法也可以采用递归法：

1. **循环实现：**

   ```go
   // InOrderByCircle 中序遍历的循环实现
   func (tree *BTree) InOrderByCircle() {
   	var stack stack.Lstack
   	stack.InitStack(0)
   
   	// TODO: 要设置一个访问标志位，看看出栈的那个结点是从左边回来的还是从右边回来的
   	var beforeNode *TNode
   	// 指示访问的当前位置
   	p := tree.root
   
   	stack.Push(tree.root)
   
   	for !stack.IsEmpty() {
   
   		// 左孩子存在就入栈
   		if p.left != nil && p.left != beforeNode {
   			stack.Push(p.left)
   			p = p.left
   		} else {
   			// 没有左孩子，出栈并执行相关操作
   			// 先保存之前出栈的结点
   			beforeNode = p
   			// 栈顶出栈并执行相关操作
   			topNode, _ := stack.Pop()
   			p = topNode.(*TNode)
   			fmt.Print(p.element)
   			if p.right != nil {
   				p = p.right
   				stack.Push(p)
   			}
   
   		}
   
   	}
   }
   
   ```

   

2. **递归实现**

   ```go
   // InOrderByRec 中序遍历的递归实现的入口函数
   func (tree *BTree) InOrderByRec() {
   	inOrderByRecursion(tree.root)
   }
   
   // inOrderByRecursion 中序遍历的递归实现
   func inOrderByRecursion(node *TNode) {
   	if node != nil {
   		// 先转向操作左子树
   		inOrderByRecursion(node.left)
   		// 再触发当前结点操作
   		fmt.Print(node.element)
   		// 再转向操作左子树
   		inOrderByRecursion(node.right)
   	}
   }
   
   ```

   

##### 4.1.3.3 后序遍历

后序遍历首先遍历左子树，然后遍历右子树，最后访问根结点，在遍历左、右子树时，仍然先遍历左子树，然后遍历右子树，最后遍历根结点。

以上面的数组对应的二叉树为例，后序遍历的输出结果应该是：894526731

先序遍历的实现可以采用循环法也可以采用递归法：

1. **循环实现：**

   ```go
   // PostOrderByCircle 后序遍历的循环实现
   func (tree *BTree) PostOrderByCircle() {
   	// 声明个栈，用来存结点
   	var stack stack.Lstack
   	stack.InitStack(0)
   	// 访问标识，看看刚刚出栈的结点是不是栈顶的孩子
   	var beforeNode *TNode
   
   	// 根节点先入栈
   	stack.Push(tree.root)
   	for !stack.IsEmpty() {
   		// 获取栈顶元素，有左或右孩子，则左右孩子入栈，其中右孩子先入栈
   		topNode, _ := stack.Top()
   		node := topNode.(*TNode)
   		// 当前栈顶结点的左孩子存在且刚刚出栈的结点不是此节点的左右孩子，证明左孩子未被访问过
   		if node.left != nil && beforeNode != node.right && beforeNode != node.left {
   			stack.Push(node.left)
   		} else if node.right != nil && beforeNode != node.right {
   			// 栈顶结点的右孩子存在且刚刚出栈的结点不是这个结点的右孩子，就证明当前栈顶结点的右孩子还未被访问过
   			stack.Push(node.right)
   		} else {
   			// 左右孩子均不存在或者左右孩子以及都被访问过了
   			curNode, _ := stack.Pop()
   			node = curNode.(*TNode)
   			beforeNode = node
   			fmt.Print(node.element)
   		}
   	}
   }
   ```

   

2. **递归实现**

   ```go
   // PostOrderByRec 后序遍历的递归实现的入口函数
   func (tree *BTree) PostOrderByRec() {
   	postOrderByRecursion(tree.root)
   }
   
   // postOrderByRecursion 后序遍历的递归实现
   func postOrderByRecursion(node *TNode) {
   	if node != nil {
   		// 先转向操作左右子树
   		postOrderByRecursion(node.left)
   		postOrderByRecursion(node.right)
   		// 再触发当前结点操作
   		fmt.Print(node.element)
   	}
   }
   
   ```

##### 4.1.3.4 层序遍历

层序遍历即按照树的结构，每一层的去遍历节点。以上面的数组对应的二叉树为例，后序遍历的输出结果应该是：123456789

```go
// LayerOrder 层序遍历的循环实现（好像没有递归的）
func (tree *BTree) LayerOrder() {
	// 声明一个队列，这里使用的是之前写的使用链表实现的队列
	var queue queue.Queue
	// 初始化队列长度，因为是链表实现的队列，且队里只存这一层加上少部分上一层的节点，所以不用设太大
	queue.InitQueue(7)

	queue.InQueue(tree.root)
	// 队列非空的时候一直循环
	for !queue.IsEmpty() {
		// 队头元素出队
		out, _ := queue.OutQueue()
		node := out.(*TNode)
		fmt.Print(node.element)
		if node.left != nil {
			queue.InQueue(node.left)
		}
		if node.right != nil {
			queue.InQueue(node.right)
		}
	}

}
```



##### 4.1.3.5 树高

采用后序遍历的思想，树的高度可以将节点依次入栈、出栈，栈的最大高度作为树的高度。

golang 代码如下（**ltree.go**）：

```go
// GetDepth 返回树的深度（高度）
// TODO: 还有递归的方法，待定
func (tree *BTree) GetDepth() int {
	// 使用栈来完成，这一块的栈就可以用我们之前写的了
	var stack stack.Lstack
	// 初始化栈，栈容量自增长
	stack.InitStack(0)
	// 树的深度，初始化为栈的初始高度 0
	maxDepth := stack.Height()

	// 树的根节点入栈，作为栈底元素
	stack.Push(tree.root)

	// 左右孩子访问标志位，指向刚刚出栈的那个节点
	var accessed *TNode

	for {

		current, ok := stack.Top()

		// 栈内仍有元素，栈顶仍然可以取
		if ok {
			// 类型断言，转换成 *TNode 型
			currentNode, ok := current.(*TNode)
			if ok {
				// 类型转换成功，可以下一步判断了
				// 左孩子存在且左右孩子均未被访问过，入栈（左孩子未被访问入栈可以理解，右孩子没被访问主要是后序遍历思想，先左后右最后中间，
				// 当刚刚出栈的节点是当前栈顶的右孩子的时候，说明左孩子已经访问过了，不需要再访问了）
				if currentNode.left != nil && accessed != currentNode.left && accessed != currentNode.right {
					stack.Push(currentNode.left)
					// 判断高度，现在高的话高度就增加
					if maxDepth < stack.Height() {
						maxDepth = stack.Height()
					}
					continue
				} else if currentNode.right != nil && accessed != currentNode.right {
					// 右孩子存在且未被访问，入栈
					stack.Push(currentNode.right)
					// 判断高度，现在高的话高度就增加
					if maxDepth < stack.Height() {
						maxDepth = stack.Height()
					}
					continue
				} else {
					// 左右都不存在，弹出当前栈顶并重来循环
					_, _ = stack.Pop()
					// 标记当前出栈的节点
					accessed = currentNode
					continue
				}
			}
		}
		// 循环退出条件：栈内空了或者上面栈的 Top 没有元素了，又或者类型转换失败了（栈内存的不是*TNode 类型）
		if stack.IsEmpty() || !ok {
			break
		}
	}
	// 左右结点都访问过了就弹出
	return maxDepth
}

```

tree_test.go 中的测试代码：

```go
func TestGetDepth(t *testing.T) {
	// 这是树的入口
	var btree BTree
	/*
		 		1
			  /   \
			 2     3
			/ \   / \
		   4   5 6   7
			  /
			 8
	*/
	var node1, node2, node3, node4, node5, node6, node7, node8 TNode
	node1.SetElement(1)
	node2.SetElement(2)
	node3.SetElement(3)
	node4.SetElement(4)
	node5.SetElement(5)
	node6.SetElement(6)
	node7.SetElement(7)
	node8.SetElement(8)
	node1.SetLeft(&node2)
	node1.SetRight(&node3)
	node2.SetLeft(&node4)
	node2.SetRight(&node5)
	node3.SetLeft(&node6)
	node3.SetRight(&node7)
	node5.SetLeft(&node8)

	btree.InitTree(&node1)

	fmt.Println("树的高度是", btree.GetDepth())
}
```

同样也可以采用递归来做，代码如下：

```go
// GetDepthRec 二叉树的高度的递归实现的入口函数
func (tree *BTree) GetDepthRec() int {
	return GetDepthByRecursion(tree.root)
}

// GetDepthByRecursion 二叉树的高度的递归实现
func GetDepthByRecursion(node *TNode) int {
	if node != nil {
		left := GetDepthByRecursion(node.left)
		right := GetDepthByRecursion(node.right)
		if left < right {
			return right + 1
		}
		return left + 1
	}
	return 0
}
```



##### 4.1.3.6 树宽

树宽需要使用到层序遍历的思想，每层元素数的最大值就是树宽，以上面的数结构为例，树宽为：4。

```go
// GetWidth 二叉树的宽度
func (tree *BTree) GetWidth() int {
	if tree.root == nil {
		return 0
	}
	// 声明一个队列，这里使用的是之前写的使用链表实现的队列
	var queue queue.Queue
	// 初始化队列长度，因为是链表实现的队列，且队里只存这一层加上少部分上一层的节点，所以不用设太大
	queue.InitQueue(12)

	queue.InQueue(tree.root)
	// count 每层节点数，width 宽度
	count := 1
	width := 1
	// 队列非空
	for !queue.IsEmpty() {
		// 临时保存下一层的节点数
		var size = 0
		for i := 0; i < count; i++ {
			// 出栈
			curNode, _ := queue.OutQueue()
			node := curNode.(*TNode)
			if node.left != nil {
				queue.InQueue(node.left)
				size++
			}
			if node.right != nil {
				queue.InQueue(node.right)
				size++
			}
		}
		// 下一层没有节点了，退出
		if size == 0 {
			break
		}
		if size > width {
			width = size
		}
		// 重新开始下一层了
		count = size
	}
	return width
}

```



**注：以上所有代码均在 ltree.go 文件内，测试代码在 tree_test.go**，**在进行所有二叉树的操作之前，最好首先对是否是空树进行判断，否则就可能在空指针上面兜圈子，很烦！！！**我这上面的文档在做的时候没有考虑到这一块，不过已经在代码中更新了，可以看 ltree.go 文件。

#### 4.1.4 不同类型二叉树的判别

##### A）完全二叉树

**算法思路：**

判断一棵树是否是完全二叉树的思路 ：

​		1>如果树为空，则直接返回错
　　2>如果树不为空：层序遍历二叉树
　　2.1>如果一个结点左右孩子都不为空，则pop该节点，将其左右孩子入队列；
　　2.1>如果遇到一个结点，左孩子为空，右孩子不为空，则该树一定不是完全二叉树；
　　2.2>如果遇到一个结点，左孩子不为空，右孩子为空；或者左右孩子都为空；则该节点之后的队列中的结点都为叶子节点；该树才是完全二叉树，否则就不是完全二叉树；

**代码实现：**

```go
// IsCompleteBTree 判断当前二叉树是否为完全二叉树(使用层序遍历思想)
func (tree *BTree) IsCompleteBTree() bool {
	// 树为空则直接退出
	if tree.root == nil {
		return false
	}
	// 声明队列并进行初始化
	var queue queue.Queue
	queue.InitQueue(12)

	// 根节点入队
	queue.InQueue(tree.root)

	// 开始循环判断(只要队列非空，就一直循环下去直到队列为空或者触发退出机制)
	for !queue.IsEmpty() {
		curNode, _ := queue.OutQueue()
		node := curNode.(*TNode)
		// 左右孩子节点都存在
		if node.left != nil && node.right != nil {
			queue.InQueue(node.left)
			queue.InQueue(node.right)
		} else if node.left == nil && node.right != nil {
			// 左孩子为空右孩子不为空，则不是完全二叉树
			return false
		} else {
			// 左右孩子都为空或者左孩子为空右孩子不为空，
			// 则队列中这个节点之后的所有节点都是叶子节点，才能使得这棵树是完全二叉树
			for !queue.IsEmpty() {
				curNode, _ = queue.OutQueue()
				node := curNode.(*TNode)
				if node.left != nil || node.right != nil {
					return false
				}
			}
		}
	}
	return true
}
```



##### B）满二叉树

判断是不是满二叉树（国内的定义）只要看 ：2^(树高) - 1 == 节点个数  是否成立，成立就是满二叉树，否则就不是。这是满二叉树判定的充要条件。

**代码实现：**

```go
// IsFullBTree 判断当前二叉树是否为满二叉树(国内定义的满二叉树)
// 得到树高和节点的个数，判断是否满足：2^(树高) - 1 = 节点个数
func (tree *BTree) IsFullBTree() bool {
	if tree.root == nil {
		return true
	}
	// 如果2^(树高) - 1 != 节点个数，则不是满二叉树
	if math.Pow(2, float64(tree.GetDepth()))-1 != float64(tree.GetTreeNodeNumber()) {
		return false
	}
	return true
}

// GetTreeNodeNumber 遍历计算树上的节点总数(使用层序)
func (tree *BTree) GetTreeNodeNumber() int {
	if tree.root == nil {
		return 0
	}
	var count int
	// 声明队列并进行初始化
	var queue queue.Queue
	queue.InitQueue(12)

	// 根节点入队
	queue.InQueue(tree.root)
	count++
	for !queue.IsEmpty() {
		curNode, _ := queue.OutQueue()
		node := curNode.(*TNode)
		if node.left != nil {
			queue.InQueue(node.left)
			count++
		}
		if node.right != nil {
			queue.InQueue(node.right)
			count++
		}
	}
	return count
}
```



#### 4.1.5 二叉树的特殊应用

##### 4.1.5.1 查找树ADT-----二叉查找树

二叉查找树（二叉搜索树）是二叉树的一个十分重要的应用。二叉查找树的性质：对于树上的 每一个节点 X ，他的左子树中所有关键字的值小于 X 的关键字值，而他的右子树中所有关键字的值大于 X 的关键字值。例如：

<img src="img\二叉查找树.png" alt="二叉查找树" style="zoom:80%;" />

一般来说，二叉查找树的平均深度是O(log N)，且二叉查找树的节点元素排列是有规律的，所以一般的查找的复杂度和树的高度相等都是O(log N)。这里为了简单起见我假设节点中的关键字都是整数，这样方便操作。其实后期如果有复杂的业务需求，节点的关键字值可以是其他的，然后自己写一个比较函数比较自己存的关键字值就行了。

二叉查找树上的操作大致分为以下三类：

1. 查找

   因为二叉查找树满足左小右大的规律，所以在上面查找的时候，首先待查找元素值和根节点的元素值进行比较，相等就返回根节点，大于根节点则向右走在右子树上重复同样的操作，小于根节点则向左子树走。。。如果查到了某个节点该向左孩子（或者右孩子）继续查的时候，该节点的左孩子（或者右孩子）为空，不存在，即为查失败。

   ```go
   // Search 二叉搜索树的查找函数
   // 查到了返回该节点指针，没查到返回 nil
   func Search(tree *BTree, value int) *TNode {
   	if tree.root == nil {
   		return nil
   	}
   	// 非递归查找
   	// 指示访问节点
   	p := tree.root
   	for p != nil {
   		if p.element > value {
   			p = p.left
   		} else if p.element < value {
   			p = p.right
   		} else {
   			return p
   		}
   	}
   	// 没查到
   	return nil
   }
   
   ```

   

2. 插入

   插入可以按照查找时候的思路，先查找树上是否已经存在待插入元素，如果已经存在就不插入，如果不存在的话就插入到访问到的最后一个节点上。

   ```go
   // Insert 向一棵二叉搜索树中插入一个值
   func Insert(tree *BTree, value int) bool {
   	if tree.root == nil {
   		return false
   	}
   
   	// 指示访问节点
   	p := tree.root
   	// 指示要插入的那个节点
   	var beforeNode *TNode
   	for p != nil {
   		if p.element > value {
   			beforeNode = p
   			p = p.left
   		} else if p.element < value {
   			beforeNode = p
   			p = p.right
   		} else {
   			// 树上已经有元素了
   			return false
   		}
   	}
   	// 构造节点
   	var node TNode
   	node.SetElement(value)
   	if value > beforeNode.element {
   		// 插在右边
   		beforeNode.SetRight(&node)
   	} else {
   		// 插在左边
   		beforeNode.SetLeft(&node)
   	}
   	return true
   }
   ```

   

3. 删除

   删除可以分为三种情况：

   - 待删除结点时一个叶子节点，可以直接删除。

   - 待删除结点只有一个孩子结点，即左右子树存在且只存在一个的时候，可以在父节点调整指针绕过该节点进行删除。

   - 如果该节点具有两个孩子，一般的策略是使用该节点右子树上的最小的数据来替代该节点的数据，并且递归地删除右子树上刚刚查到的最小数据节点。

     ```go
     // Delete 在二叉搜索树中删除一个元素 （入口函数）
     // 删除成功返回 true ，失败（节点元素不存在或者其他因素）返回 false
     func Delete(tree *BTree, value int) bool {
     	if tree.root == nil {
     		return false
     	}
     	// 查找有咩有这个元素
     	delNode := Search(tree, value)
     	// 没有返回 false
     	if delNode == nil {
     		return false
     	}
     	// 树上存在此节点，删除
     	_ = deleteNode(tree.root, value)
     	return true
     }
     
     // deleteNode这是删除的递归函数
     func deleteNode(node *TNode, value int) *TNode {
     	if node == nil {
     		return nil
     	} else if value < node.element {
     		// 向左，递归
     		node.left = deleteNode(node.left, value)
     	} else if value > node.element {
     		// 向右，递归
     		node.right = deleteNode(node.right, value)
     		// 从这儿往下说明找到了删除的元素
     	} else if node.left != nil && node.right != nil {
     		// 左右孩子都存在
     		// 使用右侧最小的替代
     		temp := FindMin(node.right)
     		node.SetElement(temp.element)
     		node.right = deleteNode(node.right, temp.element)
     	} else {
     		// 就一个或者零个孩子
     		// 一般来说，需要一个指示位指示待删除结点，然后后期释放掉该资源（下面1、2两步）
     		//1、 temp := node
     		if node.left == nil {
     			// 左侧没有孩子
     			// 直接放入右孩子
     			node = node.right
     			//2、 free(temp)
     		} else if node.right == nil {
     			// 右侧没孩子
     			// 直接放左侧的孩子
     			node = node.left
     			//2、 free(temp)
     		}
     	}
     	return node
     }
     
     // FindMin 查找子树最小的元素
     func FindMin(node *TNode) *TNode {
     	if node == nil {
     		return nil
     	}
     	for node.left != nil {
     		node = node.left
     	}
     	return node
     }
     
     ```

     **注意：以上代码在 tree/bstree.go 文件内，测试代码仍在 tree/tree_test.go 文件内，测试函数的名称是 TestBSTree**

##### 4.1.5.2 表达式树

表达式树的树叶是操作数，比如说是常量或者变量，其他节点为操作符。这里由于所有的操作都是二元的，因此这棵树恰好为表达式树。

<img src="img\表达式树案例.png" alt="表达式树案例" style="zoom:50%;" />

如上图，这是表达式<img src="img\公式1.png" alt="公式1" style="zoom:50%;" />的表达式树。我们可以通过递归地计算左子树与右子树在根处的值来得到表达式树的值。

表达式树的中序遍历的结果即为我们平常所写的表达式类型，例如上图的表达式树的中序遍历为：<img src="img\公式1.png" alt="公式1" style="zoom:50%;" />，表达式树的先序遍历结果即为前缀表达式（波兰表示），表达式树的后续遍历被称为后缀表达式（逆波兰表示）。

这里生成与计算表达式树有两个思路：

​	第一个，重新构建表达式树的结点结构，里面存放左指针、右指针、运算符与数值。这样每个叶子结点只有数值域非空其他均为空，然后再计算的时候，使用递归的方法从下往上，递归地计算左子树与右子树在根处的值并保存到数值域来得到表达式树的值。<img src="img\一类表达式树结点结构.png" alt="一类表达式树结点结构" style="zoom:80%;" />

​	第二个，可以在之前树的结点结构的基础之上，让数据域可以存放操作符和数值（比如采用 go 中的interface作为数据域或者在数据域中存放字符、ASCII 码等）。

​	以上这两个方法各有利弊，有些时候一个内存占用大，一个占用小；有些时候表达式的值得来的方便。

这里我们采用第二个方式，结点数据域内存放字符串string类型（数据值与操作符都可以使用 string 类型存储，只需要在循环或者递归的时候判断一下，叶子结点全是数值，非叶子结点全是操作符）。

##### 4.1.5.3 AVL树

##### 4.1.5.4 伸展树

##### 4.1.5.5 B树

### 4.2 树

#### 4.2.1 树的定义与特性

#### 4.2.1 树的实现

#### 4.2.1 K 叉树

#### 4.2.1 树的顺序表示法

## 5、优先级队列

## 6、映射、哈希表和跳跃表

## 7、搜索树

## 8、排序

## 9、查找

## 10、图

## 11、字符串

## 12、文本处理

## 13、内存管理