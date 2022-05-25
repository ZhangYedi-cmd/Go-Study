package main

import "fmt"

func sayName() {
	var a string = "zyd"
	fmt.Println(a)

	var (
		name string
	)
	name = "cmy"
	fmt.Println(name)

	if name == "cmy" {
		// 占位符
		fmt.Println("name is %s", name)
	}
	// 声明数组
	strings := []string{"google", "Chrome"}
	// for 循环
	for index, str := range strings {
		fmt.Println(index, str)
	}

	// 声明一个map
	nameMap := make(map[string]string)

	nameMap["name"] = "zyd"
	nameMap["age"] = "18"

	// 遍历Map
	for key, value := range nameMap {
		fmt.Println(key, value)
	}
}

func main() {
	sayName()
}
