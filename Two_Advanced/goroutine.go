package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func goToOne() {
	tStart := time.Now()
	fmt.Println("One线程启动时间:", tStart)
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("One线程-奇数:", i)
		}
	}
	tEnd := time.Now()
	fmt.Println("One线程结束时间:", tEnd)
	fmt.Println("One线程总用时:", tEnd.Sub(tStart))
}

func goToTwo() {
	tStart := time.Now()
	fmt.Println("Two线程启动时间:", tStart)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("Two线程-偶数:", i)
		}
	}
	tEnd := time.Now()
	fmt.Println("Two线程结束时间:", tEnd)
	fmt.Println("Two线程总用时:", tEnd.Sub(tStart))
}

func main23() {
	// --------------------------------------------- 题目1 ：go关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	// 考察点 ： go 关键字的使用、协程的并发执行。
	//go goToOne()
	//go goToTwo()

	// 函数启动需要给足够的时间执行线程任务,否则会发生线程未执行,主函数已启动结束！
	//time.Sleep(1000 * time.Millisecond)

	// ------------------------------------------------ 题目2 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	// 考察点 ：协程原理、并发任务调度。

	// 添加执行任务
	addScheduler := TaskScheduler{tasks: []Task{}}
	for i := 0; i < 5; i++ {
		taskName := "任务" + strconv.Itoa(i)
		addScheduler.NewTaskScheduler(taskName, func() {
			time.Sleep(1 * time.Second)
			fmt.Printf("《《《《《《《《《《《《《《《 添加任务 %d \n", i)
		})
	}

	// 等待所有任务完成的同步机制
	var wg sync.WaitGroup
	taskResults := TaskResults{}

	// 每个任务执行创建单独线程 收集时间
	for _, task := range addScheduler.tasks {
		wg.Add(1) // 需要等待任务数量

		go func(t Task) {
			defer wg.Done() // wg.Done() 将WaitGroup计数器减1，defer保证无论函数任何情况运行，该语句都会被执行，确保计数器正确递减，避免 WaitGroup 永远等待的情况。

			// 开始执行任务
			startTime := time.Now()
			defer func() {
				if recover() != nil {
					fmt.Println("执行失败了,任务名称:", task.Name)
				}

				endTime := time.Now()
				taskResults.addResult(TaskResult{
					TaskName:     task.Name,
					StartTime:    startTime,
					EndTime:      endTime,
					DurationTime: endTime.Sub(startTime),
				})
			}()

			// 实际执行任务
			t.Func()
		}(task)
	}

	// 等待所有任务完成后再打印结果
	wg.Wait()
	for _, taskResult := range taskResults.TaskResults {
		fmt.Printf("任务%s,执行时间为%d\n", taskResult.TaskName, taskResult.DurationTime)
	}
}

type Task struct {
	Name string
	Func func()
}

type TaskResult struct {
	TaskName     string
	StartTime    time.Time
	EndTime      time.Time
	DurationTime time.Duration
}

type TaskResults struct {
	TaskResults []TaskResult
}

type TaskScheduler struct {
	tasks []Task
}

func (scheduler *TaskScheduler) NewTaskScheduler(name string, Func func()) {
	scheduler.tasks = append(scheduler.tasks, Task{name, Func})
}

func (result *TaskResults) addResult(r TaskResult) {
	result.TaskResults = append(result.TaskResults, r)
}
