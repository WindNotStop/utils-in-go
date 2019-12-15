package channel

import "fmt"

func LimitedPipeline(){
	in := []interface{}{1,2,3,4,5}
	limitedNum := 3
	done := make(chan interface{})
	defer close(done)
	//流生成器,将in切片流化
	generator := func(done <-chan interface{},in ...interface{}) <-chan interface{}{
		inStream := make(chan interface{})
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
	limitedStream := func(done <-chan interface{},inStream <-chan interface{},limitedNum int) <-chan interface{}{
		results := make(chan interface{})
		go func() {
			defer close(results)
			for i:=0;i<limitedNum;i++{
				select {
				case results<-<-inStream:
				case <-done:
					return
				}
			}
		}()
		return results
	}

	for v := range limitedStream(done,generator(done,in...),limitedNum){
		fmt.Println(v)
	}
}
