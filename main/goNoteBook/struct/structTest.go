package main

import "fmt"

func myFunc2() {
	fmt.Println("this is func2 ")
}

// 定义一个结构体
type Animal struct {
	name string
	age  int
}

func main() {
	//myFunc2()
	// 实现一个结构体
	var point *Animal
	var p = Animal{
		name: "lyz",
		age:  10,
	}
	point = &p
	fmt.Println(point)
}
