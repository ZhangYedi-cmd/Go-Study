package main

import "fmt"

type People struct {
	name string
	age  int
	//sayHello func()
}

type user struct {
	name     string
	password string
}

func checkPassword(u user, pwd string) bool {
	return u.password == pwd
}
func checkPassword2(u *user, pwd string) bool {
	// 空指针
	if u == nil {
		fmt.Println("13213123")
		return false
	}
	return u.password == pwd
}

func main() {
	// 指针学习
	var p1 *People

	// 指针取值
	p1 = &People{name: "zyd", age: 18}

	fmt.Println(p1)

	// 获取属性
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	u := user{name: "zyd", password: "123"}
	//var u2 *user
	println(checkPassword(u, "123"))
	println(checkPassword2(&u, ""))
}
