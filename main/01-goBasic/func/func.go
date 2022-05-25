package main

import "fmt"

// func 函数名 （参数1 数据类型）返回值的数据类型
func myFunc(a int, b string) string {
	return string(a) + b
}

// 返回多个参数
func func2(a, b int) (string, int) {
	return string(a), b
}

// 声明类型
type cb func()

// 闭包
func wrapper(work func()) func() {
	return func() {
		fmt.Println("我是新增的功能呀")
		work()
	}
}

// 指针传参
func pointArgs(p1 *int, p2 *int) int {
	var res int = 0
	res += *p1
	res += *p2
	return res
}

func sayHello() {
	fmt.Println("hello word")
}

func main() {
	//Go 语言函数
	//fmt.Println(myFunc(1, "B"))

	// 闭包函数
	//wrapper(sayHello)()

	// 指针传参
	var p1 int = 1
	var p2 int = 2

	// &取内存地址（真正的值）
	// var p1 *int  声明一个int指针
	println(pointArgs(&p1, &p2))
}
