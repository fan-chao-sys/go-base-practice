package main

import (
	"encoding/json"
	"fmt"
	"go-base-practice/utils"
)

// 基本CRUD操作
// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 要求 ：
//  编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
//  编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
//  编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
//  编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

type Student struct {
	ID     uint `gorm:"primary_key;auto_increment"`
	Name   string
	Age    int
	Gender string
}

func main() {
	stu := []Student{{
		Name:   "小郭",
		Age:    20,
		Gender: "一年级",
	}, {
		Name:   "小范",
		Age:    10,
		Gender: "二年级",
	}, {
		Name:   "张三",
		Age:    20,
		Gender: "三年级",
	}}
	// 初始化表
	utils.DB().AutoMigrate(&stu)
	// 创建数据
	//DB().Save(&stu)
	// 查询
	var arrStu []Student
	utils.DB().Debug().Model(&stu).Where("age > ?", 18).Scan(&arrStu)
	constructionLogin(arrStu, 1, "查询出")

	//  编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	utils.DB().Model(&Student{}).Where("Name=?", "张三").Update("Gender", "四年级")
	constructionLogin(arrStu, 2, "更新后")

	//  编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
	utils.DB().Where("age < ?", 15).Delete(&Student{})
	constructionLogin(arrStu, 3, "删除后")
}

func constructionLogin(arrStu []Student, num int, s string) {

	for i := range arrStu {
		jst, _ := json.Marshal(arrStu[i])
		fmt.Printf(s+"%s\n", string(jst))
	}
}
