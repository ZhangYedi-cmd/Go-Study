package main

import "fmt"

func hello() {
	// 声明变量
	var a_ string = "Runoob"
	fmt.Println(a_)
	// int 类型 不初始化默认为0
	var num int
	//num = 1
	fmt.Println(num)

	// bool 默认为false
	var res bool
	fmt.Println(res)

	// 事先声明变量？
	//var num2 int
	num2 := 2 // 如果变量已经被声明过了，则使用 := 赋值会报错
	// var num2 int = 2
	fmt.Println(num2)

	/**
	多变量声明
	*/
	var vname1, vname2, vname3 string
	vname1, vname2, vname3 = "v1", "v2", "v3"
	fmt.Println(vname3, vname2, vname1)

	// 因式分解法声明变量
	var (
		name1 string
		age   int
	)
	name1 = "zyd"
	age = 20
	fmt.Println(name1, age)
}

func main() {
	hello()
	fmt.Println("Hello, World!")
}
