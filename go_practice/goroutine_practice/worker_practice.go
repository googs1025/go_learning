package main

import (
	"fmt"
	"math/rand"
)

/*
	worker pool(goroutine池)

	生产者消费者模型
	可以有效控制goroutine数量
	需求:
		计算一个数字的各个位数之和，例如数字123 结果是1+2+3=6
		随机生成数字进行计算
 */

type Job struct {
	// 工作id号
	Id 		int
	// 计算的随机数
	RandNum	int
}


type Result struct {
	// 必须传对象实例
	job 	*Job
	// 求和
	sum 	int
}



func main() {

	// 需要两个chan, 注意通道里放的都是对象实例
	// 1.job 通道
	jobChan := make(chan  *Job, 128)
	// 2.结果通道
	resultChan := make(chan *Result, 128)
    // 3.创建工作池
	createPool(64, jobChan, resultChan)


	// 开启个打印结果的协程
	go func(resultChan chan *Result) {
		// 从通道取数据
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}

	}(resultChan)

	var id int // 生成新job id
	// 不断取出数据 放入
	for {
		id++

		r_num := rand.Int()
		job := &Job{
			Id: id,
			RandNum: r_num,

		}

		jobChan <- job
	}

}

// 创建工作池
/*
	num:开启goroutine数量
	jobChan:job通道
	resultChan:result通道
 */
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 直接开启数个协程!! (用for开启)
	for i := 0; i < num; i++ {

		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行计算
			// 遍历job管道所有数据，进行相加
			// 主要的业务逻辑
			for job := range jobChan {
				r_num := job.RandNum

				var sum int // 定义一个新变量
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10

				}
				// 记录计算结果
				r := &Result{
					job: job,
					sum: sum,
				}
				// 放回通道中
				resultChan <- r
			}

		}(jobChan, resultChan)




	}

}
