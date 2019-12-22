package simple_goroutine_pool

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	DONE = "Done"
)

type job struct {
	id int
}

func doWork(j job){
	//模拟任务执行
	fmt.Println(j.id)
	//监控协程数
	fmt.Println("goroutine num:",runtime.NumGoroutine())
}

func simplePool(num int, done chan interface{}, jobChan <-chan job) <-chan string{
	retChan := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i:=0;i<num;i++{
		go func() {
			defer wg.Done()
			for {
				select {
				case <-done:
					retChan <- DONE
					return
				case j := <-jobChan:
					doWork(j)
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(retChan)
	}()
	return retChan
}

func SimpleGoroutinePool() {
	//用于输入任务的chan
	jobChan := make(chan job)
	done := make(chan interface{})
	//模拟任务输入
	go func() {
		for i := 0;i<1000;i++ {
			jobChan <- job{i}
		}
		close(done)
	}()
	//建立一个大小为10的协程池
	retChan := simplePool(10,done,jobChan)
	for r := range retChan{
		fmt.Println("res:",r)
	}
}

