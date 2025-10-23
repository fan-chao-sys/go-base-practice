package main

import (
	"fmt"
	"sync"
)

func main() {
	// 题目1 ：用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
	// 考察点 ：通道的基本使用、协程间通信。
	//numChan := make(chan int)
	//go func() {
	//	for i := 1; i <= 10; i++ {
	//		numChan <- i
	//		fmt.Printf("发生通道值:%d \n", i)
	//	}
	//	close(numChan)
	//}()
	//
	//go func() {
	//	for v := range numChan {
	//		fmt.Printf("接受通道值:%d \n", v)
	//	}
	//}()
	//time.(1000 * time.Second)

	// 等待线程完成,阻塞主线程等待用户输入后退出
	//var str string
	//_, err := fmt.Scanln(&str)
	//if err != nil {return}

	// 题目2 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
	// 考察点 ：通道的缓冲机制。
	// 创建容量 10 通道
	chanIntTen := make(chan int, 10)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go producer(chanIntTen, &wg)
	go consumer(chanIntTen, &wg)
	wg.Wait()
}

// 生产者:向带缓冲通道发送100个整数
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		ch <- i
		if i%10 == 0 {
			fmt.Printf("已发送第 %d 个数据,当前缓冲区大小:%d\n", i, len(ch))
		}
	}
	close(ch)
	fmt.Println("producer done", len(ch))
}

// 消费者: 从通道接受数据打印
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	count := 0
	for num := range ch {
		count++
		if count%20 == 0 {
			fmt.Printf("已接受第:%d 个数据:%d,当前缓冲区大小: %d\n", count, num, len(ch))
		}
	}
	fmt.Println("consumer done", count)
}
