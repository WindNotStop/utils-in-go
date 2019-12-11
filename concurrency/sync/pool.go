package sync

import (
	"fmt"
	"sync"
)

//模拟连接资源
type conn struct {
	ip   string
	port int
}

func Pool() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("create a conn")
			return conn{
				ip:   "localhost",
				port: 5672,
			}
		},
	}
	//获取一个实例，没有就new一个
	conn := pool.Get()
	fmt.Println("conn1:", conn)
	//将实例放回池中
	pool.Put(conn)
	conn = pool.Get()
	fmt.Println("conn2:", conn)
}
