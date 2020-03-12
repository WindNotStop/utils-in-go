package sync

import (
	"fmt"
	"sync"
)

func Map() {
	myMap := sync.Map{}
	//存
	myMap.Store("a", 97)
	myMap.Store(10, 100)
	//取
	if v, ok := myMap.Load("a"); ok {
		fmt.Println(v)
	}
	if v, ok := myMap.Load(10); ok {
		fmt.Println(v)
	}
	//遍历
	myMap.Range(func(key, value interface{}) bool {
		fmt.Println(key," ",value)
		return true
	})
	//删除
	myMap.Delete(10)
	//先读key对应值，若key不存在，则存储value
	if act, ok := myMap.LoadOrStore(10, 20); ok {
		fmt.Println("key existed,v=", act)
	} else {
		fmt.Println("key not existed,v=", act)
	}

}
