package channel

import "fmt"

func LimitedPipeline(){
	in := []int{1,2,3,4,5}
	done := make(chan interface{})
	defer close(done)
	//流生成器,将in切片流化
	generator := func(done <-chan interface{},in []int) <-chan int{
		inStream := make(chan int)
		go func() {
			defer close(inStream)
			for _, v := range in{
				select {
				case inStream <-v:
				case <-done:
					return
				}
			}
		}()
		return inStream
	}
	//num限制流输出量
	limitedStream := func(done <-chan interface{},inStream <-chan int,num int) <-chan int{
		results := make(chan int)
		go func() {
			defer close(results)
			for i:=0;i<num;i++{
				select {
				case results<-<-inStream:
				case <-done:
					return
				}
			}
		}()
		return results
	}

	for v := range limitedStream(done,generator(done,in),3){
		fmt.Println(v)
	}
}

