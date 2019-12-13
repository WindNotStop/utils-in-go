package channel

import "fmt"

func Pipeline(){
	//样例输入
	in := []int{1,2,3,4,5}
	//流处理，将样例的值加1
	simpleStream := func(in []int) <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for _, v := range in{
				results <- v + 1
			}
		}()
		return results
	}
	//消费(打印)结果数据
	consumer := func(results <-chan int) {
		for v := range results{
			fmt.Println(v)
		}
	}

	results := simpleStream(in)
	consumer(results)
}
