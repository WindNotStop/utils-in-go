# codestyle

常用代码设计参考

- **withOption** 
```go
 f(args..., With*(opt), With*(opt))
```
- **middleware** 
```go
 service = *Middleware(middleware)(service)
```
 - **strategy** 
 ```go
 Strategy(s*).Start(args...)
```
 - **callback** 
 ```go
 f(args..., callback)
```