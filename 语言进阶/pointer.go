package main

import (
	"fmt"
)

type Num struct {
	num int
}

func (n *Num) getNum() int {
	return n.num
}

// 定义函数,接受整数指针作为参数
func repSlice(num int) int {
	// 通过指针访问并修改返回的值
	return num + 10
}

func repSliceTwo(numArrSlice []int) {
	for i := range numArrSlice {
		(numArrSlice)[i] *= 2
	}
}

func main() {
	// --------------------------------------------  题目1 ：定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
	// 考察点 ：指针的使用、值传递与引用传递的区别。
	value := Num{num: 1}
	fmt.Printf("题目1,修改前值：%d", value.getNum())
	repSlice(value.getNum()) // 调用函数,传递变量的地址
	fmt.Println("题目12,修改后值：", repSlice(2))
	fmt.Println("题目1,修改后值：", value)
	fmt.Printf("题目1,初始化调用值: %d\n\n\n\n", value.getNum())

	// ---------------------------------------------- 题目2 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
	// 考察点 ：指针运算、切片操作。
	numArr := []int{1, 2, 3, 4, 5}
	fmt.Printf("题目2,修改前值numArr：%d\n", numArr)
	numArrSlice := numArr[:]
	fmt.Println("题目2,切片值：", numArrSlice)
	repSliceTwo(numArrSlice)
	fmt.Println("题目2,修改后原数组值：", numArr)
	fmt.Println("题目2,修改后切片数组值：", numArrSlice)
}
