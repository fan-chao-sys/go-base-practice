package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type shareCount struct {
	mu    sync.Mutex
	count int
}

func (s *shareCount) add() {
	s.mu.Lock()
	s.count += 1
	s.mu.Unlock()
}

func (s *shareCount) get() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

func main41() {
	// --------------------------------------------------------------------------            题目1 ：用 sync.Mutex 保护一个共享计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	// 考察点 ： sync.Mutex 的使用、并发数据安全。

	// 等待所有线程组
	wt := sync.WaitGroup{}
	sc := shareCount{sync.Mutex{}, 0}
	fmt.Printf("多线程共享下计数值,未执行前为:%d\n", sc.get())

	for i := 0; i < 10; i++ {
		wt.Add(1) // 等待线程 + 1

		fmt.Printf("<<<<<<<<< 多线程第 %d 个,执行!!\n", i+1)
		go func() {
			defer wt.Done() // 线程执行完一定减1

			if r := recover(); r != nil {
				fmt.Println("panic:", r)
			}

			// 调用方法 递增1000次
			for i := 0; i < 1000; i++ {
				sc.add()
			}
		}()
	}

	// 等待所有线程执行完后再执行下面方法
	wt.Wait()
	fmt.Printf("多线程共享下计数值,执行后为:%d\n\n", sc.get())

	// -----------------------------------------------------------------    题目2 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	// 考察点 ：原子操作、并发数据安全。

	var num int64 = 0
	fmt.Printf("多线程下无锁原子计数器-初始值:%d\n", num)

	for i := 0; i < 10; i++ {
		wt.Add(1)
		go func() {
			defer wt.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt64(&num, 1)
			}
		}()
	}

	wt.Wait()
	fmt.Printf("多线程下无锁原子计数器-结果值:%d\n", num)
}
