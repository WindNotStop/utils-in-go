package channel

import "fmt"

func SimpleChannel() {
	//创建一个带缓冲为1的channel
	simple := make(chan interface{}, 1)
	//写
	simple <- 1
	//读
	v, ok := <-simple
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("channel closed")
	}
}
