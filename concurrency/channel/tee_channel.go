package channel

import (
	"fmt"
	"sync"
)

func TeeChannel() {
	in := []interface{}{1, 2, 3, 4, 5}
	wg := sync.WaitGroup{}
	done := make(chan interface{})
	defer close(done)
	//返回out1,out2...outn复制副本
	tee := func(done <-chan interface{}, n int, in ...interface{}) []chan interface{} {
		outs := make([]chan interface{}, n)
		for i, _ := range outs {
			outs[i] = make(chan interface{})
		}
		go func() {
			defer func() {
				for _, v := range outs {
					close(v)
				}
			}()
			for _, val := range in {
				//如果去掉for循环就变成了fan out模式的生成器
				for i := 0; i < n; i++ {
					select {
					case <-done:
						return
					//nil用来阻塞已经复制完成的channel
					case outs[i] <- val:
					}
				}
			}
		}()
		return outs
	}
	res := tee(done, 3, in...)
	wg.Add(3)

	go func() {
		defer wg.Done()
		for v := range res[0] {
			fmt.Println(v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range res[1] {
			fmt.Println(v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range res[2] {
			fmt.Println(v)
		}
	}()
	wg.Wait()
}
