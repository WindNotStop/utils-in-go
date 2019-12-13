package channel

import (
	"fmt"
	"time"
)

func SimpleForSelect(){
	//每隔一秒向通道发送一个数据
	in := make(chan int, 1)
	go func() {
		for{
			in <- 1
			time.Sleep(time.Second)
		}
	}()
	//读
	for{
		select {
			case v := <-in:
				fmt.Println(v)
		}
	}
}
