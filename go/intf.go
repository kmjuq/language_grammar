package main

import (
	"fmt"
	"math"
)

// -- 接口 --
type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) toGeometry() geometry {
	return c
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func 接口结构() {
	fmt.Println("-- 接口 --")
	circle1 := circle{radius: 5}
	measure(circle1)

	// 类型不为空，值为 nil 的接口
	var c *circle = nil
	var g geometry = c
	fmt.Println(g == nil)
}
