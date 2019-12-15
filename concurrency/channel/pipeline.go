package channel

import "fmt"

func Pipeline() {
	//样例输入
	in := []interface{}{1, 2, 3, 4, 5}
	//流处理
	simpleStream := func(in ...interface{}) <-chan interface{} {
		results := make(chan interface{})
		go func() {
			defer close(results)
			for _, v := range in {
				results <- v
			}
		}()
		return results
	}
	//消费(打印)结果数据
	consumer := func(results <-chan interface{}) {
		for v := range results {
			fmt.Println(v)
		}
	}

	results := simpleStream(in...)
	consumer(results)
}
