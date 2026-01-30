package asm

import "fmt"

// 声明汇编函数，不需要实现
func add(a, b int) int

func Asm汇编语法() {
    fmt.Println("-- 汇编语法 --")
    result := add(10, 20)
    fmt.Printf("10 + 20 = %d\n", result)
}