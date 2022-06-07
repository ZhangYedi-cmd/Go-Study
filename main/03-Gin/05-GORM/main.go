package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Student2 struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:123@(localhost)/ydl?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("connect Database error!")
	}

	db.AutoMigrate(&Student2{}) // 自动迁移表

	s1 := Student2{
		Name: "zyd",
		Age:  18,
	}

	// 增
	db.Create(&s1)

	var student Student2
	// 查
	db.First(&student, 1)

	fmt.Println(student)

	// 条件查询
	var studentList []Student2
	db.Where("age > ?", 7).Find(&studentList)
	fmt.Println(studentList)

	//改
	db.Where("name = ?", "zyd").First(&Student2{}).Update("name", "djy")

	// 删
	db.Where("name = ?", "zyd").Find(&Student2{}).Delete(&Student2{})

	// 批量删除
	//db.Where("name LIKE ?", "%djy%").Delete(&Student2{})
	// 物理删除
	db.Where("name LIKE ?", "%djy%").Unscoped().Delete(&Student2{})

	defer db.Close()

}
