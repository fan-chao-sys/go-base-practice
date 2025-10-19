package 实用练习

import (
	"fmt"
	"time"
)

// Task 定义任务类型，接收任务名称和任务函数
type Task struct {
	Name string
	Func func()
}

// TaskResult 定义任务执行结果，包含任务名称、开始时间、结束时间和执行时长
type TaskResult struct {
	TaskName     string
	StartTime    time.Time
	EndTime      time.Time
	Duration     time.Duration
	ErrorMessage string
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks []Task
}

// NewScheduler 创建新的任务调度器
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: []Task{},
	}
}

// AddTask 向调度器添加任务
func (s *Scheduler) AddTask(name string, taskFunc func()) {
	s.tasks = append(s.tasks, Task{
		Name: name,
		Func: taskFunc,
	})
}

// Run 执行所有任务并返回结果
func (s *Scheduler) Run() []TaskResult {
	results := make([]TaskResult, len(s.tasks))
	resultChan := make(chan TaskResult, len(s.tasks))

	// 为每个任务启动一个协程
	for _, task := range s.tasks {
		go func(t Task) {
			startTime := time.Now()
			var errMsg string

			// 执行任务并捕获可能的恐慌
			defer func() {
				if r := recover(); r != nil { // recover() 内置函数,用于捕获和处理程序中的 panic（恐慌），防止程序因未处理的异常而直接崩溃
					errMsg = fmt.Sprintf("任务执行出错: %v", r)
				}

				endTime := time.Now()
				resultChan <- TaskResult{
					TaskName:     t.Name,
					StartTime:    startTime,
					EndTime:      endTime,
					Duration:     endTime.Sub(startTime),
					ErrorMessage: errMsg,
				}
			}()

			t.Func()
		}(task)
	}

	// 收集所有任务结果
	for i := 0; i < len(s.tasks); i++ {
		results[i] = <-resultChan
	}

	return results
}

func main() {
	// 创建调度器
	scheduler := NewScheduler()

	// 添加示例任务
	scheduler.AddTask("任务1", func() {
		time.Sleep(1 * time.Second) // 模拟任务执行时间
		fmt.Println("任务1执行完毕")
	})

	scheduler.AddTask("任务2", func() {
		time.Sleep(2 * time.Second)
		fmt.Println("任务2执行完毕")
	})

	scheduler.AddTask("任务3", func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("任务3执行完毕")
	})

	// 添加一个可能出错的任务
	scheduler.AddTask("可能出错的任务", func() {
		time.Sleep(1500 * time.Millisecond)
		// 模拟一个错误
		panic("发生了一个意外错误") // panic()内置函数，用于停止当前函数正常执行流程，并开始恐慌过程
	})

	// 执行所有任务
	fmt.Println("开始执行所有任务...")
	startTime := time.Now()
	results := scheduler.Run()
	totalDuration := time.Now().Sub(startTime)

	// 输出结果统计
	fmt.Println("\n所有任务执行完成，统计结果如下：")
	fmt.Printf("总执行时间: %v\n\n", totalDuration)

	for _, result := range results {
		fmt.Printf("任务名称: %s\n", result.TaskName)
		fmt.Printf("开始时间: %s\n", result.StartTime.Format("2006-01-02 15:04:05.000"))
		fmt.Printf("结束时间: %s\n", result.EndTime.Format("2006-01-02 15:04:05.000"))
		fmt.Printf("执行时长: %v\n", result.Duration)

		if result.ErrorMessage != "" {
			fmt.Printf("状态: 失败 - %s\n", result.ErrorMessage)
		} else {
			fmt.Println("状态: 成功")
		}
		fmt.Println("------------------------")
	}
}
