package main

import (
	"fmt"
	"math"
	"strconv"
)

// Shape 接口定义,包括面积周长
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return r.Width*2 + r.Height*2
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Area 计算圆形面积
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Perimeter 计算圆形周长
func (c Circle) Perimeter() float64 {
	return math.Pi * c.Radius * 2
}

func main22() {
	// ---------------------------------------------------------------------------- 题目1 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
	//		 在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
	// 考察点 ：接口的定义与实现、面向对象编程风格。

	re := Rectangle{Width: 1, Height: 2}
	cr := Circle{Radius: 2}
	fmt.Println(re.Area())
	fmt.Println(re.Perimeter())
	fmt.Println("\n", cr.Area())
	fmt.Println(cr.Perimeter())
	// **** 利用接口实现多态
	var shape Shape
	shape = re
	re.Area()
	re.Perimeter()
	shape = cr
	shape.Area()
	shape.Perimeter()

	// --------------------------------------------------------------------------- 题目2 ：用组合方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
	//		为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
	// 考察点 ：组合的使用、方法接收者。
	em := Employee{
		Person: Person{
			Name: "s",
			Age:  20,
		},
		EmployeeId: 1,
	}
	fmt.Println("\n\n\n", em.printInfo())

}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeId int
}

func (e Employee) printInfo() string {
	return e.Name + "" + strconv.Itoa(e.Age)
}
