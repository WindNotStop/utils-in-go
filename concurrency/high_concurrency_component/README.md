# high_concurrency_component

高并发组件

- **simple_goroutine_pool** - 一个简单的协程池
- **string_[]byte** - string与[]byte快速转换
  > 用于[]byte分配在堆上且不会修改的情况下,能提升10到20倍性能
- **simple_filter** - 一个简单的过滤器
- **simple_bloom** - 一个简单的布隆过滤器
- **simple_cache** - 一个简单的freecache缓存