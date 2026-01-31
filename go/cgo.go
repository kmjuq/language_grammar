package main

/*
#include <stdio.h>
#include <stdlib.h>

// 多架构条件编译：根据CPU架构选择对应的汇编实现
// macOS Intel(x86-64) 预定义宏：__x86_64__
// macOS M4(ARM64) 预定义宏：__aarch64__
#ifdef __x86_64__
// x86-64 架构（Intel）：AT&T 格式汇编，复用你原有代码（仅补充注释）
int c_add(int a, int b) {
    asm volatile (
        "addl %%ebx, %%eax;"  // 32位加法：eax = eax + ebx（a+b）
        : "=a" (a)            // 输出约束：结果写入a，复用eax寄存器
        : "a" (a), "b" (b)    // 输入约束：a存入eax，b存入ebx
        : "cc"                // 被修改的状态：cc表示条件码寄存器被修改
    );
    return a;
}
#elif defined(__aarch64__)
// ARM64 架构（M4/M1-M3）：GNU ARM64 格式汇编，适配Apple Silicon
int c_add(int a, int b) {
    asm volatile (
        "add %w0, %w1, %w2;" // 32位加法：%w0 = %w1 + %w2（a+b）
                             // %w 表示使用32位子寄存器（匹配C的int类型，ARM64寄存器默认64位）
        : "=r" (a)           // 输出约束：结果写入a，编译器自动分配通用寄存器
        : "r" (a), "r" (b)   // 输入约束：a、b存入任意通用寄存器，编译器自动管理
        // 无需声明被修改寄存器：未使用固定寄存器，编译器自动处理
    );
    return a;
}
#else
// 兜底：非目标架构编译直接报错，避免运行时异常
#error "unsupported architecture: only x86-64(Intel) and aarch64(M4) are supported on macOS"
#endif

// 原有的C打印函数，无需修改（纯C代码天然跨架构）
void c_print(const char* s) {
    printf("C层打印: %s\n", s);
}
*/
import "C" // 注意：import "C" 必须紧跟CGO注释块，中间无空行（CGO语法强制要求）
import (
	"fmt"
	"unsafe"
)

func cgo语法() {
	fmt.Println("-- cgo syntax --")
	// Go 调用 C 函数：基础类型传参（无需修改）
	a, b := C.int(10), C.int(20)
	cRes := C.c_add(a, b)
	fmt.Printf("Go调用C函数 c_add(10,20) = %d\n", cRes)

	// Go 调用 C 函数：字符串传参（无需修改，注意内存释放）
	goStr := "Hello Cgo"
	cStr := C.CString(goStr)
	defer C.free(unsafe.Pointer(cStr)) // 使用 defer 确保C内存释放，避免内存泄漏
	C.c_print(cStr)
}
