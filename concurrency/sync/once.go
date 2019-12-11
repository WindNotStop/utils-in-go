package sync

import (
	"fmt"
	"sync"
)

func Once() {
	var num int
	once := sync.Once{}

	inc := func() {
		num++
	}

	dec := func() {
		num--
	}

	for i := 0; i < 100; i++ {
		once.Do(inc)
	}

	//once计算的是Do的调用次数，故不会执行dec函数
	once.Do(dec)
	fmt.Println(num)
}
