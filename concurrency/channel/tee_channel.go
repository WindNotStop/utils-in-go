package channel

import (
	"fmt"
	"sync"
)

func TeeChannel(){
	in := []interface{}{1,2,3,4,5}
	wg := sync.WaitGroup{}
	done := make(chan interface{})
	defer close(done)
	//返回out1,out2复制副本
	tee := func(done <-chan interface{},in ...interface{}) (_, _<-chan interface{}) {
		out1 := make(chan interface{})
		out2 := make(chan interface{})
		go func() {
			defer close(out1)
			defer close(out2)
			for _,val := range in {
				var out1, out2 = out1, out2
				//如果去掉for循环就变成了fan out模式的生成器
				for i := 0; i < 2; i++ {
					select {
					case <-done:
					//nil用来阻塞已经复制完成的channel
					case out1 <- val:
						out1 = nil
					case out2 <- val:
						out2 = nil
					}
				}
			}
		}()
		return out1, out2
	}
	res1, res2 := tee(done,in...)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range res1{
			fmt.Println(v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range res2{
			fmt.Println(v)
		}
	}()
	wg.Wait()
}