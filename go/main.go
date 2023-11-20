package main

import "fmt"
import (
    "math"
    "time"
    "embed"
)

const s string = "constant"

//go:embed folder/single_file.txt
var fileString string
//go:embed folder/single_file.txt
var fileByte []byte
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS


func main() {
	fmt.Println("-- 变量声明 --")
	var a = "initial"
	var b, c int = 1, 2
	// 该声明 相当于 var i = 1 或者 var i int = 1
	i := 1
	// 消费变量，go 语言中不允许有未使用的变量
	fmt.Println(a, b, c, i)

	fmt.Println("-- 常量声明 --")
	const d = 4000
	const f int = 5000
	fmt.Println(d, f, s)

	fmt.Println("-- for 循环 --")
	count := 1
	for count <= 3 {
		fmt.Println(count)
		count++
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	fmt.Println("-- if/else 结构 --")
	// 在判断条件之前，可以使用一个表达式
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	fmt.Println("-- switch 结构 --")
	// go 语言的switch 的case 默认是自动跳过的，不想跳过可以使用 XXX 关键字
	num := 2
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("three")
	}
	whatAmI := func(i interface{}) {
		// switch 可以直接使用表达式的
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	fmt.Println("-- 数组结构 --")
	// 数组声明
	var arrA [5]int
	// 数组的字面量声明
	arrB := [5]int{1, 2, 3, 4, 5}
	arrA[4] = 100
	fmt.Println(arrA, arrB)
	fmt.Println("数组的长度", len(arrA))

	fmt.Println("-- 切片结构 --")
	// go 语言的切片就是内存地址的映射，本身不存储数据，可以通过切片来操作具体的内存数据
	// 切片的声明
	var sliceA []string
	fmt.Println("uninit:", sliceA, sliceA == nil, len(sliceA) == 0)
	// 给切片申请内存
	sliceA = make([]string, 3)
	fmt.Println("emp:", sliceA, "len:", len(sliceA), "cap:", cap(sliceA))

	fmt.Println("-- map结构 --")
	// 创建一个map数据结构 key 为string value 为 int
	m := make(map[string]int)
	// map 的增删改查
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)
	v1 := m["k1"]
	v2, ok1 := m["k2"]
	_, ok2 := m["k3"]
	fmt.Println(v1, v2, ok1, ok2)
	delete(m, "k2")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)
	// map 数据结构的字面量声明
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	fmt.Println("-- range语法 --")
	sum := 0
	for _, num := range arrB {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range arrB {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	fmt.Println("-- 函数语法 --")
	fmt.Println(add2(3, 4))
	fmt.Println(add3(3, 4, 5))
    fmt.Println(vals())
    fmt.Println(vals1())
    // 可变参数
	sum1(1, 2, 3, 4, 5)
    // 指针参数的使用
    zval := 10
    fmt.Println("zval",zval)
    zeroval(zval)
    fmt.Println("zeroval:", zval)
    zeroptr(&zval)
    fmt.Println("zeroptr:", zval)
    fmt.Println("pointer:", &zval)
    // 闭包的使用
    nextInt := intSeq()
    fmt.Println(intSeq()())
    fmt.Println(intSeq()())
    fmt.Println(intSeq()())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    fmt.Println("-- struct结构 --")
    p1 := person{name:"Sean",age:50}
    fmt.Println(p1.name)
    // 指针的使用
    p2 := &p1
    fmt.Println(p2.age)
    p2.age = 51
    fmt.Println(p2.age)

    // struct 结构体的字面量表达方式
    dog := struct {
        name string
        isGood bool
    }{
        "Rex",
        true,
    }
    fmt.Println(dog)

    // 方法
    r1 := rect{width: 10, height: 5}
    fmt.Println(r1)
    r2 := r1
    fmt.Println(r2)

    // 结构嵌入
    co := container {
        base: base {
            num: 1,
        },
        str : "some name",
    }
    fmt.Println(co)
    // 可以直接调用内嵌的结构体的方法
    fmt.Println(co.describe())
    // 内嵌结构体的方法可以看作是结构体自己的方法，算直接实现接口
    type describer interface {
        describe() string
    }
    var desc describer = co
    fmt.Println("describer:",desc.describe())
    fmt.Println("-- 接口 --")
    circle1 := circle{radius:5}
    measure(circle1)

    fmt.Println("-- goroutines --")
    func1("direct")
    go func1("goroutine")
    go func(msg string){
        fmt.Println(msg)
    }("going")
    time.Sleep(time.Second)
    fmt.Println("done")

    fmt.Println("-- channels --")
    msgChan := make(chan string)
    go func() {
        msgChan <- "ping"
    }()
    msg := <- msgChan
    fmt.Println(msg)
    // 缓冲 channel
    msgChan2 := make(chan string,2)
    msgChan2 <- "buffered"
    msgChan2 <- "channel"

    fmt.Println(<-msgChan2)
    fmt.Println(<-msgChan2)
    // 异步channel
    done := make(chan bool,1)
    go worker(done)
    <-done

    pings := make(chan string,1)
    pongs := make(chan string,1)
    ping(pings,"passed message")
    pong(pings,pongs)
    fmt.Println(<-pongs)

    // select 语法，主要用来处理多个channel的消费问题
    c1 := make(chan string)
    c2 := make(chan string)
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }

    // channel close
    jobs := make(chan int ,5)
    doneChan := make(chan bool)

    go func(){
        for {
            j,more := <- jobs
            if more {
                fmt.Println("received job",j)
            } else {
                fmt.Println("received all jobs")
                doneChan <- true
                return
            }
        }
    }()
    
    for j:=1;j<=3;j++ {
        jobs <- j
        fmt.Println("sent job",j)
    }
    // close 函数会关闭channel，防止继续加入数据到channel中。
    close(jobs)
    fmt.Println("sent all jobs")
    <- doneChan

    // channel 的 range 用户
    queue:= make(chan string,2)
    queue <- "one"
    queue <- "two"
    close(queue)
    for elem := range queue {
        fmt.Println(elem)
    }
    
    fmt.Println("-- defer --")
    deferfunc()

    fmt.Println("-- embed directive --")
    print(fileString)
    print(string(fileByte))
    content1, _ := folder.ReadFile("folder/file1.hash")
    print(string(content1))
    content2, _ := folder.ReadFile("folder/file2.hash")
    print(string(content2))

    fmt.Println("-- panic --")
    defer func(){
        if r:= recover();r!=nil {
            fmt.Println("Recoverd. Error:\n",r)
        }
    }()
    mayPanic()
    fmt.Println("After mayPanic")
}

// -- 函数语法 --
func add2(a int, b int) int {
	return a + b
}

func add3(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

// 可以直接给返回值设置变量，然后直接返回，也是可以的
func vals1() (x int,y int){
    x = 3
    y = 7
    return 
}

func sum1(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// 递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

// 参数为指针
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}


// -- struct 结构体 --
type person struct {
    name string
    age  int
}

// 函数返回值是struct的指针
func newPerson(name string) *person {
    p := person{name:name}
    p.age = 42
    return &p
}

type rect struct {
    width,height int
}

// 方法 是结构体的方法，当创建结构体时，可以直接通过结构体调用方法
func (r *rect) area() int {
    return r.width * r.height
}

// 结构嵌入
type base struct {
    num int
}

func (b base) describe() string {
    return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
    base
    str string
}


// -- 接口 --
type geometry interface {
    area() float64
    perim() float64
}

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}


// -- 泛型 --
func MapKeys[K comparable, V any](m map[K]V) []K {
    r := make([]K, 0, len(m))
    for k := range m {
        r = append(r, k)
    }
    return r
}

type List[T any] struct {
    head, tail *element[T]
}

type element[T any] struct {
    next *element[T]
    val  T
}

func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}

func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

// -- goroutines --
func func1(from string) {
    for i:=0;i<3;i++ {
        fmt.Println(from,":",i)
    }
}

// -- channels --
func worker(done chan bool){
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    done <- true
}

// 通过chan关键左右的符号，可以限制参数只能接收或者发送消息
func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

// -- defer --
func deferfunc(){
    defer fmt.Println("defer ")
    fmt.Println("func call")
}

// -- panic --
func mayPanic(){
    panic("a problem")
}

