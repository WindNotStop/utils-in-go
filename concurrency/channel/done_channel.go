package channel

import (
	"fmt"
	"time"
)

func DoneChannel(){
	//用来终止协程
	done := make(chan interface{})
	//样例输入
	in := []interface{}{1,2,3,4,5}
	//流处理
	simpleStream := func(done <-chan interface{}, in ...interface{}) <-chan interface{} {
		results := make(chan interface{})
		go func() {
			defer close(results)
			defer fmt.Println("goroutine closed")
			for _, v := range in{
				select {
				case results <- v:
					//模拟实际操作
					time.Sleep(time.Second)
				case <-done:
					return
				}
			}
		}()
		return results
	}
	//消费(打印)结果数据
	consumer := func(done chan interface{}, results <-chan interface{}) {
		for v := range results{
			//关闭doneChannel
			if v.(int) == 2{
				close(done)
			}
			fmt.Println(v)
		}
	}

	results := simpleStream(done, in...)
	consumer(done, results)
}
