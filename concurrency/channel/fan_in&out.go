package channel

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func FanInOut(){
	//预设并发副本数
	num := runtime.NumCPU()
	fmt.Println("num:",num)
	done := make(chan interface{})
	defer close(done)
	//模拟高耗时的操作
	outStream := func(done <-chan interface{},in ...interface{})<-chan interface{} {
		out := make(chan interface{})
		go func() {
			defer close(out)
			for _,v := range in{
				time.Sleep(time.Second)
				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}()
		return out
	}
	outs := make([]<-chan interface{}, num)
	for i:=0;i<num;i++{
		outs[i] = outStream(done,i)
	}
	//合并输出，即扇入
	inStream := func(done <-chan interface{},outs []<-chan interface{})<-chan interface{} {
		results := make(chan interface{})
		wg := sync.WaitGroup{}
		wg.Add(len(outs))
		//多路并发器
		multiplex := func(out <-chan interface{}) {
			defer wg.Done()
			for v := range out {
				select {
				case results<-v:
				case <-done:
					return
				}
			}
		}
		//并发操作，即扇出
		for _,v := range outs {
			go multiplex(v)
		}
		go func() {
			wg.Wait()
			close(results)
		}()
		return results
	}

	t := time.Now()
	results := inStream(done, outs)
	for v := range results{
		fmt.Println(v)
	}
	fmt.Println(time.Since(t))
}

