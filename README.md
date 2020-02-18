# kazaana

error tracking for go

由于发生error时的调用栈输出只截止上一个函数入口, 如果打印的`error`已经通过多个函数函数传递...头皮发麻

> 阿库娅大人说: "今天能敲代码就今天敲吧，反正明天也不一定学去敲，如果敲不了，最后咕咕咕就是了！"
> 
> 我老婆说: "我讨厌的事有三件『办不到、好累、好麻烦』这三句话非常不好, 会抹杀人类所拥有的无限可能。"

那干脆自己针对发生`error`的地方保存调用栈, 在后续回调中依然可以找到出错位置

每次输出时附带`error`捕获的时间, 格式为`2006-01-02 15:04:05 .Unix() .UnixNano()`

调用栈输出格式参考了`vscode`下的输出形式

因为不想为`nil`判断烦恼, 所以使用了结构体

# 速食一览

运行`./test/main.go`即可预览输出效果

```
error happen:
     2020-02-18 22:05:37 1582034737 1582034737352519000
     parsing time "1970-01-01 08:00:00" as "2006-01-02 15:04:051": cannot parse "" as "1"
     /Users/m_quadra/go/src/github.com/m_quadra/kazaana/test/main.go:29 +0x10a188d
     /Users/m_quadra/go/src/github.com/m_quadra/kazaana/test/main.go:20 +0x10a1633
     /Users/m_quadra/go/src/github.com/m_quadra/kazaana/test/main.go:11 +0x10a1634
     /usr/local/go/src/runtime/proc.go:203 +0x102acbd
     /usr/local/go/src/runtime/asm_amd64.s:1357 +0x1053080
```

# 食用指北

```
kazaana.FirstCallers = 5
```

首次捕获错误信息时保存的调用栈层数, 默认为5

```
kazaana.New(err)
```

构建新的`kazaana.Error`, 并保存当前调用栈信息


```
err := kazaana.Error{}
err.HasError()
```

检查是否发生错误, 若是, 则输出错误信息

```
err := kazaana.Error{}
err.CheckError()
```

检查是否发生错误, 无输出, 用于`kazaana.Error`传递过程中的检查

```
kazaana.HasError(err)
```

检查`error`是否发生错误, 若是, 则输出错误信息
