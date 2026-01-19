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
// 指针接收者：*rect
func (r *rect) area() int {
	r.width = 10 // 修改原对象的字段
	return r.width * r.height
}

// 值接收者：rect（无*）
func (r rect) perimeter() int {
	r.height = 20 // 仅修改副本，原对象不受影响
	return 2 * (r.width + r.height)
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
	p3 := newPerson("Tom")
	fmt.Println(p3)

	// struct 结构体的字面量表达方式
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

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

	r := rect{width: 5, height: 5}
	// 调用指针接收者方法，原对象被修改
	fmt.Println("area:", r.area())    // 输出：area: 50（width被改为10）
	fmt.Println("原r.width:", r.width) // 输出：原r.width: 10（原对象已变）

	// 调用值接收者方法，仅修改副本
	fmt.Println("perimeter:", r.perimeter()) // 输出：perimeter: 60（height副本改为20）
	fmt.Println("原r.height:", r.height)      // 输出：原r.height: 5（原对象不变）
}
