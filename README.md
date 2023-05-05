# kazaana

个人瞎折腾的错误包装工具。



默认的`error`缺少调用栈与参数信息, 多次传递后直接头皮发麻。

> 阿库娅大人说: "今天能敲代码就今天敲吧，反正明天也不一定学去敲，如果敲不了，最后咕咕咕就是了！"
> 
> 我老婆说: "我讨厌的事有三件『办不到、好累、好麻烦』这三句话非常不好, 会抹杀人类所拥有的无限可能。"

添加了调用栈保存、参数、时间。具体输出由外部实现, 用例格式参考了`panic`。



## 概览

错误包装

```go
func f1() error {
	err1 := f0()
	if err1 != nil {
		return kazaana.Wrap(err1, 1)
	}

	// do something...
	return nil
}
```



日志输出(example)

```go
[Kazaana] error catch, 2023-05-05 22:03:31
err0
    Time: 2023-05-05 22:03:31.918
    Args: 0
    /xxxx/kazaana/internal/example/main.go:13 +main.f0
    /xxxx/kazaana/internal/example/main.go:17 +main.f1
    /xxxx/kazaana/internal/example/main.go:26 +main.main
    /usr/local/go/src/runtime/proc.go:250 +runtime.main
    /usr/local/go/src/runtime/asm_amd64.s:1598 +runtime.goexit
        Time: 2023-05-05 22:03:31.918
        Args: 1
        /xxxx/kazaana/internal/example/main.go:19 +main.f1
```



## 获取

```shell
go get github.com/M-Quadra/kazaana/v3
```



## 说明

- 自定义新建错误

```
kazaana.Config.Creator = xxx
```

影响方法`kazaana.New`通过string创建自定义错误。类似`errors.New`。

- 自定义错误包装

```
kazaana.Config.Wrapper = yyy
```

影响方法`kazaana.Wrap`, 包装现有error。默认追加记录当前代码位置。

- 其他

```
errors.As(err, target)
````

当err与target均为`kazaana.Error`时, target将使用err本身赋值, 而不会寻找被包装的error

```
errors.Unwrap(err)
```

将返回内部error, 即`err.Raw`



## Todo

- [ ] errors.Join 支持

[填坑记录](./his.md)