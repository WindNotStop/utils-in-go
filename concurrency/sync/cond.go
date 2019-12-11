package sync

import (
	"fmt"
	"sync"
	"time"
)

func Cond() {
	cLock := sync.NewCond(&sync.Mutex{})
	queue := make(chan int, 10)

	for i := 0; i < 10; i++ {
		cLock.L.Lock()
		//检查队列长度，达到2时调用wait阻塞自己
		for len(queue) == 2 {
			cLock.Wait()
		}
		//入队
		queue <- i
		go func() {
			time.Sleep(time.Second)
			cLock.L.Lock()
			fmt.Println("出队", <-queue)
			cLock.L.Unlock()
			//唤醒单个协程
			cLock.Signal()
			//唤醒所有协程
			//cLock.Broadcast()
		}()
		cLock.L.Unlock()
	}
	fmt.Println("出队", <-queue)
	fmt.Println("出队", <-queue)
}
