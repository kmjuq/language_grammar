package main

import (
	"embed"
	"fmt"
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
	var a = "initial"
	var b, c int = 1, 2
	// 该声明 相当于 var i = 1 或者 var i int = 1
	i := 1
	// 消费变量，go 语言中不允许有未使用的变量
	fmt.Println(a, b, c, i)
}

func 常量声明() {
	fmt.Println("-- 常量声明 --")
	const d = 4000
	const f int = 5000
	fmt.Println(d, f, s)
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
	// 在判断条件之前，可以使用一个表达式
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
	// 切片的声明
	var sliceA []string
	fmt.Println("uninit:", sliceA, sliceA == nil, len(sliceA) == 0)
	// 给切片申请内存
	sliceA = make([]string, 3)
	fmt.Println("emp:", sliceA, "len:", len(sliceA), "cap:", cap(sliceA))
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
	defer fmt.Println("defer ")
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
	接口()
	goroutine()
	channel()
	defer语法()
	embed资源文件嵌入()
	panic语法()
}
