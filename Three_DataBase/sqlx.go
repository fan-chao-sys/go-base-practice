package main

import (
	"encoding/json"
	"fmt"
	"go-base-practice/utils"
)

type Employee struct {
	Name       string
	Department string
	Salary     float64
}

type Books struct {
	ID     uint    `db:"id"` // 标签 db:"id"：指定该字段对应数据库表中的 id 列
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	// 使用SQL扩展库进行查询
	// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
	// 要求 ：
	//  编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	//  编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
	//funcName()
	//funcNameSelect()

	// 实现类型安全映射
	//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
	//要求 ：
	//定义一个 Book 结构体，包含与 books 表对应的字段。
	//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
	//createBooks()
	var arrBooks []Books
	utils.DB().Model(&Books{}).Where("Price > ?", 50).Scan(&arrBooks)
	for i := range arrBooks {
		bjs, _ := json.MarshalIndent(arrBooks[i], "", " ")
		fmt.Println("book大于50书集:" + string(bjs))
	}
}

func createBooks() {
	book := []Books{
		{
			Title:  "安徒生童话",
			Author: "安徒生",
			Price:  500,
		},
		{
			Title:  "收破烂的毛毛球",
			Author: "小毛",
			Price:  100,
		},
		{
			Title:  "简爱",
			Author: "安黛因",
			Price:  20,
		},
	}
	utils.DB().AutoMigrate(&Books{})
	utils.DB().Create(&book)
}

func funcNameSelect() {
	var arrEm []Employee
	utils.DB().Debug().Model(&Employee{}).Where("Department = ?", "技术部").Scan(&arrEm)
	for i := range arrEm {
		js, _ := json.MarshalIndent(arrEm[i], "", " ")
		fmt.Println("技术部员工:%s \n " + string(js))
	}
	var emp Employee
	utils.DB().Debug().Model(&Employee{}).Order("Salary Desc").First(&emp)
	json, _ := json.Marshal(emp)
	fmt.Println("最大薪资员工：", string(json))
}

func funcName() []Employee {
	employee := []Employee{
		{
			Name:       "John1",
			Department: "信息部",
			Salary:     100,
		},
		{
			Name:       "John2",
			Department: "技术部",
			Salary:     2000,
		},
		{
			Name:       "John3",
			Department: "技术部",
			Salary:     1000,
		},
		{
			Name:       "John4",
			Department: "财务部",
			Salary:     1000,
		}, {
			Name:       "John5",
			Department: "信息部",
			Salary:     500,
		},
	}

	err := utils.DB().AutoMigrate(&Employee{})
	if err != nil {
		return nil
	}
	utils.DB().Create(&employee)
	return employee
}
