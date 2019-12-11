package sync

import (
	"fmt"
	"sync"
	"time"
)

func Mutex() {
	lock := sync.Mutex{}
	var num int

	//inc
	go func() {
		lock.Lock()
		defer lock.Unlock()
		num++
		fmt.Println("inc:", num)
	}()

	//dec
	go func() {
		lock.Lock()
		defer lock.Unlock()
		num--
		fmt.Println("dec:", num)
	}()

	time.Sleep(time.Second)
	fmt.Println(num)
}
