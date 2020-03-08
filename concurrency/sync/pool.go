package sync

import (
	"sync"
)

func Pool() {
	//池中存放可复用的切片
	slicePool := &sync.Pool{
		New: func() interface{} {
			v := make([]interface{}, 0, 100)
			return v
		},
	}
	//获取一个实例
	slice := slicePool.Get().([]interface{})
	//对实例进行操作
	slice = append(slice, 1, 2, 3)
	//清空切片
	slice = slice[:0]
	//放回实例
	slicePool.Put(slice)
}
