package channel

import "fmt"

func SimpleChannel(){
	//创建一个带缓冲为1的channel
	myChan := make(chan interface{}, 1)
	//写
	myChan <- 1
	//读
	v, ok := <-myChan
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("channel closed")
	}
}

