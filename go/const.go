package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func 常量声明() {
	fmt.Println("-- 常量声明 --")
	const d = 4000
	const f int = 5000
	fmt.Println(d, f, s)
}
