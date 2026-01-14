package main

import "fmt"

// -- struct 结构体 --
type person struct {
	name string
	age  int
}

// 函数返回值是struct的指针
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

type rect struct {
	width, height int
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

func struct结构() {
	fmt.Println("-- struct结构 --")
	p1 := person{name: "Sean", age: 50}
	fmt.Println(p1.name)
	// 指针的使用
	p2 := &p1
	fmt.Println(p2.age)
	p2.age = 51
	fmt.Println(p2.age)

	// struct 结构体的字面量表达方式
	dog := struct {
		name   string
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
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	fmt.Println(co)
	// 可以直接调用内嵌的结构体的方法
	fmt.Println(co.describe())
	// 内嵌结构体的方法可以看作是结构体自己的方法，算直接实现接口
	type describer interface {
		describe() string
	}
	var desc describer = co
	fmt.Println("describer:", desc.describe())
}
