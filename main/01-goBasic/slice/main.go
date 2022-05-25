package main

import "fmt"

func main() {
	arr := [6]int{}

	for i := range arr {
		arr[i] = i
	}

	fmt.Println(arr)

	//数组切片
	fmt.Println(arr[:3])

	fmt.Println(arr[len(arr)-1])

	fmt.Println(arr)

	c := append(arr[:], 3)
	fmt.Println(c)

	// Slice 数据类型， 不定长的序列
	d := make([]string, 10)
	e := append(d, "1")
	fmt.Println(e)

}
