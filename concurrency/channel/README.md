# channel

通过channel通信来实现内存访问同步

- **SimpleChannel** - 一个简单的channel通信
- **SimpleForSelect** - 一个简单的for-select循环
- **Pipeline** - 一个简单的流处理
- **DoneChannel** - 一个可终止协程的channel
- **OrChannel** - 一个管理多个done的channel
- **ErrorsChannel** - 一个可传递errors的channel
- **LimitedPipeline** - 一个限制长度的流处理
- **FanIn&Out** - 一种将高耗时的流操作复制多个副本并发的处理
- **TeeChannel** - 一个将流数据复制多份的channel
- **BridgeChannel** - 一个处理多个channel的channel

channel操作对应状态表

| 操作| channel状态 | 结果 |
| --- | --- | --- |
| read | nil | 阻塞 |
| | 打开且非空 | 输出值 |
| | 打开但空 | 阻塞 |
| | 关闭 | 默认值，false |
    
| 操作| channel状态 | 结果 |
| --- | --- | --- |
| write | nil | 阻塞 |
| | 打开且不满 | 输入值 |
| | 打开但满 | 阻塞 |
| | 关闭 | panic |
    
| 操作| channel状态 | 结果 |
| --- | --- | --- |
| close | nil | panic |
| | 打开且非空 | 关闭 |
| | 打开但空 | 关闭 |
| | 关闭 | panic |


