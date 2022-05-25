package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串常用的方法
	fmt.Println("hello string")
	a := "hello"
	fmt.Println(strings.Contains(a, "ell"))                           // true
	fmt.Println(strings.Count(a, "l"))                                // 2
	fmt.Println(strings.Index(a, "he"))                               // 0
	fmt.Println(strings.Join([]string{"h", "e", "l", "l", "o"}, "-")) //h-e-l-l-o
	//fmt.Println(strings.Cut(a, "h"))                                  // ello true

	//var (
	//	str string
	//	res bool
	//)
	before, after, res := strings.Cut(a, "h")
	fmt.Println(before, after, res)
	fmt.Println(a[:3]) //字符串切片
	//strings.Index()
}
