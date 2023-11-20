#include <stdio.h>
#include <stdlib.h>
// extern 说明符
// 在当前文件中声明，这个变量是其他文件定义的，告诉编译器，不需要在这里为它分配内存空间。
extern int a[];
// 宏
#define PI 3.14
#define SQUARE(X) X*X
#define SQUARE1(X) ((X)*(X))
#define PRINT_NUMS_TO_PRODUCT(a, b) { \
  int product = (a) * (b); \
  for (int i = 0; i < product; i++) { \
    printf("%d\n", i); \
  } \
}
#define QUADP(a, b, c) ((-(b) + sqrt((b) * (b) - 4 * (a) * (c))) / (2 * (a)))
#define QUADM(a, b, c) ((-(b) - sqrt((b) * (b) - 4 * (a) * (c))) / (2 * (a)))
#define QUAD(a, b, c) QUADP(a, b, c), QUADM(a, b, c)
// # 运算符 和 ## 运算符
#define STR(X) #X
#define XNAME(n) "x"#n
#define MK_ID(n) i##n
#define X(...) #__VA_ARGS__
// #undef
#define LIMIT 400
#undef LIMIT
// #include
// #include <foo.h> 尖括号表示该文件是系统提供的，不需要写路径
// #include "foo.h" 双引号里面，表示文件是用户提供的，需要指定路径

// #if...#elif...#else...#endif
#if 0 // 该用法常用来当作注释使用
const double pi = 3.1415;
#endif
#define FOO 1

// #ifdef...#endif
// 主要用来判断是否重复加载某个库
#define EXTRA_HAPPY
#ifdef IBMPC
    #include "ibmpc.h"
#elif MAC
    #include "mac.h"
#else
    #include "general.h"
#endif

// #ifndef...#endif
// 和 #ifdef 作用相反，用来判断某个宏没有被定义过
// 防止重复加载，如果某个宏没有被定义过，则执行指定的操作
#ifndef MYHEADER_H
    #define MYHEADER_H
    #include "myheader.h"
#endif

enum colors {RED,GREEN,BLUE};
enum {ONE,TWO};

// static 的两种作用范围
// 作用于全局变量时，表示该变量只用于当前文件
// 作用于局部变量时，表示该变量的值会在函数每次执行后得到保留
static int link = 1;

int staticBlock() {
    static int staticInt = 0;
    staticInt++;
    return staticInt;
}

typedef enum {
    SHEEP,
    WHEAT,
    WOOD,
    BRICK,
    ORE
} RESOURCE;

int main() {
    // restrict
    int* restrict pt = (int*) malloc(10*sizeof(int));
    RESOURCE r = SHEEP;
    // 预处理器 preprocessor
    // SQUARE(3 + 4)如果是函数，输出的应该是49（7*7）
    // 宏是原样替换，所以替换成3 + 4*3 + 4，最后输出19
    printf("%d\n",SQUARE(3 + 4));
    printf("%d\n",SQUARE1(3 + 4));
    int MK_ID(1);
    i1 = 10;
    printf("%s\n",X(1,2,3));
    #if FOO
        printf("defined\n"); 
    #else
        printf("not defined\n");
    #endif
    #ifdef EXTRA_HAPPY
        printf("I'm extra happy!\n");
    #endif
    // 预定义宏
    printf("This function: %s\n", __func__);
    printf("This file: %s\n", __FILE__);
    printf("This line: %d\n", __LINE__);
    printf("Compiled on: %s %s\n", __DATE__, __TIME__);
    printf("C Version: %ld\n", __STDC_VERSION__);
    freopen("output.txt","w",stdout);
    printf("hello\n");
    const double PI1 = 3.14159;
    const int arr[] = {1,2,3,4};
    int* const x; // 表示指针包含的地址不可修改
    const int* y; // 表示指针指向的值不可修改
    const char* const c; // 表示内存地址以及指向的值都不可修改
    char* val = getenv("JAVA_HOME");
    printf("Value: %s\n",val);
    char* s = "春天";
    printf("%s\n",s);
    return 0;
}

