package channel

import "fmt"

func SimpleChannel(){
	//创建一个带缓冲为1的channel
	myChan := make(chan int, 1)
	//写
	myChan <- 1
	//读
	v := <-myChan
	fmt.Println(v)
}

