package channel

import (
	"fmt"
	"time"
)

func SimpleForSelect() {
	//每隔一秒向通道发送一个数据
	in := make(chan interface{}, 1)
	go func() {
		for {
			in <- 1
			time.Sleep(time.Second)
		}
	}()
	//超时计时器
	idleDuration := 2 * time.Second
	idleDelay := time.NewTimer(idleDuration)
	defer idleDelay.Stop()
	//读
	for {
		idleDelay.Reset(idleDuration)
		select {
		case v := <-in:
			fmt.Println(v)
		//超时退出
		case <-idleDelay.C:
			return
		}
	}
}
