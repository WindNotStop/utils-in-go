package main

import (
	"fmt"
	"sync/atomic"
)

func CAS() {
	var temp int64 = 233
	old := temp
	var new int64 = 2333
	//比较temp和old，如果相同，则把new赋值给temp
	if ok := atomic.CompareAndSwapInt64(&temp, old, new); ok {
		fmt.Println(temp)
	}
}
