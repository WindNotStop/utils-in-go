package sync

import (
	"fmt"
	"sync"
	"time"
)

func RWMutex() {
	lock := sync.RWMutex{}
	var num int

	//read
	go func() {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("read:", num)
	}()

	//write
	go func() {
		lock.Lock()
		defer lock.Unlock()
		num++
		fmt.Println("write inc:", num)
	}()

	time.Sleep(time.Second)
	fmt.Println(num)
}
