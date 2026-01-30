package main

import (
	"embed"
	"fmt"
	"grammar/asm"
	folderUtil "grammar/folder"
	"time"
)

const s string = "constant"

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileByte []byte

//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func 变量声明() {
	fmt.Println("-- 变量声明 --")
	// := 短变量声明 声明 + 赋值二合一; 只能用在函数内的"局部""变量";
	// = 是单纯的赋值操作; var 是变量申明
	// := 在接受函数调用的多返回值时比较好用
	var a = "initial"
	var e, c int = 1, 2
	d := 3
	// 该声明 相当于 var i = 1 或者 var i int = 1
	g := 1
	// 消费变量，go 语言中不允许有未使用的变量
	fmt.Println(a, e, c, d, g)

	// 如果只申明了没有赋值，会给每个类型默认的 “零值”
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// 类型转换
	var h = float64(g)
	fmt.Printf("类型转换 %v", h)
}

func for循环() {
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
}

func ifelse分支() {
	fmt.Println("-- if/else 结构 --")
	// 在判断条件之前, 可以使用一个表达式, 表达式只可以使用 :=
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func switch分支() {
	fmt.Println("-- switch 结构 --")
	// go 语言的switch 的case 默认是自动跳过的，不想跳过可以使用 fallthrough 关键字,不过基本不使用，case可以多条件
	num := 2
	switch num {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3, 4:
		fmt.Println("3 or 4")
		fallthrough
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("早上好！")
	case t.Hour() < 17:
		fmt.Println("下午好！")
	default:
		fmt.Println("晚上好！")
	}
}

func 数组结构() {
	fmt.Println("-- 数组结构 --")
	// 数组声明
	var arrA [5]int
	// 数组的字面量声明
	arrB := [5]int{1, 2, 3, 4, 5}
	arrA[4] = 100
	fmt.Println(arrA, arrB)
	fmt.Println("数组的长度", len(arrA))
}

func 切片结构() {
	fmt.Println("-- 切片结构 --")
	// go 语言的切片就是内存地址的映射，本身不存储数据，可以通过切片来操作具体的内存数据
	// 切片的声明, 切片的零值是 nil
	var sliceA []string
	fmt.Println("uninit:", sliceA, sliceA == nil, len(sliceA) == 0)
	// 给切片申请内存
	sliceA = make([]string, 3)
	fmt.Println("emp:", sliceA, "len:", len(sliceA), "cap:", cap(sliceA))

	// 切片字面量
	primes := []int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
	var s1 = primes[:4]
	var s2 = primes[1:]
	// 给切片添加元素，会直接覆盖切片后的原数组元素
	s = append(s, 99)
	fmt.Println(s)
	var s3 = primes[:]
	fmt.Println(s1, s2, s3)
	// 切片的长度就是它所包含的元素个数。
	// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
	fmt.Println(len(s1), cap(s1))
}

func map结构() {
	fmt.Println("-- map结构 --")
	// 创建一个map数据结构 key 为string value 为 int
	m := make(map[string]int)
	// map 的增删改查
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)
	v1 := m["k1"]
	if v2, ok1 := m["k2"]; ok1 {
		fmt.Println("v2:", v2)
	}
	_, ok2 := m["k3"]
	fmt.Println(v1, ok2)
	delete(m, "k2")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)
	// map 数据结构的字面量声明
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

func range结构() {
	fmt.Println("-- range语法 --")
	// 数组的字面量声明
	arrB := [5]int{1, 2, 3, 4, 5}
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
}

func goroutine() {
	fmt.Println("-- goroutines --")
	var func1 func(from string)
	func1 = func(from string) {
		for i := 0; i < 3; i++ {
			fmt.Println(from, ":", i)
		}
	}
	func1("direct")
	go func1("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}

func defer语法() {
	fmt.Println("-- defer --")
	// defer 语法需要的变量会在 defer语法 申明的时候捕获，比如说在defer改变变量后，也不会在defer语法中生效
	num := 3
	fun1 := func(num int) {
		fmt.Println("fun1", num)
	}
	defer fun1(num)
	num += 1
	fmt.Println("func call")
}

// 用于将静态资源编译到二进制文件中
func embed资源文件嵌入() {
	fmt.Println("-- embed directive --")
	print(fileString)
	print(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))
	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}

func panic语法() {
	fmt.Println("-- panic --")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd. Error:\n", r)
		}
	}()
	panic("a problem")
}

func 方法权限() {
	// 大写开头的方法就是被公开的方法
	folderUtil.PublicFunc()
	// 小写的就是私有方法
	// folderUtil.privateFunc()
}

func 指针类型() {
	fmt.Println("-- 指针类型 --")
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}

func 类型判断() {
	var i any = "hello"

	s := i.(string)
	fmt.Println(s)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	whatAmI := func(i any) {
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
}

func 类型别名() {
	fmt.Println("\n类型定义语法:")

	// 类型定义
	type MyInt int

	// 类型别名
	type IntAlias = int

	var a MyInt = 42
	var b IntAlias = 100
	var c int = 200

	fmt.Printf("MyInt: %d, 类型: %T\n", a, a)
	fmt.Printf("IntAlias: %d, 类型: %T\n", b, b)
	fmt.Printf("int: %d, 类型: %T\n", c, c)

	// 类型别名可以直接赋值
	b = c
	fmt.Printf("IntAlias 赋值后: %d\n", b)

	// 类型定义需要显式转换
	// a = c // 编译错误
	a = MyInt(c)
	fmt.Printf("MyInt 转换后: %d\n", a)
}

// 核心：go generate 触发代码生成（零侵入，规范写法）
//
//go:generate go run ./gen/loggen.go -src=../struct.go
func main() {
	变量声明()
	常量声明()
	for循环()
	ifelse分支()
	switch分支()
	数组结构()
	切片结构()
	map结构()
	range结构()
	函数语法()
	struct结构()
	goroutine()
	channel()
	defer语法()
	embed资源文件嵌入()
	panic语法()
	方法权限()
	指针类型()
	接口结构()
	类型判断()
	类型别名()
	asm.Asm汇编语法()
	cgo语法()
}
