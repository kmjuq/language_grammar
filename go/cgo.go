package main

/*
#include <stdio.h>
#include <stdlib.h>

// 简单的汇编函数
int c_add(int a, int b) {
    asm volatile (
        "addl %%ebx, %%eax;"
        : "=a" (a)
        : "a" (a), "b" (b)
        : "cc"
    );
    return a;
}

void c_print(const char* s) {
    printf("C层打印: %s\n", s);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func cgo语法() {
	fmt.Println("-- cgo syntax --")
	// Go 调用 C 函数：基础类型传参
	a, b := C.int(10), C.int(20)
	cRes := C.c_add(a, b)
	fmt.Printf("Go调用C函数 c_add(10,20) = %d\n", cRes)

	// Go 调用 C 函数：字符串传参
	goStr := "Hello Cgo"
	cStr := C.CString(goStr)
	defer C.free(unsafe.Pointer(cStr)) // 使用 defer 确保内存释放
	C.c_print(cStr)
}
