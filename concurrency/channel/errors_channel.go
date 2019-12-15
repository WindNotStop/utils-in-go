package channel

import (
	"errors"
	"fmt"
)

type ResWithErr struct {
	res int
	err error
}

type arith struct {
	a int
	b int
}

func ErrorsChannel() {
	//输入
	ariths := []arith{{1, 1}, {2, 0}, {3, 2}, {4, 2}, {10, 3}}
	done := make(chan interface{})
	defer close(done)
	//计算逻辑
	errors := func(done <-chan interface{}, ariths ...arith) <-chan ResWithErr {
		results := make(chan ResWithErr)
		go func() {
			defer close(results)
			for _, v := range ariths {
				var result ResWithErr
				//如果除数为0或不能整除则报错
				if v.b == 0 || v.a%v.b != 0 {
					result.err = errors.New(fmt.Sprintf("input error : a = %d, b = %d", v.a, v.b))
				} else {
					result.res = v.a / v.b
				}
				select {
				case results <- result:
				case <-done:
					return
				}
			}
		}()
		return results
	}
	results := errors(done, ariths...)
	for res := range results {
		if res.err != nil {
			fmt.Println(res.err)
		} else {
			fmt.Println(res.res)
		}
	}
}
