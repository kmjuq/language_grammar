package main

import "fmt"

func 函数语法() {
	fmt.Println("-- 函数语法 --")
	fmt.Println(add2(3, 4))
	fmt.Println(add3(3, 4, 5))
	fmt.Println(vals())
	fmt.Println(vals1())
	// 可变参数
	sum1(1, 2, 3, 4, 5)
	// 指针参数的使用
	zval := 10
	fmt.Println("zval", zval)
	zeroval(zval)
	fmt.Println("zeroval:", zval)
	zeroptr(&zval)
	fmt.Println("zeroptr:", zval)
	fmt.Println("pointer:", &zval)
	// 闭包的使用
	fmt.Println(intSeq()())
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}

// 可变参数
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

// 参数为指针
func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

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
func vals1() (x int, y int) {
	x = 3
	y = 7
	return
}
