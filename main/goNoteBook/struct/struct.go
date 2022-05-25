package main

import "fmt"

// 定义一个结构体，声明一个结构体应该有哪些属性
type Student struct {
	name     string
	age      int
	sayHello func()
}

func main() {
	p1 := Student{name: "zyd"}
	var ip *Student // 默认空指针
	ip = &p1
	fmt.Println(ip)
}
