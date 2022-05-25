package main

import "fmt"

func main() {
	// 创建数组
	var arr [5]int // 没赋值的默认为0
	arr[4] = 10
	fmt.Println(arr)

	b := [10]int{1, 3, 5, 3, 8, 7}
	fmt.Println(b) // 数组赋值

	// 二维数组
	var twoArr [4][4]int
	for i := range twoArr {
		for j := 0; j < len(twoArr[0]); j++ {
			twoArr[i][j] = i + j
		}
	}

	fmt.Println(twoArr)

}
